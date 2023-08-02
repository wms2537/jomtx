package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimTxn = "claim_txn"

var _ sdk.Msg = &MsgClaimTxn{}

func NewMsgClaimTxn(creator string, txnId string) *MsgClaimTxn {
	return &MsgClaimTxn{
		Creator: creator,
		TxnId:   txnId,
	}
}

func (msg *MsgClaimTxn) Route() string {
	return RouterKey
}

func (msg *MsgClaimTxn) Type() string {
	return TypeMsgClaimTxn
}

func (msg *MsgClaimTxn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimTxn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimTxn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
