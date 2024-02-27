package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func CmdRevealNumber() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reveal-number [gameId] [number] [salt]",
		Short: "Broadcast message revealNumber",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			gameId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("couldn't parse gameId %s into uint64", args[0])
			}

			number, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("couldn't parse number %s into int64", args[1])
			}

			salt := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return fmt.Errorf("couldn't get client context: %w", err)
			}

			msg := types.NewMsgRevealNumber(
				clientCtx.GetFromAddress().String(),
				gameId,
				number,
				salt,
			)
			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("basic validation failed: %w", err)
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
