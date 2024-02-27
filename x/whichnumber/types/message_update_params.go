package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdateParams = "update_params"

func (msg *MsgUpdateParams) Route() string {
	return RouterKey
}

func (msg *MsgUpdateParams) Type() string {
	return TypeMsgUpdateParams
}

func (msg *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	player, err := GetPlayerAddress(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{player}
}

func (msg *MsgUpdateParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateParams) ValidateBasic() error {
	if msg.Params.CommitTimeout == 0 {
		return fmt.Errorf("commit timeout cannot be 0")
	}

	if msg.Params.RevealTimeout == 0 {
		return fmt.Errorf("reveal timeout cannot be 0")
	}

	if msg.Params.MaxPlayersPerGame == 0 {
		return fmt.Errorf("max players per game cannot be 0")
	}

	return nil
}
