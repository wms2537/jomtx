package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jomluz/jomtx/x/jomtx/types"
)

func (k msgServer) PinIpfs(goCtx context.Context, msg *types.MsgPinIpfs) (*types.MsgPinIpfsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	// msg.Creator
	_ = ctx

	size := strconv.FormatUint(msg.ExpectedSize, 10)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.PinIpfsEventType,
			sdk.NewAttribute(types.PinIpfsEventCID, msg.Cid),
			sdk.NewAttribute(types.PinIpfsEventExpectedSize, size),
		),
	)

	return &types.MsgPinIpfsResponse{}, nil
}
