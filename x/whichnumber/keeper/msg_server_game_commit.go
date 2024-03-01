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
	ctx.GasMeter().ConsumeGas(types.GameCommitGas, "Commit number")

	game, found := m.k.GetStoredGame(ctx, request.GameId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game_id: %d", request.GameId)
	}

	// check if the player can commit
	if err := m.canCommit(ctx, game, request.Player); err != nil {
		return nil, err
	}

	if err := m.k.sendCoinsToModule(ctx, request.Player, game.EntryFee); err != nil {
		m.k.Logger(ctx).Error("Error depositing entry fee", "error", err)
		return nil, types.ErrGameFailedToDepositEntryFee
	}

	commit := &types.NumberCommit{
		PlayerAddress: request.Player,
		Commit:        request.Commit,
		CreatedAt:     ctx.BlockTime(),
	}

	game.PlayerCommits = append(game.PlayerCommits, commit)
	m.k.SetStoredGame(ctx, game)
	m.emitCommitEvents(ctx, request, game)

	return &types.MsgCommitNumberResponse{}, nil
}

func (m msgServer) canCommit(ctx sdk.Context, game types.Game, player string) error {
	// check if the game is finished
	if game.Status == types.GameStatus_GAME_STATUS_FINISHED {
		return types.ErrGameFinished
	}

	// check if the game is committing
	if game.Status != types.GameStatus_GAME_STATUS_COMMITTING {
		return types.ErrGameNotCommitting
	}

	params := m.k.GetStoredParams(ctx)
	numPlayers := uint64(len(game.PlayerCommits))

	// if the game is full we can move the game to the next state
	if numPlayers >= params.MaxPlayersPerGame {
		return sdkerrors.Wrapf(types.ErrGameFull, "max players: %d", params.MaxPlayersPerGame)
	}

	// check if the game exists and that the commit timeout hasn't passed
	if !game.CommitTimeout.IsZero() && ctx.BlockTime().After(game.CommitTimeout) {
		return sdkerrors.Wrapf(types.ErrCommitTimeout, "commit timeout: %s", game.CommitTimeout.Format(time.RFC3339))
	}

	// check if the player is the creator
	if game.Creator == player {
		return types.ErrCreatorNotPlayer
	}

	// charge for loop gas up front
	loopGas := uint64(types.LoopIterationGas * len(game.PlayerCommits))
	ctx.GasMeter().ConsumeGas(loopGas, "Loop iteration")
	var loopGasConsumed uint64

	// check if the player has already committed
	for _, commit := range game.PlayerCommits {
		loopGasConsumed += types.LoopIterationGas
		if commit.PlayerAddress == player {
			// if the loop ended early refund the player
			refundAmount := loopGas - loopGasConsumed
			if refundAmount > 0 {
				ctx.GasMeter().RefundGas(refundAmount, "Loop iteration refund")
			}
			return types.ErrPlayerAlreadyCommitted
		}
	}

	return nil
}

func (m msgServer) emitCommitEvents(ctx sdk.Context, request *types.MsgCommitNumber, game types.Game) {
	// emit event for the player making the fee deposit
	// (even though the bank module emits an event about the transfer,
	// we want to emit a custom one that provides more context)
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGamePlayerDeposit{
			Player: request.Player,
			GameId: strconv.FormatInt(game.Id, 10),
			Amount: game.EntryFee.String(),
		},
	); err != nil {
		m.k.Logger(ctx).Error("Error emitting player deposit event", "error", err)
	}

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
}
