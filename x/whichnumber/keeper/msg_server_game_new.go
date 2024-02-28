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

	systemInfo, found := m.k.GetStoredSystemInfo(ctx)
	if !found {
		return nil, types.ErrCannotGetSystemInfo
	}

	params := m.k.GetStoredParams(ctx)
	si := systemInfo
	newIndex := si.NextId

	blockTime := ctx.BlockTime()
	game := types.Game{
		Id:            newIndex,
		Creator:       newGame.Creator,
		SecretNumber:  newGame.SecretNumber,
		EntryFee:      newGame.EntryFee,
		Reward:        newGame.Reward,
		CommitTimeout: blockTime.Add(time.Second * time.Duration(params.CommitTimeout)),
		Status:        types.GameStatus_GAME_STATUS_COMMITTING,
		BeforeId:      types.NoFifoId,
		AfterId:       types.NoFifoId,
	}

	if err := game.Validate(); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidGame, "id: %d; error: %s", newIndex, err)
	}

	creator, err := types.GetPlayerAddress(newGame.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "creator: %s", newGame.Creator)
	}

	// send the reward coins from the creator to the module
	if err := m.k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(newGame.Reward)); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrFailedToSendCoinsToModule, "error: %s", err)
	}

	// emit event for the coins sent to the module
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGameCreatorDeposit{
			Creator: newGame.Creator,
			GameId:  strconv.FormatInt(newIndex, 10),
			Amount:  newGame.Reward.String(),
		},
	); err != nil {
		m.k.Logger(ctx).Error("Error emitting creator deposit event", "error", err)
	}

	m.k.SendToFifoTail(ctx, &game, &systemInfo)
	m.k.SetStoredGame(ctx, game)
	systemInfo.NextId++
	m.k.SetStoredSystemInfo(ctx, systemInfo)

	// emit event for the new game
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGameNew{
			Creator:       newGame.Creator,
			GameId:        strconv.FormatInt(newIndex, 10),
			EntryFee:      newGame.EntryFee.String(),
			MaxPlayers:    params.MaxPlayersPerGame,
			Reward:        newGame.Reward.String(),
			CommitTimeout: game.CommitTimeout.Format(time.RFC3339),
			Status:        types.GameStatus_GAME_STATUS_COMMITTING,
		},
	); err != nil {
		m.k.Logger(ctx).Error("Error emitting new game event", "error", err)
	}

	return &types.MsgNewGameResponse{
		GameId: newIndex,
	}, nil
}
