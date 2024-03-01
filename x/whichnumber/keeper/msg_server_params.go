package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (m msgServer) UpdateParams(goCtx context.Context, params *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.GasMeter().ConsumeGas(types.ParamsUpdateGas, "UpdateParams")

	if _, err := sdk.AccAddressFromBech32(params.Authority); err != nil {
		return nil, fmt.Errorf("invalid authority address: %w", err)
	}

	/*
		TODO: enable this
		if authority := m.k.GetAuthority(); !strings.EqualFold(params.Authority, authority) {
			return nil, fmt.Errorf("unauthorized, authority does not match the module's authority: got %s, want %s", params.Authority, authority)
		}*/

	if err := params.Params.Validate(); err != nil {
		return nil, err
	}

	// update the params
	m.k.SetParams(ctx, params.Params)
	// emit event
	if err := ctx.EventManager().EmitTypedEvent(
		&types.EventParamsUpdated{
			Params: params.Params,
		}); err != nil {
		m.k.Logger(ctx).Error("failed to emit event", "error", err)
	}
	return &types.MsgUpdateParamsResponse{}, nil
}
