package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTxn = "create_txn"
	TypeMsgUpdateTxn = "update_txn"
	TypeMsgDeleteTxn = "delete_txn"
)

var _ sdk.Msg = &MsgCreateTxn{}

func NewMsgCreateTxn(creator string, invoiceNo string, proofs []string, items string, remarks string, files []string) *MsgCreateTxn {
	return &MsgCreateTxn{
		Creator:   creator,
		InvoiceNo: invoiceNo,
		Proofs:    proofs,
		Items:     items,
		Remarks:   remarks,
		Files:     files,
	}
}

func (msg *MsgCreateTxn) Route() string {
	return RouterKey
}

func (msg *MsgCreateTxn) Type() string {
	return TypeMsgCreateTxn
}

func (msg *MsgCreateTxn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTxn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTxn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// var _ sdk.Msg = &MsgUpdateTxn{}

// func NewMsgUpdateTxn(creator string, id uint64, invoiceNo string, quantity string, items string, remarks string, files []string) *MsgUpdateTxn {
// 	return &MsgUpdateTxn{
// 		Id:        id,
// 		Creator:   creator,
// 		InvoiceNo: invoiceNo,
// 		Quantity:  quantity,
// 		Items:     items,
// 		Remarks:   remarks,
// 		Files:     files,
// 	}
// }

// func (msg *MsgUpdateTxn) Route() string {
// 	return RouterKey
// }

// func (msg *MsgUpdateTxn) Type() string {
// 	return TypeMsgUpdateTxn
// }

// func (msg *MsgUpdateTxn) GetSigners() []sdk.AccAddress {
// 	creator, err := sdk.AccAddressFromBech32(msg.Creator)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return []sdk.AccAddress{creator}
// }

// func (msg *MsgUpdateTxn) GetSignBytes() []byte {
// 	bz := ModuleCdc.MustMarshalJSON(msg)
// 	return sdk.MustSortJSON(bz)
// }

// func (msg *MsgUpdateTxn) ValidateBasic() error {
// 	_, err := sdk.AccAddressFromBech32(msg.Creator)
// 	if err != nil {
// 		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
// 	}
// 	return nil
// }

// var _ sdk.Msg = &MsgDeleteTxn{}

// func NewMsgDeleteTxn(creator string, id uint64) *MsgDeleteTxn {
// 	return &MsgDeleteTxn{
// 		Id:      id,
// 		Creator: creator,
// 	}
// }
// func (msg *MsgDeleteTxn) Route() string {
// 	return RouterKey
// }

// func (msg *MsgDeleteTxn) Type() string {
// 	return TypeMsgDeleteTxn
// }

// func (msg *MsgDeleteTxn) GetSigners() []sdk.AccAddress {
// 	creator, err := sdk.AccAddressFromBech32(msg.Creator)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return []sdk.AccAddress{creator}
// }

// func (msg *MsgDeleteTxn) GetSignBytes() []byte {
// 	bz := ModuleCdc.MustMarshalJSON(msg)
// 	return sdk.MustSortJSON(bz)
// }

// func (msg *MsgDeleteTxn) ValidateBasic() error {
// 	_, err := sdk.AccAddressFromBech32(msg.Creator)
// 	if err != nil {
// 		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
// 	}
// 	return nil
// }
