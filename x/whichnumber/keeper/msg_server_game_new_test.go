package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/zale144/whichnumber/x/whichnumber/testutil"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func TestMsgServer_NewGame(t *testing.T) {
	msgServer, k, context, _, _ := setupMsgServerWithMock(t)

	// create a new game
	resp, err := msgServer.NewGame(context, &types.MsgNewGame{
		Creator:      testutil.Bob,
		SecretNumber: 42,
		EntryFee:     sdk.NewCoin("stake", sdk.NewInt(100)),
		Reward:       sdk.NewCoin("stake", sdk.NewInt(1000)),
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), resp.GameId)

	ctx := sdk.UnwrapSDKContext(context)

	// check if the game was created
	game, found := k.GetStoredGame(ctx, resp.GameId)
	require.True(t, found)
	require.Equal(t, testutil.Bob, game.Creator)
	require.Equal(t, int64(42), game.SecretNumber)
	require.Equal(t, sdk.NewCoin("stake", sdk.NewInt(100)), game.EntryFee)
	require.Equal(t, sdk.NewCoin("stake", sdk.NewInt(1000)), game.Reward)
	require.Equal(t, types.GameStatus_GAME_STATUS_COMMITTING, game.Status)
	require.Equal(t, int64(1), game.Id)
	require.Equal(t, int64(-1), game.BeforeId)
	require.Equal(t, int64(-1), game.AfterId)
}

func TestMsgServer_NewGame_InvalidEntryFee(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerWithMock(t)

	// create a new game
	_, err := msgServer.NewGame(context, &types.MsgNewGame{
		Creator:      testutil.Bob,
		SecretNumber: 42,
		EntryFee:     sdk.NewCoin("stake", sdk.NewInt(0)),
		Reward:       sdk.NewCoin("stake", sdk.NewInt(1000)),
	})
	require.Error(t, err)
	require.Equal(t, "id: 1; error: entry fee cannot be zero: game is invalid", err.Error())
}

func TestMsgServer_NewGame_MinimumReward(t *testing.T) {
	msgServer, k, context, _, _ := setupMsgServerWithMock(t)
	ctx := sdk.UnwrapSDKContext(context)

	// update params, set min reward to 1000
	params := k.GetStoredParams(ctx)
	params.MinReward = sdk.NewCoin("stake", sdk.NewInt(1000))
	k.SetParams(ctx, params)

	// create a new game
	_, err := msgServer.NewGame(context, &types.MsgNewGame{
		Creator:      testutil.Bob,
		SecretNumber: 42,
		EntryFee:     sdk.NewCoin("stake", sdk.NewInt(100)),
		Reward:       sdk.NewCoin("stake", sdk.NewInt(999)),
	})
	require.Error(t, err)
	require.Equal(t, "reward: 999stake; minimum: 1000stake: invalid reward", err.Error())
}

func TestMsgServer_NewGame_InvalidReward(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerWithMock(t)

	// create a new game
	_, err := msgServer.NewGame(context, &types.MsgNewGame{
		Creator:      testutil.Bob,
		SecretNumber: 42,
		EntryFee:     sdk.NewCoin("stake", sdk.NewInt(100)),
		Reward:       sdk.NewCoin("stake", sdk.NewInt(0)),
	})
	require.Error(t, err)
	require.Equal(t, "reward: 0stake; minimum: 1000stake: invalid reward", err.Error())

}

func TestMsgServer_NewGame_InvalidCreator(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerWithMock(t)

	// create a new game
	_, err := msgServer.NewGame(context, &types.MsgNewGame{
		Creator:      "invalid",
		SecretNumber: 42,
		EntryFee:     sdk.NewCoin("stake", sdk.NewInt(100)),
		Reward:       sdk.NewCoin("stake", sdk.NewInt(1000)),
	})
	require.Error(t, err)
	require.Equal(t, "failed to deposit reward", err.Error())
}

func TestMsgServer_NewGame_FailedToSendCoinsToModule(t *testing.T) {
	msgServer, k, context, _, bankMock := setupMsgServerWithMockNoBankExpect(t)

	// mock failed to send coins to module
	bankMock.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("")).Times(1)

	// create a new game
	_, err := msgServer.NewGame(context, &types.MsgNewGame{
		Creator:      testutil.Bob,
		SecretNumber: 42,
		EntryFee:     sdk.NewCoin("stake", sdk.NewInt(100)),
		Reward:       sdk.NewCoin("stake", sdk.NewInt(1000)),
	})
	require.Error(t, err)
	require.Equal(t, "failed to deposit reward", err.Error())

	// check if the game was created
	ctx := sdk.UnwrapSDKContext(context)
	game, found := k.GetStoredGame(ctx, 2)
	require.False(t, found)
	require.Equal(t, types.Game{}, game)
}
