package keeper_test

import (
	"testing"

	testkeeper "github.com/jomluz/jomtx/testutil/keeper"
	"github.com/jomluz/jomtx/x/jomtx/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.JomtxKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
