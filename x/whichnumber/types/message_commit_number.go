package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCommitNumber = "commit_number"

var _ sdk.Msg = &MsgCommitNumber{}

func NewMsgCommitNumber(player string, gameId int64, commit string) *MsgCommitNumber {
	return &MsgCommitNumber{
		Player: player,
		GameId: gameId,
		Commit: commit,
	}
}

func (msg *MsgCommitNumber) Route() string {
	return RouterKey
}

func (msg *MsgCommitNumber) Type() string {
	return TypeMsgCommitNumber
}

func (msg *MsgCommitNumber) GetSigners() []sdk.AccAddress {
	player, err := GetPlayerAddress(msg.Player)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{player}
}

func (msg *MsgCommitNumber) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCommitNumber) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Player); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid player address (%s)", err)
	}
	return nil
}
