package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/wms2537/jomtx/testutil/keeper"
	"github.com/wms2537/jomtx/testutil/nullify"
	"github.com/wms2537/jomtx/x/jomtx/keeper"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

func createNTxn(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Txn {
	items := make([]types.Txn, n)
	for i := range items {
		items[i].Id = keeper.AppendTxn(ctx, items[i])
	}
	return items
}

func TestTxnGet(t *testing.T) {
	keeper, ctx := keepertest.JomtxKeeper(t)
	items := createNTxn(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTxn(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTxnRemove(t *testing.T) {
	keeper, ctx := keepertest.JomtxKeeper(t)
	items := createNTxn(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTxn(ctx, item.Id)
		_, found := keeper.GetTxn(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTxnGetAll(t *testing.T) {
	keeper, ctx := keepertest.JomtxKeeper(t)
	items := createNTxn(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTxn(ctx)),
	)
}

func TestTxnCount(t *testing.T) {
	keeper, ctx := keepertest.JomtxKeeper(t)
	items := createNTxn(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTxnCount(ctx))
}
