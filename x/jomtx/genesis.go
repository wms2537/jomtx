package jomtx

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wms2537/jomtx/x/jomtx/keeper"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the txn
	for _, elem := range genState.TxnList {
		k.SetTxn(ctx, elem)
	}

	// Set txn count
	k.SetTxnCount(ctx, genState.TxnCount)
	// Set if defined
	if genState.SystemInfo != nil {
		k.SetSystemInfo(ctx, *genState.SystemInfo)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TxnList = k.GetAllTxn(ctx)
	genesis.TxnCount = k.GetTxnCount(ctx)
	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = &systemInfo
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
