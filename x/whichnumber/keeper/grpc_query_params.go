package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (k Keeper) GetParams(c context.Context, _ *types.QueryGetParamsRequest) (*types.QueryGetParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetStoredParams(ctx)

	return &types.QueryGetParamsResponse{
		Params: params,
	}, nil
}
