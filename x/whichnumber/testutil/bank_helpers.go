package testutil

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (escrow *MockBankKeeper) ExpectAny(context context.Context) {
	escrow.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	escrow.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
}

func coinsOf(amount uint64, denom string) sdk.Coins {
	return sdk.Coins{
		sdk.Coin{
			Denom:  denom,
			Amount: sdk.NewInt(int64(amount)),
		},
	}
}

func (escrow *MockBankKeeper) ExpectPay(context context.Context, who string, amount uint64) *gomock.Call {
	return escrow.ExpectPayWithDenom(context, who, amount, sdk.DefaultBondDenom)
}

func (escrow *MockBankKeeper) ExpectPayWithDenom(context context.Context, who string, amount uint64, denom string) *gomock.Call {
	whoAddr, err := sdk.AccAddressFromBech32(who)
	if err != nil {
		panic(err)
	}
	return escrow.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), whoAddr, types.ModuleName, coinsOf(amount, denom))
}

func (escrow *MockBankKeeper) ExpectRefund(context context.Context, who string, amount uint64) *gomock.Call {
	return escrow.ExpectRefundWithDenom(context, who, amount, sdk.DefaultBondDenom)
}

func (escrow *MockBankKeeper) ExpectRefundWithDenom(context context.Context, who string, amount uint64, denom string) *gomock.Call {
	whoAddr, err := sdk.AccAddressFromBech32(who)
	if err != nil {
		panic(err)
	}
	return escrow.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), types.ModuleName, whoAddr, coinsOf(amount, denom))
}
