package keeper

import (
	"context"
	"strconv"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jomluz/jomtx/x/jomtx/types"
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
	val.Customer = msg.Creator

	k.SetTxn(ctx, val)
	return &types.MsgClaimTxnResponse{}, nil
}
