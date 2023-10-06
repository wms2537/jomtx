package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

func (k Keeper) ForfeitExpiredGames(goCtx context.Context) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get FIFO information
	systemInfo, found := k.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	txnIndex := systemInfo.FifoHeadIndex
	var storedTxn types.Txn
	for {
		// Finished moving along
		if txnIndex == systemInfo.FifoTailIndex {
			break
		}
		storedTxn, found = k.GetTxn(ctx, txnIndex)
		if !found {
			panic(fmt.Sprintf("Fifo head game not found %d", systemInfo.FifoHeadIndex))
		}
		timestamp, err := time.Parse(time.RFC3339, storedTxn.Timestamp)
		if err != nil {
			panic(err)
		}
		deadline := timestamp.Add(time.Minute * 15)
		if deadline.Before(ctx.BlockTime()) {
			// Game is past deadline
			systemInfo.FifoHeadIndex++
			storedTxn.Customer = "unclaimed"
			txnIndex = systemInfo.FifoHeadIndex

		} else {
			// All other games after are active anyway
			break
		}
	}

	k.SetSystemInfo(ctx, systemInfo)
}
