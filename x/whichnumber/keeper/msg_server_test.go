package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/zale144/whichnumber/testutil/keeper"
	"github.com/zale144/whichnumber/x/whichnumber/keeper"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.WhichnumberKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
