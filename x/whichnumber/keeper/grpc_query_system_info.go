package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zale144/whichnumber/x/whichnumber/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetSystemInfo(goCtx context.Context, _ *types.QueryGetSystemInfoRequest) (*types.QueryGetSystemInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	systemInfo, found := k.GetStoredSystemInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "system info not found")
	}

	return &types.QueryGetSystemInfoResponse{
		SystemInfo: systemInfo,
	}, nil
}
