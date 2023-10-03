package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/wms2537/jomtx/x/jomtx/keeper"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

func SimulateMsgPinIpfs(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgPinIpfs{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the PinIpfs simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "PinIpfs simulation not implemented"), nil, nil
	}
}
