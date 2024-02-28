package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/zale144/whichnumber/x/whichnumber/testutil"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func TestKeeper_MsgServer_CommitNumber(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGame(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	// commit a number
	_, err := msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	game, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)
	require.Equal(t, 1, len(game.PlayerCommits))
	require.Equal(t, testutil.Bob, game.PlayerCommits[0].PlayerAddress)
	require.Equal(t, "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d", game.PlayerCommits[0].Commit)

	systemInfo, found := keeper.GetStoredSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:     2,
		FifoHeadId: 1,
		FifoTailId: 1,
	}, systemInfo)
}

func TestKeeper_MsgServer_CommitNumber_GameNotFound(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGame(t)

	_, err := msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: 2,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.Error(t, err)
	require.Equal(t, "game_id: 2: game by id not found", err.Error())
}

func TestKeeper_MsgServer_CommitNumber_GameFinished(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGame(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	// commit a number
	_, err := msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	game, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)
	game.Status = types.GameStatus_GAME_STATUS_FINISHED
	keeper.SetStoredGame(ctx, game)

	_, err = msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.Error(t, err)
	require.Equal(t, "game is already finished", err.Error())
}

func TestKeeper_MsgServer_CommitNumber_GameNotCommitting(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGame(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	// commit a number
	_, err := msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	game, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)
	game.Status = types.GameStatus_GAME_STATUS_REVEALING
	keeper.SetStoredGame(ctx, game)

	_, err = msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.Error(t, err)
	require.Equal(t, "game is not committing", err.Error())
}

func TestKeeper_MsgServer_CommitNumber_GameFull(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGame(t)

	gameId := int64(1)

	// commit a number
	_, err := msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	// commit a number
	_, err = msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Charlie,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	_, err = msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Dave,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})

	_, err = msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Eve,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.Error(t, err)
	require.Equal(t, "max players: 3: game is full", err.Error())
}

func TestKeeper_MsgServer_CommitNumber_CommitTimeout(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGame(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	// commit a number
	_, err := msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	game, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)
	game.CommitTimeout = ctx.BlockTime().Add(-1)
	keeper.SetStoredGame(ctx, game)

	_, err = msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Charlie,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.Error(t, err)
	expectedError := fmt.Sprintf("commit timeout: %s: commit timeout has passed", game.CommitTimeout.Format("2006-01-02T15:04:05Z07:00"))
	require.Equal(t, expectedError, err.Error())
}

func TestKeeper_MsgServer_CommitNumber_CreatorNotPlayer(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGame(t)

	gameId := int64(1)

	_, err := msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Alice,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.Error(t, err)
	require.Equal(t, "game creator is not a player", err.Error())
}

func TestKeeper_MsgServer_CommitNumber_PlayerAlreadyCommitted(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneGame(t)

	gameId := int64(1)

	// commit a number
	_, err := msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	_, err = msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.Error(t, err)
	require.Equal(t, "player has already committed", err.Error())
}
