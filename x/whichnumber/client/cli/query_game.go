package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func CmdListGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-games",
		Short: "list all games",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return fmt.Errorf("error reading page request: %w", err)
			}

			params := &types.QueryGetGamesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GetGames(context.Background(), params)
			if err != nil {
				return fmt.Errorf("failed to list games: %w", err)
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-game [index]",
		Short: "shows a game",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			argId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("couldn't parse index")
			}

			params := &types.QueryGetGameRequest{
				Id: argId,
			}

			res, err := queryClient.GetGame(context.Background(), params)
			if err != nil {
				return fmt.Errorf("failed to get game: %w", err)
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
