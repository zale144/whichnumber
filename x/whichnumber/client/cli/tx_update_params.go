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

func CmdUpdateParams() *cobra.Command {
	// Action: whichnumberd tx whichnumber update-params 10 20 5 10 1000stake
	cmd := &cobra.Command{
		Use:   "update-params [commitTimeout] [revealTimeout] [maxPlayersPerGame] [minDistanceToWin] [minReward]",
		Short: "Update a parameter",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return fmt.Errorf("couldn't get client context: %w", err)
			}

			commitTimeout, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("couldn't parse commitTimeout %s into int64", args[0])
			}

			revealTimeout, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("couldn't parse revealTimeout %s into int64", args[1])
			}

			maxPlayersPerGame, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return fmt.Errorf("couldn't parse maxPlayersPerGame %s into int64", args[2])
			}

			minDistanceToWin, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return fmt.Errorf("couldn't parse minDistanceToWin %s into int64", args[3])
			}

			minReward, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return fmt.Errorf("couldn't parse minReward %s into int64", args[4])
			}

			msg := &types.MsgUpdateParams{
				Authority: clientCtx.GetFromAddress().String(),
				Params: types.Params{
					CommitTimeout:     commitTimeout,
					RevealTimeout:     revealTimeout,
					MaxPlayersPerGame: maxPlayersPerGame,
					MinDistanceToWin:  minDistanceToWin,
					MinReward:         minReward,
				},
			}

			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("basic validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
