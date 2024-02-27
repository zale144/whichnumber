package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateParams{}, "whichnumber/UpdateParams", nil)
	cdc.RegisterConcrete(&MsgNewGame{}, "whichnumber/NewGame", nil)
	cdc.RegisterConcrete(&MsgCommitNumber{}, "whichnumber/CommitNumber", nil)
	cdc.RegisterConcrete(&MsgRevealNumber{}, "whichnumber/RevealNumber", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpdateParams{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgNewGame{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgCommitNumber{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgRevealNumber{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
