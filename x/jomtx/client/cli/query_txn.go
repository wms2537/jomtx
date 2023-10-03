package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

// func CmdListTxn() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "list-txn",
// 		Short: "list all txn",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			pageReq, err := client.ReadPageRequest(cmd.Flags())
// 			if err != nil {
// 				return err
// 			}

// 			queryClient := types.NewQueryClient(clientCtx)

// 			params := &types.QueryAllTxnRequest{
// 				Pagination: pageReq,
// 			}

// 			res, err := queryClient.TxnAll(cmd.Context(), params)
// 			if err != nil {
// 				return err
// 			}

// 			return clientCtx.PrintProto(res)
// 		},
// 	}

// 	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
// 	flags.AddQueryFlagsToCmd(cmd)

// 	return cmd
// }

func CmdShowTxn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-txn [id]",
		Short: "shows a txn",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetTxnRequest{
				Id: id,
			}

			res, err := queryClient.Txn(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
