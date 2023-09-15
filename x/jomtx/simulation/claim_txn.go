package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jomluz/jomtx/x/jomtx/keeper"
	"github.com/jomluz/jomtx/x/jomtx/types"
)

func SimulateMsgClaimTxn(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgClaimTxn{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ClaimTxn simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ClaimTxn simulation not implemented"), nil, nil
	}
}