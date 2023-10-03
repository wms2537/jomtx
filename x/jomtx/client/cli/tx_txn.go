package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

func CmdCreateTxn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-txn [invoice-no] [proofs] [items] [remarks] [files] [total] [currency] [decimals]",
		Short: "Create a new txn",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argInvoiceNo := args[0]
			argProofs := strings.Split(args[1], listSeparator)
			argItems := args[2]
			argRemarks := args[3]
			argFiles := strings.Split(args[4], listSeparator)
			argTotal, _ := strconv.ParseUint(args[5], 10, 64)
			argCurrency := args[6]
			argDecimals, _ := strconv.ParseUint(args[7], 10, 32)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateTxn(clientCtx.GetFromAddress().String(), argInvoiceNo, argProofs, argItems, argRemarks, argFiles, argTotal, argCurrency, uint32(argDecimals))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// func CmdUpdateTxn() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "update-txn [id] [invoice-no] [quantity] [items] [remarks] [files]",
// 		Short: "Update a txn",
// 		Args:  cobra.ExactArgs(6),
// 		RunE: func(cmd *cobra.Command, args []string) (err error) {
// 			id, err := strconv.ParseUint(args[0], 10, 64)
// 			if err != nil {
// 				return err
// 			}

// 			argInvoiceNo := args[1]

// 			argQuantity := args[2]

// 			argItems := args[3]

// 			argRemarks := args[4]

// 			argFiles := strings.Split(args[5], listSeparator)

// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			msg := types.NewMsgUpdateTxn(clientCtx.GetFromAddress().String(), id, argInvoiceNo, argQuantity, argItems, argRemarks, argFiles)
// 			if err := msg.ValidateBasic(); err != nil {
// 				return err
// 			}
// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)

// 	return cmd
// }

// func CmdDeleteTxn() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "delete-txn [id]",
// 		Short: "Delete a txn by id",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			id, err := strconv.ParseUint(args[0], 10, 64)
// 			if err != nil {
// 				return err
// 			}

// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			msg := types.NewMsgDeleteTxn(clientCtx.GetFromAddress().String(), id)
// 			if err := msg.ValidateBasic(); err != nil {
// 				return err
// 			}
// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)

// 	return cmd
// }
