package keeper

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	systemInfo, found := k.GetStoredSystemInfo(ctx)
	if !found {
		k.Logger(ctx).Error("System info not found")
		return
	}

	// get the current block time
	now := ctx.BlockTime()
	// get the stored params
	params := k.GetStoredParams(ctx)

	// iterate over all active games
	k.IterateStoredGames(ctx, systemInfo.FifoHeadId, func(game types.Game) (stop bool, nextId int64) {
		// if the game has no players and the commit/reveal timeout has passed, delete the game and refund the creator
		if len(game.PlayerCommits) == 0 && now.After(game.CommitTimeout) || len(game.PlayerReveals) == 0 && !game.RevealTimeout.IsZero() && now.After(game.RevealTimeout) {
			// delete the game
			k.RemoveFromFifo(ctx, &game, &systemInfo)
			k.RemoveStoredGame(ctx, game.Id)

			// refund the creator
			if err := k.sendCoinsToPlayer(ctx, game.Creator, game.Reward); err != nil {
				k.Logger(ctx).Error("Error refunding creator", "error", err)
				return true, types.NoFifoId // stop iteration
			}
			return false, systemInfo.FifoHeadId
		}

		numPlayerCommits := uint64(len(game.PlayerCommits))

		// if the commit timeout has passed or the game is full, start the reveal and wait for the next iteration
		if (now.After(game.CommitTimeout) || numPlayerCommits == params.MaxPlayersPerGame) && game.Status == types.GameStatus_GAME_STATUS_COMMITTING {
			game.RevealTimeout = ctx.BlockTime().Add(time.Second * time.Duration(params.RevealTimeout))
			game.Status = types.GameStatus_GAME_STATUS_REVEALING

			k.SetStoredGame(ctx, game)
			// emit event for the reveal timeout
			if err := ctx.EventManager().EmitTypedEvent(
				&types.EventGameCommitFinished{
					GameId:          strconv.FormatInt(game.Id, 10),
					CommitTimeout:   game.RevealTimeout.Format(time.RFC3339),
					NumberOfCommits: numPlayerCommits,
				},
			); err != nil {
				k.Logger(ctx).Error("Error emitting commit finished event", "error", err)
			}

			return false, game.AfterId
		}

		// if we are either still committing, or the reveal timeout has not passed yet and
		// not all players have revealed their number,
		// wait for the next iteration
		if game.Status == types.GameStatus_GAME_STATUS_COMMITTING ||
			game.Status == types.GameStatus_GAME_STATUS_REVEALING &&
				now.Before(game.RevealTimeout) &&
				len(game.PlayerReveals) < len(game.PlayerCommits) {
			return false, game.AfterId
		}

		game.Status = types.GameStatus_GAME_STATUS_FINISHED

		// game has ended, delete it from the FIFO
		k.RemoveFromFifo(ctx, &game, &systemInfo)

		// calculate winners
		totalProximity := k.selectWinners(&game)

		// distribute rewards to winners based on proximity to secret number
		if err := k.distributeRewards(ctx, &game, totalProximity); err != nil {
			k.Logger(ctx).Error("Error distributing rewards", "error", err)
			return true, types.NoFifoId // stop iteration
		}

		numLosers := len(game.PlayerReveals) - len(game.Winners)

		// calculate how much to refund to game creator
		losersEntryFees := sdk.NewCoin(game.EntryFee.Denom, game.EntryFee.Amount.Mul(sdk.NewInt(int64(numLosers))))

		// send any remaining funds to game creator
		if losersEntryFees.IsPositive() {
			if err := k.sendCoinsToPlayer(ctx, game.Creator, losersEntryFees); err != nil {
				k.Logger(ctx).Error("Error refunding creator", "error", err)
				return true, types.NoFifoId // stop iteration
			}
		}

		k.SetStoredGame(ctx, game)

		// emit event for game end
		if err := ctx.EventManager().EmitTypedEvent(
			&types.EventGameEnd{
				GameId:  strconv.FormatInt(game.Id, 10),
				Winners: game.Winners,
			},
		); err != nil {
			k.Logger(ctx).Error("Error emitting game end event", "error", err)
		}
		return false, systemInfo.FifoHeadId // continue iteration
	})

	// update system info
	k.SetStoredSystemInfo(ctx, systemInfo)
}

func (k Keeper) selectWinners(game *types.Game) (totalProximity uint64) {
	// iterate over guesses and add winners
	for _, guess := range game.PlayerReveals {
		if !guess.IsWinner {
			continue
		}
		game.Winners = append(game.Winners, &types.Winner{
			Player:    guess.PlayerAddress,
			Proximity: guess.Proximity,
		})
		// add proximity to total
		totalProximity += guess.Proximity
	}
	return
}

func (k Keeper) distributeRewards(ctx sdk.Context, game *types.Game, totalProximity uint64) error {
	for i := range game.Winners {
		// calculate reward proportionally based on proximity
		rewProx := big.NewInt(0).Mul(game.Reward.Amount.BigInt(), big.NewInt(int64(game.Winners[i].Proximity)))
		rewProx = big.NewInt(0).Div(rewProx, big.NewInt(int64(totalProximity)))
		reward := sdk.NewCoin(game.Reward.Denom, sdk.NewIntFromBigInt(rewProx))
		game.Winners[i].Reward = reward.String()
		toSend := reward.Add(game.EntryFee) // add entry fee to reward

		// send reward to winner
		if err := k.sendCoinsToPlayer(ctx, game.Winners[i].Player, toSend); err != nil {
			return fmt.Errorf("failed to distribute player reward: %w", err)
		}
	}
	return nil
}

func (k Keeper) sendCoinsToPlayer(ctx sdk.Context, player string, coins sdk.Coin) error {
	creatorAddress, err := types.GetPlayerAddress(player)
	if err != nil {
		return fmt.Errorf("invalid creator address: %w", err)
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddress, sdk.NewCoins(coins)); err != nil {
		return fmt.Errorf("failed to send coins to player %s: %w", player, err)
	}
	return nil
}
