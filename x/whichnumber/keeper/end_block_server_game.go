package keeper

import (
	"fmt"
	"math"
	"math/big"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	systemInfo, found := k.GetStoredSystemInfo(ctx)
	if !found {
		k.Logger(ctx).Error("System info not found")
		return
	}

	// iterate over all active games
	k.IterateStoredGames(ctx, systemInfo.FifoHeadId, func(game types.Game) (stop bool, nextId int64) {
		// get the current block time
		now := ctx.BlockTime()

		// if the game has no players and the commit/reveal timeout has passed, delete the game and refund the creator
		if len(game.PlayerCommits) == 0 && now.After(game.CommitTimeout) || len(game.PlayerReveals) == 0 && !game.RevealTimeout.IsZero() && now.After(game.RevealTimeout) {
			// add game to list of games to delete
			k.RemoveFromFifo(ctx, &game, &systemInfo)
			k.RemoveStoredGame(ctx, game.Id)

			// refund the creator
			if err := k.sendCoinsToPlayer(ctx, game.Creator, game.Reward); err != nil {
				k.Logger(ctx).Error("Error refunding creator", "error", err)
				return true, types.NoFifoId // stop iteration
			}
			return false, systemInfo.FifoHeadId
		}

		// no player revealed so there's no reveal timeout set yet
		if game.RevealTimeout.IsZero() {
			return false, game.AfterId
		}

		// if the reveal timeout has not passed and not all players have revealed, wait for the next iteration
		if now.Before(game.RevealTimeout) && len(game.PlayerReveals) < len(game.PlayerCommits) {
			return false, game.AfterId
		}

		// game has ended, delete it from the FIFO
		k.RemoveFromFifo(ctx, &game, &systemInfo)
		k.RemoveStoredGame(ctx, game.Id)

		// calculate winners
		winners, totalProximity := k.selectWinners(game, ctx)

		// distribute rewards to winners based on proximity to secret number
		winners, err := k.distributeRewards(ctx, game, winners, totalProximity)
		if err != nil {
			k.Logger(ctx).Error("Error distributing rewards", "error", err)
			return true, types.NoFifoId // stop iteration
		}

		numLosers := len(game.PlayerReveals) - len(winners)

		// calculate how much to refund to game creator
		losersEntryFees := sdk.NewCoin(game.EntryFee.Denom, game.EntryFee.Amount.Mul(sdk.NewInt(int64(numLosers))))

		// send any remaining funds to game creator
		if losersEntryFees.IsPositive() {
			if err := k.sendCoinsToPlayer(ctx, game.Creator, losersEntryFees); err != nil {
				k.Logger(ctx).Error("Error refunding creator", "error", err)
				return true, types.NoFifoId // stop iteration
			}
		}

		// emit event for game end
		if err := ctx.EventManager().EmitTypedEvent(
			&types.EventGameEnd{
				GameId:  strconv.FormatInt(game.Id, 10),
				Winners: winners,
			},
		); err != nil {
			k.Logger(ctx).Error("Error emitting game end event", "error", err)
		}
		return false, systemInfo.FifoHeadId // continue iteration
	})

	// update system info
	k.SetStoredSystemInfo(ctx, systemInfo)
}

func (k Keeper) selectWinners(game types.Game, ctx sdk.Context) (winners []*types.Winner, totalProximity uint64) {
	// calculate the correct guess range
	params := k.GetStoredParams(ctx)
	minGuess := game.SecretNumber - int64(params.MinDistanceToWin)
	maxGuess := game.SecretNumber + int64(params.MinDistanceToWin)

	// iterate over guesses and calculate winners
	for _, guess := range game.PlayerReveals {
		if guess.Number >= minGuess && guess.Number <= maxGuess {
			// calculate proximity to secret number
			proximity := uint64(math.Abs(float64(guess.Number - game.SecretNumber)))

			// add player to winners list along with their proximity
			winners = append(winners, &types.Winner{
				Player:    guess.PlayerAddress,
				Proximity: proximity,
			})

			// add proximity to total
			totalProximity += proximity
		}
	}
	return
}

func (k Keeper) distributeRewards(ctx sdk.Context, game types.Game, winners []*types.Winner, totalProximity uint64) ([]*types.Winner, error) {
	for i := range winners {
		// calculate reward proportionally based on proximity
		rewProx := big.NewInt(0).Mul(game.Reward.Amount.BigInt(), big.NewInt(int64(winners[i].Proximity)))
		rewProx = big.NewInt(0).Div(rewProx, big.NewInt(int64(totalProximity)))
		reward := sdk.NewCoin(game.Reward.Denom, sdk.NewIntFromBigInt(rewProx))
		winners[i].Reward = reward.String()
		toSend := reward.Add(game.EntryFee) // add entry fee to reward

		// send reward to winner
		if err := k.sendCoinsToPlayer(ctx, winners[i].Player, toSend); err != nil {
			return nil, fmt.Errorf("failed to distribute player reward: %w", err)
		}
	}
	return winners, nil
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
