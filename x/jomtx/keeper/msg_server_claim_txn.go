package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jomluz/jomtx/x/jomtx/types"
)

func (k msgServer) ClaimTxn(goCtx context.Context, msg *types.MsgClaimTxn) (*types.MsgClaimTxnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgClaimTxnResponse{}, nil
}
