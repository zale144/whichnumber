package keeper

import (
	"context"
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

	// check if the game is finished
	if game.Status == types.GameStatus_GAME_STATUS_FINISHED {
		return nil, types.ErrGameFinished
	}

	// check if the game is committing
	if game.Status == types.GameStatus_GAME_STATUS_COMMITTING {
		return nil, sdkerrors.Wrapf(types.ErrCommitNotTimeout, "seconds left: %d", int(game.CommitTimeout.Sub(ctx.BlockTime()).Seconds()))
	}

	// check  that the reveal timeout hasn't passed
	if !game.RevealTimeout.IsZero() && ctx.BlockTime().After(game.RevealTimeout) {
		return nil, types.ErrRevealTimeout
	}

	// check if the player has already revealed
	_, revFound := m.k.GetNumberReveal(game, request.Player)
	if revFound {
		return nil, types.ErrPlayerAlreadyRevealed
	}

	numberCommit, found := m.k.GetNumberCommit(game, request.Player)
	if !found {
		return nil, types.ErrPlayerNotCommitted
	}

	// reveal the Number
	// calculate the Number's commitment, must match the one stored
	commit := types.CalculateNumberCommit(request.Number, request.Salt)
	if commit != numberCommit.Commit {
		return nil, types.ErrInvalidReveal
	}

	// store the reveal
	reveal := &types.NumberReveal{
		PlayerAddress: request.Player,
		Number:        request.Number,
		Salt:          request.Salt,
		CreatedAt:     ctx.BlockTime(),
	}

	params := m.k.GetStoredParams(ctx)
	reveal.HasWon(game.SecretNumber, params)

	game.PlayerReveals = append(game.PlayerReveals, reveal)
	m.k.SetStoredGame(ctx, game)

	// emit event for the new reveal
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

	return &types.MsgRevealNumberResponse{}, nil
}
