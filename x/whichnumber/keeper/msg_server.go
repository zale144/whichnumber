package keeper

import (
	"context"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

type msgServer struct {
	k Keeper
}

func (m msgServer) NewGame(goCtx context.Context, newGame *types.MsgNewGame) (*types.MsgNewGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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
		BeforeId:      types.NoFifoId,
		AfterId:       types.NoFifoId,
	}

	if err := game.Validate(); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidGame, "id: %d; error: %s", newIndex, err)
	}

	creator, err := types.GetPlayerAddress(newGame.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, newGame.Creator)
	}

	// send the reward coins from the creator to the module
	if err = m.k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(newGame.Reward)); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrFailedToSendCoinsToModule, "%s", err)
	}

	m.k.SendToFifoTail(ctx, &game, &systemInfo)
	m.k.SetStoredGame(ctx, game)
	systemInfo.NextId++
	m.k.SetStoredSystemInfo(ctx, systemInfo)

	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventNewGame{
			Creator:       newGame.Creator,
			GameId:        strconv.FormatInt(newIndex, 10),
			EntryFee:      newGame.EntryFee.String(),
			MaxPlayers:    params.MaxPlayersPerGame,
			Reward:        newGame.Reward.String(),
			CommitTimeout: game.CommitTimeout.Format(time.RFC3339),
		},
	); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrFailedToEmitEvent, "%s", err)
	}

	return &types.MsgNewGameResponse{
		GameId: newIndex,
	}, nil
}

func (m msgServer) CommitNumber(goCtx context.Context, request *types.MsgCommitNumber) (*types.MsgCommitNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	game, found := m.k.GetStoredGame(ctx, request.GameId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%d", request.GameId)
	}

	numPlayers := uint64(len(game.PlayerCommits))
	params := m.k.GetStoredParams(ctx)

	// check if the game is full
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
		if player.PlayerAddress == request.Player {
			return nil, types.ErrPlayerAlreadyCommitted
		}
	}

	playerAddr, err := types.GetPlayerAddress(request.Player)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, request.Player)
	}

	if err = m.k.bankKeeper.SendCoinsFromAccountToModule(ctx, playerAddr, types.ModuleName, sdk.NewCoins(game.EntryFee)); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrFailedToSendCoinsToModule, "%s", err)
	}

	commit := &types.NumberCommit{
		PlayerAddress: request.Player,
		Commit:        request.Commit,
		CreatedAt:     ctx.BlockTime(),
	}

	game.PlayerCommits = append(game.PlayerCommits, commit)

	m.k.SetStoredGame(ctx, game)

	return &types.MsgCommitNumberResponse{}, nil
}

func (m msgServer) RevealNumber(goCtx context.Context, request *types.MsgRevealNumber) (*types.MsgRevealNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	game, found := m.k.GetStoredGame(ctx, request.GameId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%d", request.GameId)
	}

	// check  that the commit timeout has passed
	if !game.CommitTimeout.IsZero() && ctx.BlockTime().Before(game.CommitTimeout) {
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

	// store reveal timeout if it's the first reveal
	if game.RevealTimeout.IsZero() {
		game.RevealTimeout = ctx.BlockTime().Add(time.Second * time.Duration(params.RevealTimeout))
		// emit event for the reveal timeout
		if err := ctx.EventManager().EmitTypedEvent(
			&types.EventRevealTimeout{
				GameId:          strconv.FormatInt(game.Id, 10),
				RevealTimeout:   game.RevealTimeout.Format(time.RFC3339),
				NumberOfPlayers: uint64(len(game.PlayerCommits)),
			},
		); err != nil {
			return nil, sdkerrors.Wrapf(types.ErrFailedToEmitEvent, "%s", err)
		}
	}

	game.PlayerReveals = append(game.PlayerReveals, reveal)
	m.k.SetStoredGame(ctx, game)

	return &types.MsgRevealNumberResponse{}, nil
}

func (m msgServer) UpdateParams(goCtx context.Context, params *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: check if the sender is the admin
	// update the params
	m.k.SetParams(ctx, params.Params)
	return &types.MsgUpdateParamsResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

var _ types.MsgServer = msgServer{}
