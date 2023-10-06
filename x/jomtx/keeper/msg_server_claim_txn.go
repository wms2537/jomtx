package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wms2537/jomtx/x/jomtx/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) ClaimTxn(goCtx context.Context, msg *types.MsgClaimTxn) (*types.MsgClaimTxnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := msg.ValidateBasic()
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrInvalidCreator.Error(), msg.Creator)
	}
	id, err := strconv.ParseUint(msg.TxnId, 10, 64)
	if err != nil {
		return nil, err
	}
	val, found := k.GetTxn(
		ctx,
		id,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	if val.Customer != "" {
		return nil, status.Error(codes.AlreadyExists, "txn already claimed")
	}
	val.Customer = msg.Creator

	k.SetTxn(ctx, val)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.TxnClaimedEventType,
			sdk.NewAttribute(types.TxnClaimedEventCreator, val.Creator),
			sdk.NewAttribute(types.TxnClaimedEventTxnId, msg.TxnId),
			sdk.NewAttribute(types.TxnClaimedEventCustomer, msg.Creator),
			sdk.NewAttribute(types.TxnClaimedEventItems, val.Items),
			sdk.NewAttribute(types.TxnClaimedEventTotal, fmt.Sprintf("%d", val.Total)),
			sdk.NewAttribute(types.TxnClaimedEventCurrency, val.Currency),
			sdk.NewAttribute(types.TxnClaimedEventDecimals, fmt.Sprintf("%d", val.Decimals)),
			sdk.NewAttribute(types.TxnClaimedEventTimestamp, val.Timestamp),
		),
	)

	return &types.MsgClaimTxnResponse{}, nil
}
