package keeper

import (
	"context"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (m msgServer) NewGame(goCtx context.Context, newGame *types.MsgNewGame) (*types.MsgNewGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.GasMeter().ConsumeGas(types.GameNewGas, "New game")

	params := m.k.GetStoredParams(ctx)
	// check if the reward is not less than the minimum amount
	if newGame.Reward.IsLT(params.MinReward) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidReward, "reward: %s; minimum: %s", newGame.Reward, params.MinReward)
	}

	systemInfo, found := m.k.GetStoredSystemInfo(ctx)
	if !found {
		return nil, types.ErrCannotGetSystemInfo
	}

	game := types.Game{
		Id:            systemInfo.NextId,
		Creator:       newGame.Creator,
		SecretNumber:  newGame.SecretNumber,
		EntryFee:      newGame.EntryFee,
		Reward:        newGame.Reward,
		CommitTimeout: ctx.BlockTime().Add(time.Second * time.Duration(params.CommitTimeout)),
		Status:        types.GameStatus_GAME_STATUS_COMMITTING,
		BeforeId:      types.NoFifoId,
		AfterId:       types.NoFifoId,
	}

	if err := game.Validate(); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidGame, "id: %d; error: %s", game.Id, err)
	}

	if err := m.k.sendCoinsToModule(ctx, newGame.Creator, game.Reward); err != nil {
		m.k.Logger(ctx).Error("Error depositing reward", "error", err)
		return nil, types.ErrGameFailedToDepositReward
	}

	m.k.sendToFifoTail(ctx, &game, &systemInfo)
	m.k.SetStoredGame(ctx, game)
	m.emitNewGameEvents(ctx, game, params.MaxPlayersPerGame)

	systemInfo.NextId++
	m.k.SetStoredSystemInfo(ctx, systemInfo)

	return &types.MsgNewGameResponse{
		GameId: game.Id,
	}, nil
}

func (m msgServer) emitNewGameEvents(ctx sdk.Context, game types.Game, maxPlayers uint64) {
	// emit event for the creator making the reward deposit
	// (even though the bank module emits an event about the transfer,
	// we want to emit a custom one that provides more context)
	gameId := strconv.FormatInt(game.Id, 10)

	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGameCreatorDeposit{
			Creator: game.Creator,
			GameId:  gameId,
			Amount:  game.Reward.String(),
		},
	); err != nil {
		m.k.Logger(ctx).Error("Error emitting creator deposit event", "error", err)
	}

	// emit event for the new game
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGameNew{
			Creator:       game.Creator,
			GameId:        gameId,
			EntryFee:      game.EntryFee.String(),
			MaxPlayers:    maxPlayers,
			Reward:        game.Reward.String(),
			CommitTimeout: game.CommitTimeout.Format(time.RFC3339),
			Status:        types.GameStatus_GAME_STATUS_COMMITTING,
		},
	); err != nil {
		m.k.Logger(ctx).Error("Error emitting new game event", "error", err)
	}
}
