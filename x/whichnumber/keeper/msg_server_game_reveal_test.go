package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/zale144/whichnumber/x/whichnumber/keeper"
	"github.com/zale144/whichnumber/x/whichnumber/testutil"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func TestMsgServer_RevealNumber(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameAndCommits(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	game, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)

	// move the state to revealing
	game.Status = types.GameStatus_GAME_STATUS_REVEALING
	keeper.SetStoredGame(ctx, game)

	// reveal a number
	_, err := msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Number: 39,
		Salt:   "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe",
	})
	require.NoError(t, err)

	game, found = keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)
	require.Equal(t, 1, len(game.PlayerReveals))
	require.Equal(t, testutil.Bob, game.PlayerReveals[0].PlayerAddress)
	require.Equal(t, int64(39), game.PlayerReveals[0].Number)
	require.Equal(t, "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe", game.PlayerReveals[0].Salt)
	require.Equal(t, types.GameStatus_GAME_STATUS_REVEALING, game.Status)
}

func TestMsgServer_RevealNumber_InvalidReveal(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameAndCommits(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	game, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)

	// move the state to revealing
	game.Status = types.GameStatus_GAME_STATUS_REVEALING
	keeper.SetStoredGame(ctx, game)

	// reveal a number
	_, err := msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Number: 39,
		Salt:   "invalid salt",
	})
	require.Error(t, err)
	require.Equal(t, types.ErrInvalidReveal, err)
}

func TestMsgServer_RevealNumber_PlayerAlreadyRevealed(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameAndCommits(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	game, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)

	// move the state to revealing
	game.Status = types.GameStatus_GAME_STATUS_REVEALING
	keeper.SetStoredGame(ctx, game)

	// reveal a number
	_, err := msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Number: 39,
		Salt:   "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe",
	})
	require.NoError(t, err)

	// reveal the same number again
	_, err = msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Number: 39,
		Salt:   "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe",
	})
	require.Error(t, err)
	require.Equal(t, types.ErrPlayerAlreadyRevealed, err)
}

func TestMsgServer_RevealNumber_PlayerNotCommitted(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGame(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	game, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)

	// move the state to revealing
	game.Status = types.GameStatus_GAME_STATUS_REVEALING
	keeper.SetStoredGame(ctx, game)

	// reveal a number
	_, err := msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Number: 39,
		Salt:   "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe",
	})
	require.Error(t, err)
	require.Equal(t, types.ErrPlayerNotCommitted, err)
}

func TestMsgServer_RevealNumber_RevealTimeout(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGame(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	game, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)

	// move the state to revealing
	game.Status = types.GameStatus_GAME_STATUS_REVEALING
	game.RevealTimeout = ctx.BlockTime().Add(-1)
	keeper.SetStoredGame(ctx, game)

	// reveal a number
	_, err := msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Number: 39,
		Salt:   "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe",
	})
	require.Error(t, err)
	require.Equal(t, types.ErrRevealTimeout, err)
}

func TestMsgServer_RevealNumber_Committing(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameAndCommits(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	_, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)

	// reveal a number
	_, err := msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Number: 39,
		Salt:   "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe",
	})
	require.Error(t, err)
	require.EqualError(t, err, "seconds left: 60: commit timeout has not passed yet")
}

func TestMsgServer_RevealNumber_GameFinished(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGameAndCommits(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	game, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)

	// move the state to finished
	game.Status = types.GameStatus_GAME_STATUS_FINISHED
	keeper.SetStoredGame(ctx, game)

	// reveal a number
	_, err := msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Number: 39,
		Salt:   "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe",
	})
	require.Error(t, err)
	require.EqualError(t, err, "game is already finished")
}

func TestMsgServer_RevealNumber_GameNotFound(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGameAndCommits(t)

	// reveal a number
	_, err := msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: 2,
		Number: 39,
		Salt:   "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe",
	})
	require.Error(t, err)
	require.EqualError(t, err, "game_id: 2: game by id not found")
}

func setupMsgServerWithOneGameAndCommits(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	server, k, ctx := setupMsgServerWithOneGame(t)

	_, err := server.CommitNumber(ctx, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: 1,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	_, err = server.CommitNumber(ctx, &types.MsgCommitNumber{
		Player: testutil.Charlie,
		GameId: 1,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	_, err = server.CommitNumber(ctx, &types.MsgCommitNumber{
		Player: testutil.Eve,
		GameId: 1,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	return server, k, ctx
}
