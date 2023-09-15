package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jomluz/jomtx/testutil/network"
	"github.com/jomluz/jomtx/testutil/nullify"
	"github.com/jomluz/jomtx/x/jomtx/client/cli"
	"github.com/jomluz/jomtx/x/jomtx/types"
)

func networkWithTxnObjects(t *testing.T, n int) (*network.Network, []types.Txn) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	for i := 0; i < n; i++ {
		txn := types.Txn{
			Id: uint64(i),
		}
		nullify.Fill(&txn)
		state.TxnList = append(state.TxnList, txn)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.TxnList
}

func TestShowTxn(t *testing.T) {
	net, objs := networkWithTxnObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc string
		id   string
		args []string
		err  error
		obj  types.Txn
	}{
		{
			desc: "found",
			id:   fmt.Sprintf("%d", objs[0].Id),
			args: common,
			obj:  objs[0],
		},
		{
			desc: "not found",
			id:   "not_found",
			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{tc.id}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowTxn(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetTxnResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Txn)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Txn),
				)
			}
		})
	}
}

// func TestListTxn(t *testing.T) {
// 	net, objs := networkWithTxnObjects(t, 5)

// 	ctx := net.Validators[0].ClientCtx
// 	request := func(next []byte, offset, limit uint64, total bool) []string {
// 		args := []string{
// 			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
// 		}
// 		if next == nil {
// 			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
// 		} else {
// 			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
// 		}
// 		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
// 		if total {
// 			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
// 		}
// 		return args
// 	}
// 	t.Run("ByOffset", func(t *testing.T) {
// 		step := 2
// 		for i := 0; i < len(objs); i += step {
// 			args := request(nil, uint64(i), uint64(step), false)
// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListTxn(), args)
// 			require.NoError(t, err)
// 			var resp types.QueryAllTxnResponse
// 			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 			require.LessOrEqual(t, len(resp.Txn), step)
// 			require.Subset(t,
// 				nullify.Fill(objs),
// 				nullify.Fill(resp.Txn),
// 			)
// 		}
// 	})
// 	t.Run("ByKey", func(t *testing.T) {
// 		step := 2
// 		var next []byte
// 		for i := 0; i < len(objs); i += step {
// 			args := request(next, 0, uint64(step), false)
// 			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListTxn(), args)
// 			require.NoError(t, err)
// 			var resp types.QueryAllTxnResponse
// 			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 			require.LessOrEqual(t, len(resp.Txn), step)
// 			require.Subset(t,
// 				nullify.Fill(objs),
// 				nullify.Fill(resp.Txn),
// 			)
// 			next = resp.Pagination.NextKey
// 		}
// 	})
// 	t.Run("Total", func(t *testing.T) {
// 		args := request(nil, 0, uint64(len(objs)), true)
// 		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListTxn(), args)
// 		require.NoError(t, err)
// 		var resp types.QueryAllTxnResponse
// 		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
// 		require.NoError(t, err)
// 		require.Equal(t, len(objs), int(resp.Pagination.Total))
// 		require.ElementsMatch(t,
// 			nullify.Fill(objs),
// 			nullify.Fill(resp.Txn),
// 		)
// 	})
// }
