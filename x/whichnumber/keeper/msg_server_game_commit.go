package keeper

import (
	"context"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (m msgServer) CommitNumber(goCtx context.Context, request *types.MsgCommitNumber) (*types.MsgCommitNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var totalGasConsumed uint64 = types.GameCommitGas

	defer func() {
		ctx.GasMeter().ConsumeGas(totalGasConsumed, "Commit number")
	}()

	game, found := m.k.GetStoredGame(ctx, request.GameId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game_id: %d", request.GameId)
	}

	// check if the game is finished
	if game.Status == types.GameStatus_GAME_STATUS_FINISHED {
		return nil, types.ErrGameFinished
	}

	// check if the game is committing
	if game.Status != types.GameStatus_GAME_STATUS_COMMITTING {
		return nil, types.ErrGameNotCommitting
	}

	params := m.k.GetStoredParams(ctx)
	numPlayers := uint64(len(game.PlayerCommits))

	// if the game is full we can move the game to the next state
	if numPlayers >= params.MaxPlayersPerGame {
		return nil, sdkerrors.Wrapf(types.ErrGameFull, "max players: %d", params.MaxPlayersPerGame)
	}

	// check if the game exists and that the commit timeout hasn't passed
	if !game.CommitTimeout.IsZero() && ctx.BlockTime().After(game.CommitTimeout) {
		return nil, sdkerrors.Wrapf(types.ErrCommitTimeout, "commit timeout: %s", game.CommitTimeout.Format(time.RFC3339))
	}

	// check if the player is the creator
	if game.Creator == request.Player {
		return nil, types.ErrCreatorNotPlayer
	}

	// check if the player has already committed
	for _, player := range game.PlayerCommits {
		totalGasConsumed += types.LoopIterationGas
		if player.PlayerAddress == request.Player {
			return nil, types.ErrPlayerAlreadyCommitted
		}
	}

	playerAddr, err := types.GetPlayerAddress(request.Player)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, request.Player)
	}

	if err = m.k.bankKeeper.SendCoinsFromAccountToModule(ctx, playerAddr, types.ModuleName, sdk.NewCoins(game.EntryFee)); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrFailedToSendCoinsToModule, "error: %s", err)
	}

	// emit event for the coins sent to the module
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGamePlayerDeposit{
			Player: request.Player,
			GameId: strconv.FormatInt(game.Id, 10),
			Amount: game.EntryFee.String(),
		},
	); err != nil {
		m.k.Logger(ctx).Error("Error emitting player deposit event", "error", err)
	}

	commit := &types.NumberCommit{
		PlayerAddress: request.Player,
		Commit:        request.Commit,
		CreatedAt:     ctx.BlockTime(),
	}

	game.PlayerCommits = append(game.PlayerCommits, commit)

	m.k.SetStoredGame(ctx, game)

	// emit event for the new commit
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGameNewCommit{
			GameId:          strconv.FormatInt(game.Id, 10),
			Player:          request.Player,
			Commit:          request.Commit,
			NumberOfCommits: uint64(len(game.PlayerCommits)),
			Timestamp:       ctx.BlockTime().Format(time.RFC3339),
		}); err != nil {
		m.k.Logger(ctx).Error("Error emitting new commit event", "error", err)
	}

	return &types.MsgCommitNumberResponse{}, nil
}
