package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jomluz/jomtx/x/jomtx/types"
)

func TestTxnMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateTxn(ctx, &types.MsgCreateTxn{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

// func TestTxnMsgServerUpdate(t *testing.T) {
// 	creator := "A"

// 	tests := []struct {
// 		desc    string
// 		request *types.MsgUpdateTxn
// 		err     error
// 	}{
// 		{
// 			desc:    "Completed",
// 			request: &types.MsgUpdateTxn{Creator: creator},
// 		},
// 		{
// 			desc:    "Unauthorized",
// 			request: &types.MsgUpdateTxn{Creator: "B"},
// 			err:     sdkerrors.ErrUnauthorized,
// 		},
// 		{
// 			desc:    "Unauthorized",
// 			request: &types.MsgUpdateTxn{Creator: creator, Id: 10},
// 			err:     sdkerrors.ErrKeyNotFound,
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			srv, ctx := setupMsgServer(t)
// 			_, err := srv.CreateTxn(ctx, &types.MsgCreateTxn{Creator: creator})
// 			require.NoError(t, err)

// 			_, err = srv.UpdateTxn(ctx, tc.request)
// 			if tc.err != nil {
// 				require.ErrorIs(t, err, tc.err)
// 			} else {
// 				require.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestTxnMsgServerDelete(t *testing.T) {
// 	creator := "A"

// 	tests := []struct {
// 		desc    string
// 		request *types.MsgDeleteTxn
// 		err     error
// 	}{
// 		{
// 			desc:    "Completed",
// 			request: &types.MsgDeleteTxn{Creator: creator},
// 		},
// 		{
// 			desc:    "Unauthorized",
// 			request: &types.MsgDeleteTxn{Creator: "B"},
// 			err:     sdkerrors.ErrUnauthorized,
// 		},
// 		{
// 			desc:    "KeyNotFound",
// 			request: &types.MsgDeleteTxn{Creator: creator, Id: 10},
// 			err:     sdkerrors.ErrKeyNotFound,
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			srv, ctx := setupMsgServer(t)

// 			_, err := srv.CreateTxn(ctx, &types.MsgCreateTxn{Creator: creator})
// 			require.NoError(t, err)
// 			_, err = srv.DeleteTxn(ctx, tc.request)
// 			if tc.err != nil {
// 				require.ErrorIs(t, err, tc.err)
// 			} else {
// 				require.NoError(t, err)
// 			}
// 		})
// 	}
// }
