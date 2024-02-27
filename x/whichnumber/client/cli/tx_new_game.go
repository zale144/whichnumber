package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func CmdNewGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-game [secretNumber] [entryFee] [reward]", // e.g. whichnumbercli tx whichnumber new-game 42 11stake 121stake
		Short: "Broadcast message newGame",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			secretNumber, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("couldn't parse secretNumber %s into int64", args[0])
			}

			entryFee, err := sdk.ParseCoinNormalized(args[1]) // e.g. 10stake
			if err != nil {
				return fmt.Errorf("couldn't parse entryFee %s into sdk.Coin", args[1])
			}

			reward, err := sdk.ParseCoinNormalized(args[2]) // e.g. 100stake
			if err != nil {
				return fmt.Errorf("couldn't parse reward %s into sdk.Coin", args[2])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return fmt.Errorf("couldn't get client context: %w", err)
			}

			msg := types.NewMsgNewGame(
				clientCtx.GetFromAddress().String(),
				secretNumber,
				entryFee,
				reward,
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
