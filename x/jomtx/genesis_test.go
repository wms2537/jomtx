package jomtx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/wms2537/jomtx/testutil/keeper"
	"github.com/wms2537/jomtx/testutil/nullify"
	"github.com/wms2537/jomtx/x/jomtx"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TxnList: []types.Txn{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TxnCount: 2,
		SystemInfo: &types.SystemInfo{
			FifoHeadIndex: 100,
			FifoTailIndex: 50,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JomtxKeeper(t)
	jomtx.InitGenesis(ctx, *k, genesisState)
	got := jomtx.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TxnList, got.TxnList)
	require.Equal(t, genesisState.TxnCount, got.TxnCount)
	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}
