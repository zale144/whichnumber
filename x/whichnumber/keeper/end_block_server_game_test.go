package keeper_test

import (
	"context"
	"sort"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/zale144/whichnumber/testutil/keeper"
	"github.com/zale144/whichnumber/x/whichnumber"
	"github.com/zale144/whichnumber/x/whichnumber/keeper"
	"github.com/zale144/whichnumber/x/whichnumber/testutil"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

/*
TODO: more test cases, like:

- game with multiple players
- multiple games
- game with no players and commit timeout passed and reveal timeout passed
- error cases...

*/

func TestKeeper_EndBlocker(t *testing.T) {
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

	systemInfo, found := keeper.GetStoredSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:     2,
		FifoHeadId: 1,
		FifoTailId: 1,
	}, systemInfo)

	// move the state to revealing
	game, found = keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)
	game.Status = types.GameStatus_GAME_STATUS_REVEALING
	keeper.SetStoredGame(ctx, game)

	// reveal the number
	_, err = msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Number: 39,
		Salt:   "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe",
	})
	require.NoError(t, err)

	// end the block
	keeper.EndBlocker(ctx)

	game, found = keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)

	systemInfo, found = keeper.GetStoredSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:     2,
		FifoHeadId: -1,
		FifoTailId: -1,
	}, systemInfo)

	require.Equal(t, types.GameStatus_GAME_STATUS_FINISHED, game.Status)
	require.Equal(t, 1, len(game.PlayerReveals))
	require.Equal(t, int64(39), game.PlayerReveals[0].Number)
	require.Equal(t, uint64(3), game.PlayerReveals[0].Proximity)
	require.Equal(t, 1, len(game.Winners))
	require.Equal(t, game.Winners[0], &types.Winner{
		Player:    testutil.Bob,
		Proximity: 3,
		Reward:    "1000stake",
	})

	// compareEvents(t, ctx, game)
}

// TODO: fix this
func compareEvents(t *testing.T, ctx sdk.Context, game types.Game) {
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 6)

	revealTimeoutEvent := events[2]

	sort.Slice(revealTimeoutEvent.Attributes, func(i, j int) bool {
		return revealTimeoutEvent.Attributes[i].Key < revealTimeoutEvent.Attributes[j].Key
	})

	require.EqualValues(t, sdk.StringEvent{
		Type: "zale144.whichnumber.whichnumber.EventRevealTimeout",
		Attributes: []sdk.Attribute{
			{Key: "game_id", Value: "\"1\""},
			{Key: "number_of_players", Value: "\"1\""},
			{Key: "reveal_timeout", Value: "\"" + game.RevealTimeout.Format("2006-01-02T15:04:05Z") + "\""},
		},
	}, revealTimeoutEvent)

	newGameEvent := events[1]

	sort.Slice(newGameEvent.Attributes, func(i, j int) bool {
		return newGameEvent.Attributes[i].Key < newGameEvent.Attributes[j].Key
	})

	require.EqualValues(t, sdk.StringEvent{
		Type: "zale144.whichnumber.whichnumber.EventNewGame",
		Attributes: []sdk.Attribute{
			{Key: "commit_timeout", Value: "\"" + game.CommitTimeout.Format("2006-01-02T15:04:05Z") + "\""},
			{Key: "creator", Value: "\"cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3\""},
			{Key: "entry_fee", Value: "\"100stake\""},
			{Key: "game_id", Value: "\"1\""},
			{Key: "max_players", Value: "\"10\""},
			{Key: "reward", Value: "\"1000stake\""},
		},
	}, newGameEvent)

	gameEndEvent := events[0]

	sort.Slice(gameEndEvent.Attributes, func(i, j int) bool {
		return gameEndEvent.Attributes[i].Key < gameEndEvent.Attributes[j].Key
	})

	require.EqualValues(t, sdk.StringEvent{
		Type: "zale144.whichnumber.whichnumber.EventGameEnd",
		Attributes: []sdk.Attribute{
			{Key: "game_id", Value: "\"1\""},
			{Key: "winners", Value: "[{\"player\":\"" + testutil.Bob + "\",\"proximity\":\"3\",\"reward\":\"1000stake\"}]"},
		},
	}, gameEndEvent)
}

func TestKeeper_EndBlocker_GameFinished(t *testing.T) {
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

	// move the state to finished
	game.Status = types.GameStatus_GAME_STATUS_FINISHED
	keeper.SetStoredGame(ctx, game)

	// end the block
	keeper.EndBlocker(ctx)

	game, found = keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)
	require.Equal(t, types.GameStatus_GAME_STATUS_FINISHED, game.Status)
}

func TestKeeper_EndBlocker_GameCommitting(t *testing.T) {
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

	// end the block
	keeper.EndBlocker(ctx)

	game, found = keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)
	require.Equal(t, types.GameStatus_GAME_STATUS_COMMITTING, game.Status)
}

func setupMsgServerWithOneGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	server, k, ctx, _, escrow := setupMsgServerWithOneGameWithMock(t)
	escrow.ExpectAny(ctx)
	return server, k, ctx
}

func setupMsgServerWithOneGameWithMock(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	server, k, sdkCtx, ctrl, bankMock := setupMsgServerWithMock(t)

	resp, err := server.NewGame(sdkCtx, &types.MsgNewGame{
		Creator:      testutil.Alice,
		SecretNumber: 42,
		EntryFee:     sdk.NewCoin("stake", sdk.NewInt(100)),
		Reward:       sdk.NewCoin("stake", sdk.NewInt(1000)),
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), resp.GameId)

	return server, k, sdkCtx, ctrl, bankMock
}

func setupMsgServerWithMock(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	k, ctx := testkeeper.WhichNumberKeeperWithMocks(t, bankMock)

	genesis := types.DefaultGenesis()
	genesis.Params.MaxPlayersPerGame = 3
	whichnumber.InitGenesis(ctx, *k, *genesis)

	server := keeper.NewMsgServerImpl(*k)
	sdkCtx := sdk.WrapSDKContext(ctx)

	return server, *k, sdkCtx, ctrl, bankMock
}

func setupMsgServerWithMockNoBankExpect(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)

	k, ctx := testkeeper.WhichNumberKeeperWithMocks(t, bankMock)

	genesis := types.DefaultGenesis()
	genesis.Params.MaxPlayersPerGame = 3
	whichnumber.InitGenesis(ctx, *k, *genesis)

	server := keeper.NewMsgServerImpl(*k)
	sdkCtx := sdk.WrapSDKContext(ctx)

	return server, *k, sdkCtx, ctrl, bankMock
}
