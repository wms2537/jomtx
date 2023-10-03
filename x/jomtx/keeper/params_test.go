package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/wms2537/jomtx/testutil/keeper"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.JomtxKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
