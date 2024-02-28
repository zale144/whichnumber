package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (m msgServer) UpdateParams(goCtx context.Context, params *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.GasMeter().ConsumeGas(types.ParamsUpdateGas, "UpdateParams")

	// TODO: check if the sender is the admin

	// update the params
	m.k.SetParams(ctx, params.Params)
	// emit event
	err := ctx.EventManager().EmitTypedEvent(
		&types.EventParamsUpdated{
			Params: params.Params,
		})
	return &types.MsgUpdateParamsResponse{}, errors.Wrap(err, "emit event")
}
