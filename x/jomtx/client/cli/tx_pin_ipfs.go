package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/wms2537/jomtx/x/jomtx/types"
)

var _ = strconv.Itoa(0)

func CmdPinIpfs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pin-ipfs [cid] [expected-size]",
		Short: "Broadcast message pinIpfs",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCid := args[0]
			argExpectedSize, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPinIpfs(
				clientCtx.GetFromAddress().String(),
				argCid,
				argExpectedSize,
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
