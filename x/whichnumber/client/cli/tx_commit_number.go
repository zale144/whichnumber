package cli

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func CmdCommitNumber() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-number [gameId] [number]",
		Short: "Broadcast message commitNumber",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			gameId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("couldn't parse gameId %s into uint64", args[0])
			}

			number, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("couldn't parse number %s into uint64", args[1])
			}

			salt, err := salt(32)
			if err != nil {
				return fmt.Errorf("couldn't generate salt: %w", err)
			}

			cmd.Println("copy your salt for the reveal: ", salt)

			commit := types.CalculateNumberCommit(number, salt)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return fmt.Errorf("couldn't get client context: %w", err)
			}

			msg := types.NewMsgCommitNumber(
				clientCtx.GetFromAddress().String(),
				gameId,
				commit,
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

func salt(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
