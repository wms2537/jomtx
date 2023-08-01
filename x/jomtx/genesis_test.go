package jomtx_test

import (
	"testing"

	keepertest "github.com/jomluz/jomtx/testutil/keeper"
	"github.com/jomluz/jomtx/testutil/nullify"
	"github.com/jomluz/jomtx/x/jomtx"
	"github.com/jomluz/jomtx/x/jomtx/types"
	"github.com/stretchr/testify/require"
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
	// this line is used by starport scaffolding # genesis/test/assert
}
