package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

// GetStoredParams get all parameters as types.Params
func (k Keeper) GetStoredParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
