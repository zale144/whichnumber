package keeper

import (
	"context"
	"math"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (m msgServer) RevealNumber(goCtx context.Context, request *types.MsgRevealNumber) (*types.MsgRevealNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.GasMeter().ConsumeGas(types.GameRevealGas, "Reveal number")

	game, found := m.k.GetStoredGame(ctx, request.GameId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game_id: %d", request.GameId)
	}

	// check if player can reveal
	if err := canReveal(ctx, game, request); err != nil {
		return nil, err
	}

	won, proximity := m.hasWon(ctx, game.SecretNumber, request.Number)
	reveal := &types.NumberReveal{
		PlayerAddress: request.Player,
		Number:        request.Number,
		Salt:          request.Salt,
		IsWinner:      won,
		Proximity:     proximity,
		CreatedAt:     ctx.BlockTime(),
	}

	game.PlayerReveals = append(game.PlayerReveals, reveal)
	m.k.SetStoredGame(ctx, game)
	m.emitRevealEvents(ctx, game, request)

	return &types.MsgRevealNumberResponse{}, nil
}

func canReveal(ctx sdk.Context, game types.Game, request *types.MsgRevealNumber) error {
	// check if the game is finished
	if game.Status == types.GameStatus_GAME_STATUS_FINISHED {
		return types.ErrGameFinished
	}

	// check if the game is committing
	if game.Status == types.GameStatus_GAME_STATUS_COMMITTING {
		return sdkerrors.Wrapf(types.ErrCommitNotTimeout, "seconds left: %d", int(game.CommitTimeout.Sub(ctx.BlockTime()).Seconds()))
	}

	// check  that the reveal timeout hasn't passed
	if !game.RevealTimeout.IsZero() && ctx.BlockTime().After(game.RevealTimeout) {
		return types.ErrRevealTimeout
	}

	// charge for loop gas up front
	loopGas := uint64(types.LoopIterationGas * (len(game.PlayerCommits) + len(game.PlayerReveals)))
	ctx.GasMeter().ConsumeGas(loopGas, "Loop iteration")

	var loopGasConsumed uint64
	// refund the gas if the loop ends early
	defer func() {
		if refundGas := loopGas - loopGasConsumed; refundGas > 0 {
			ctx.GasMeter().RefundGas(refundGas, "Loop iteration refund")
		}
	}()

	// check if the player has already revealed
	for _, reveal := range game.PlayerReveals {
		loopGasConsumed += types.LoopIterationGas
		if reveal.PlayerAddress == request.Player {
			return types.ErrPlayerAlreadyRevealed
		}
	}

	// check if the player has committed
	var numberCommit *types.NumberCommit
	for _, commit := range game.PlayerCommits {
		loopGasConsumed += types.LoopIterationGas
		if commit.PlayerAddress == request.Player {
			numberCommit = commit
			break
		}
	}
	if numberCommit == nil {
		return types.ErrPlayerNotCommitted
	}
	// reveal the Number
	// calculate the Number's commitment, must match the one stored
	if commit := types.CalculateNumberCommit(request.Number, request.Salt); commit != numberCommit.Commit {
		return types.ErrInvalidReveal
	}

	return nil
}

func (m msgServer) hasWon(ctx sdk.Context, secretNumber, revealNumber int64) (won bool, proximity uint64) {
	params := m.k.GetStoredParams(ctx)
	minGuess := secretNumber - int64(params.MinDistanceToWin)
	maxGuess := secretNumber + int64(params.MinDistanceToWin)

	if revealNumber >= minGuess && revealNumber <= maxGuess {
		won = true
		proximity = uint64(math.Abs(float64(revealNumber - secretNumber)))
	}
	return
}

func (m msgServer) emitRevealEvents(ctx sdk.Context, game types.Game, request *types.MsgRevealNumber) {
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventGameNewReveal{
			GameId:          strconv.FormatInt(game.Id, 10),
			Player:          request.Player,
			NumberOfReveals: uint64(len(game.PlayerReveals)),
			NumberOfCommits: uint64(len(game.PlayerCommits)),
			Timestamp:       ctx.BlockTime().Format(time.RFC3339),
		}); err != nil {
		m.k.Logger(ctx).Error("Error emitting new reveal event", "error", err)
	}
}
