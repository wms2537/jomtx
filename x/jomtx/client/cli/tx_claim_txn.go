package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jomluz/jomtx/x/jomtx/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdClaimTxn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-txn [txn-id]",
		Short: "Broadcast message claim-txn",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTxnId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaimTxn(
				clientCtx.GetFromAddress().String(),
				argTxnId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
