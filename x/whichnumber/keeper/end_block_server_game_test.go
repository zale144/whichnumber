package keeper_test

import (
	"context"
	"sort"
	"testing"
	"time"

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
func TestKeeper_EndBlocker(t *testing.T) {
	type fields struct {
		cdc          codec.BinaryCodec
		storeKey     types.StoreKey
		memKey       types.StoreKey
		paramstore   ptypes.Subspace
		addressCodec codec.Codec
		bankKeeper   btypes.BankKeeper
	}
	type args struct {
		ctx types.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test EndBlocker",
			fields: fields{
				cdc:          nil,
				storeKey:     nil,
				memKey:       nil,
				// NewSubspace(cdc codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key sdk.StoreKey, tkey sdk.StoreKey, name string)
				paramstore:   ptypes.NewSubspace(
					btypes.ModuleName,
					)
				addressCodec: nil,
				bankKeeper:   nil,
			},
			args: args{
				ctx: types.UnwrapSDKContext(context.Background()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Keeper{
				cdc:          tt.fields.cdc,
				storeKey:     tt.fields.storeKey,
				memKey:       tt.fields.memKey,
				paramstore:   tt.fields.paramstore,
				addressCodec: tt.fields.addressCodec,
				bankKeeper:   tt.fields.bankKeeper,
			}
			k.EndBlocker(tt.args.ctx)
		})
	}
}
*/

func TestKeeper_EndBlocker(t *testing.T) {
	msgServer, keeper, context := setupMsgServerWithOneGame(t)
	ctx := sdk.UnwrapSDKContext(context)

	gameId := int64(1)

	_, err := msgServer.CommitNumber(context, &types.MsgCommitNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Commit: "2aa150bd4875fae49b3b7daac782fb08eff71c2bb973dd3ad9d5ae2c97279c7d",
	})
	require.NoError(t, err)

	g, found := keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)
	require.Equal(t, 1, len(g.PlayerCommits))

	systemInfo, found := keeper.GetStoredSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:     2,
		FifoHeadId: 1,
		FifoTailId: 1,
	}, systemInfo)

	g, found = keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)
	commitmentTimeoutExp := g.CommitTimeout
	g.CommitTimeout = g.CommitTimeout.Add(-1 * time.Minute)
	keeper.SetStoredGame(ctx, g)

	_, err = msgServer.RevealNumber(context, &types.MsgRevealNumber{
		Player: testutil.Bob,
		GameId: gameId,
		Number: 39,
		Salt:   "86530148b65e664713932cbe74842e27a170ab997594330bef420234e816defe",
	})
	require.NoError(t, err)

	g, found = keeper.GetStoredGame(ctx, gameId)
	require.True(t, found)

	keeper.EndBlocker(ctx)

	systemInfo, found = keeper.GetStoredSystemInfo(ctx)
	require.True(t, found)
	require.EqualValues(t, types.SystemInfo{
		NextId:     2,
		FifoHeadId: -1,
		FifoTailId: -1,
	}, systemInfo)

	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 3)

	revealTimeoutEvent := events[2]

	sort.Slice(revealTimeoutEvent.Attributes, func(i, j int) bool {
		return revealTimeoutEvent.Attributes[i].Key < revealTimeoutEvent.Attributes[j].Key
	})

	require.EqualValues(t, sdk.StringEvent{
		Type: "zale144.whichnumber.whichnumber.EventRevealTimeout",
		Attributes: []sdk.Attribute{
			{Key: "game_id", Value: "\"1\""},
			{Key: "number_of_players", Value: "\"1\""},
			{Key: "reveal_timeout", Value: "\"" + g.RevealTimeout.Format("2006-01-02T15:04:05Z") + "\""},
		},
	}, revealTimeoutEvent)

	newGameEvent := events[1]

	sort.Slice(newGameEvent.Attributes, func(i, j int) bool {
		return newGameEvent.Attributes[i].Key < newGameEvent.Attributes[j].Key
	})

	require.EqualValues(t, sdk.StringEvent{
		Type: "zale144.whichnumber.whichnumber.EventNewGame",
		Attributes: []sdk.Attribute{
			{Key: "commit_timeout", Value: "\"" + commitmentTimeoutExp.Format("2006-01-02T15:04:05Z") + "\""},
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

func setupMsgServerWithOneGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	server, k, ctx, _, escrow := setupMsgServerWithOneGameWithMock(t)
	escrow.ExpectAny(ctx)
	return server, k, ctx
}

func setupMsgServerWithOneGameWithMock(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	k, ctx := testkeeper.WhichNumberKeeperWithMocks(t, bankMock)

	whichnumber.InitGenesis(ctx, *k, *types.DefaultGenesis())

	server := keeper.NewMsgServerImpl(*k)
	sdkCtx := sdk.WrapSDKContext(ctx)

	resp, err := server.NewGame(sdkCtx, &types.MsgNewGame{
		Creator:      testutil.Alice,
		SecretNumber: 42,
		EntryFee:     sdk.NewCoin("stake", sdk.NewInt(100)),
		Reward:       sdk.NewCoin("stake", sdk.NewInt(1000)),
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), resp.GameId)

	return server, *k, sdkCtx, ctrl, bankMock
}
