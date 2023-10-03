package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPinIpfs = "pin_ipfs"

var _ sdk.Msg = &MsgPinIpfs{}

func NewMsgPinIpfs(creator string, cid string, expectedSize uint64) *MsgPinIpfs {
	return &MsgPinIpfs{
		Creator:      creator,
		Cid:          cid,
		ExpectedSize: expectedSize,
	}
}

func (msg *MsgPinIpfs) Route() string {
	return RouterKey
}

func (msg *MsgPinIpfs) Type() string {
	return TypeMsgPinIpfs
}

func (msg *MsgPinIpfs) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPinIpfs) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPinIpfs) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
