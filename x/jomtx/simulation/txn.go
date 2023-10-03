package simulation

import (
	"math/rand"

	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/wms2537/jomtx/x/jomtx/keeper"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

func SimulateMsgCreateTxn(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		msg := &types.MsgCreateTxn{
			Creator: simAccount.Address.String(),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// func SimulateMsgUpdateTxn(
// 	ak types.AccountKeeper,
// 	bk types.BankKeeper,
// 	k keeper.Keeper,
// ) simtypes.Operation {
// 	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
// 	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
// 		var (
// 			simAccount = simtypes.Account{}
// 			txn        = types.Txn{}
// 			msg        = &types.MsgUpdateTxn{}
// 			allTxn     = k.GetAllTxn(ctx)
// 			found      = false
// 		)
// 		for _, obj := range allTxn {
// 			simAccount, found = FindAccount(accs, obj.Creator)
// 			if found {
// 				txn = obj
// 				break
// 			}
// 		}
// 		if !found {
// 			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "txn creator not found"), nil, nil
// 		}
// 		msg.Creator = simAccount.Address.String()
// 		msg.Id = txn.Id

// 		txCtx := simulation.OperationInput{
// 			R:               r,
// 			App:             app,
// 			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
// 			Cdc:             nil,
// 			Msg:             msg,
// 			MsgType:         msg.Type(),
// 			Context:         ctx,
// 			SimAccount:      simAccount,
// 			ModuleName:      types.ModuleName,
// 			CoinsSpentInMsg: sdk.NewCoins(),
// 			AccountKeeper:   ak,
// 			Bankkeeper:      bk,
// 		}
// 		return simulation.GenAndDeliverTxWithRandFees(txCtx)
// 	}
// }

// func SimulateMsgDeleteTxn(
// 	ak types.AccountKeeper,
// 	bk types.BankKeeper,
// 	k keeper.Keeper,
// ) simtypes.Operation {
// 	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
// 	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
// 		var (
// 			simAccount = simtypes.Account{}
// 			txn        = types.Txn{}
// 			msg        = &types.MsgUpdateTxn{}
// 			allTxn     = k.GetAllTxn(ctx)
// 			found      = false
// 		)
// 		for _, obj := range allTxn {
// 			simAccount, found = FindAccount(accs, obj.Creator)
// 			if found {
// 				txn = obj
// 				break
// 			}
// 		}
// 		if !found {
// 			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "txn creator not found"), nil, nil
// 		}
// 		msg.Creator = simAccount.Address.String()
// 		msg.Id = txn.Id

// 		txCtx := simulation.OperationInput{
// 			R:               r,
// 			App:             app,
// 			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
// 			Cdc:             nil,
// 			Msg:             msg,
// 			MsgType:         msg.Type(),
// 			Context:         ctx,
// 			SimAccount:      simAccount,
// 			ModuleName:      types.ModuleName,
// 			CoinsSpentInMsg: sdk.NewCoins(),
// 			AccountKeeper:   ak,
// 			Bankkeeper:      bk,
// 		}
// 		return simulation.GenAndDeliverTxWithRandFees(txCtx)
// 	}
// }
