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

	now := ctx.BlockTime()
	params := k.GetStoredParams(ctx)

	k.processGames(ctx, &systemInfo, now, params)
	k.SetStoredSystemInfo(ctx, systemInfo)
}

func (k Keeper) processGames(ctx sdk.Context, systemInfo *types.SystemInfo, now time.Time, params types.Params) {
	k.IterateStoredGames(ctx, systemInfo.FifoHeadId, func(game types.Game) int64 {
		if k.shouldRemoveGame(game, now) {
			k.removeGame(ctx, &game, systemInfo)
			return systemInfo.FifoHeadId
		}

		if k.shouldStartReveal(game, now, params) {
			k.startReveal(ctx, &game, now, params)
			return game.AfterId
		}

		if k.shouldEndGame(game, now) {
			k.endGame(ctx, &game, systemInfo)
			return systemInfo.FifoHeadId
		}

		return game.AfterId
	})
}

func (k Keeper) shouldRemoveGame(game types.Game, now time.Time) bool {
	return len(game.PlayerCommits) == 0 && now.After(game.CommitTimeout) ||
		len(game.PlayerReveals) == 0 && !game.RevealTimeout.IsZero() && now.After(game.RevealTimeout)
}

func (k Keeper) removeGame(ctx sdk.Context, game *types.Game, systemInfo *types.SystemInfo) {
	k.removeFromFifo(ctx, game, systemInfo)
	k.RemoveStoredGame(ctx, game.Id)
	// refund creator
	if err := k.refundCreator(ctx, game); err != nil {
		k.Logger(ctx).Error("Error refunding creator", "error", err)
	}
}

func (k Keeper) shouldStartReveal(game types.Game, now time.Time, params types.Params) bool {
	return (now.After(game.CommitTimeout) || uint64(len(game.PlayerCommits)) == params.MaxPlayersPerGame) &&
		game.Status == types.GameStatus_GAME_STATUS_COMMITTING
}

func (k Keeper) startReveal(ctx sdk.Context, game *types.Game, now time.Time, params types.Params) int64 {
	game.RevealTimeout = now.Add(time.Second * time.Duration(params.RevealTimeout))
	game.Status = types.GameStatus_GAME_STATUS_REVEALING
	k.SetStoredGame(ctx, *game)
	// emit event for the commit stage end
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGameCommitFinished{
			GameId:          strconv.FormatInt(game.Id, 10),
			CommitTimeout:   game.RevealTimeout.Format(time.RFC3339),
			NumberOfCommits: uint64(len(game.PlayerCommits)),
		},
	); err != nil {
		k.Logger(ctx).Error("Error emitting commit finished event", "error", err)
	}

	return game.AfterId
}

func (k Keeper) shouldEndGame(game types.Game, now time.Time) bool {
	return game.Status == types.GameStatus_GAME_STATUS_REVEALING &&
		(now.After(game.RevealTimeout) || len(game.PlayerReveals) == len(game.PlayerCommits))
}

func (k Keeper) endGame(ctx sdk.Context, game *types.Game, systemInfo *types.SystemInfo) int64 {
	// calculate winners
	var totalProximity uint64
	game.Winners, totalProximity = k.selectWinners(game)
	game.Status = types.GameStatus_GAME_STATUS_FINISHED

	// game has ended, delete it from the FIFO
	k.removeFromFifo(ctx, game, systemInfo)

	// distribute rewards to winners based on proximity to secret number
	if err := k.distributeRewards(ctx, game, totalProximity); err != nil {
		k.Logger(ctx).Error("Error distributing rewards", "error", err)
		return systemInfo.FifoHeadId
	}

	k.SetStoredGame(ctx, *game)

	// refund creator
	if len(game.Winners) == 0 {
		if err := k.refundCreator(ctx, game); err != nil {
			k.Logger(ctx).Error("Error refunding creator", "error", err)
			return systemInfo.FifoHeadId
		}
	}

	// emit event for game end
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGameEnd{
			GameId:  strconv.FormatInt(game.Id, 10),
			Winners: game.Winners,
		},
	); err != nil {
		k.Logger(ctx).Error("Error emitting game end event", "error", err)
	}
	return systemInfo.FifoHeadId
}

func (k Keeper) selectWinners(game *types.Game) (winners []*types.Winner, totalProximity uint64) {
	// iterate over guesses and add winners
	for _, guess := range game.PlayerReveals {
		if !guess.IsWinner {
			continue
		}
		winners = append(winners, &types.Winner{
			Player:    guess.PlayerAddress,
			Proximity: guess.Proximity,
			Reward:    "",
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

		// emit event for distributing reward
		if err := ctx.EventManager().EmitTypedEvent(
			&types.EventGamePlayerReward{
				GameId: strconv.FormatInt(game.Id, 10),
				Player: game.Winners[i].Player,
				Amount: reward.String(),
			},
		); err != nil {
			k.Logger(ctx).Error("Error emitting game reward event", "error", err)
		}
	}
	return nil
}

func (k Keeper) refundCreator(ctx sdk.Context, game *types.Game) error {
	// calculate how much to refund to game creator
	numLosers := len(game.PlayerReveals) - len(game.Winners)
	losersEntryFees := sdk.NewCoin(game.EntryFee.Denom, game.EntryFee.Amount.Mul(sdk.NewInt(int64(numLosers))))

	// send any remaining funds to game creator
	if err := k.sendCoinsToPlayer(ctx, game.Creator, losersEntryFees); err != nil {
		k.Logger(ctx).Error("Error refunding creator", "error", err)
		return fmt.Errorf("failed to refund creator: %w", err)
	}

	// emit event for refunding creator
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGameCreatorRefund{
			GameId:  strconv.FormatInt(game.Id, 10),
			Creator: game.Creator,
			Amount:  game.Reward.String(),
		},
	); err != nil {
		k.Logger(ctx).Error("Error emitting game end event", "error", err)
	}
	return nil
}
