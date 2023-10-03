package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

func (k msgServer) CreateTxn(goCtx context.Context, msg *types.MsgCreateTxn) (*types.MsgCreateTxnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var txn = types.Txn{
		Creator:   msg.Creator,
		InvoiceNo: msg.InvoiceNo,
		Proofs:    msg.Proofs,
		Items:     msg.Items,
		Remarks:   msg.Remarks,
		Files:     msg.Files,
		Total:     msg.Total,
		Currency:  msg.Currency,
		Decimals:  msg.Decimals,
	}

	id := k.AppendTxn(
		ctx,
		txn,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.TxnCreatedEventType,
			sdk.NewAttribute(types.TxnCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.TxnCreatedEventTxnId, strconv.FormatUint(id, 10)),
		),
	)

	return &types.MsgCreateTxnResponse{
		Id: id,
	}, nil
}

// func (k msgServer) UpdateTxn(goCtx context.Context, msg *types.MsgUpdateTxn) (*types.MsgUpdateTxnResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	var txn = types.Txn{
// 		Creator:   msg.Creator,
// 		Id:        msg.Id,
// 		InvoiceNo: msg.InvoiceNo,
// 		Quantity:  msg.Quantity,
// 		Items:     msg.Items,
// 		Remarks:   msg.Remarks,
// 		Files:     msg.Files,
// 	}

// 	// Checks that the element exists
// 	val, found := k.GetTxn(ctx, msg.Id)
// 	if !found {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
// 	}

// 	// Checks if the msg creator is the same as the current owner
// 	if msg.Creator != val.Creator {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
// 	}

// 	k.SetTxn(ctx, txn)

// 	return &types.MsgUpdateTxnResponse{}, nil
// }

// func (k msgServer) DeleteTxn(goCtx context.Context, msg *types.MsgDeleteTxn) (*types.MsgDeleteTxnResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	// Checks that the element exists
// 	val, found := k.GetTxn(ctx, msg.Id)
// 	if !found {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
// 	}

// 	// Checks if the msg creator is the same as the current owner
// 	if msg.Creator != val.Creator {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
// 	}

// 	k.RemoveTxn(ctx, msg.Id)

// 	return &types.MsgDeleteTxnResponse{}, nil
// }
