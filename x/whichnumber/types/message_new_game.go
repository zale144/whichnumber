package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNewGame = "new_game"

var _ sdk.Msg = &MsgNewGame{}

func NewMsgNewGame(creator string, secretNumber int64, entryFee, reward sdk.Coin) *MsgNewGame {
	return &MsgNewGame{
		Creator:      creator,
		SecretNumber: secretNumber,
		EntryFee:     entryFee,
		Reward:       reward,
	}
}

func (msg *MsgNewGame) Route() string {
	return RouterKey
}

func (msg *MsgNewGame) Type() string {
	return TypeMsgNewGame
}

func (msg *MsgNewGame) GetSigners() []sdk.AccAddress {
	creator, err := GetPlayerAddress(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewGame) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.EntryFee.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "entry fee must be positive")
	}

	if msg.Reward.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "reward must be positive")
	}

	return nil
}
