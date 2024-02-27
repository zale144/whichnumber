package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevealNumber = "reveal_number"

var _ sdk.Msg = &MsgRevealNumber{}

func NewMsgRevealNumber(player string, gameId int64, number int64, salt string) *MsgRevealNumber {
	return &MsgRevealNumber{
		Player: player,
		GameId: gameId,
		Number: number,
		Salt:   salt,
	}
}

func (msg *MsgRevealNumber) Route() string {
	return RouterKey
}

func (msg *MsgRevealNumber) Type() string {
	return TypeMsgRevealNumber
}

func (msg *MsgRevealNumber) GetSigners() []sdk.AccAddress {
	player, err := GetPlayerAddress(msg.Player)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{player}
}

func (msg *MsgRevealNumber) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevealNumber) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Player); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid player address (%s)", err)
	}
	return nil
}
