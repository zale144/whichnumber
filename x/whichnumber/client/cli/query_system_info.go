package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func CmdShowSystemInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-system-info",
		Short: "shows systemInfo",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetSystemInfoRequest{}

			res, err := queryClient.GetSystemInfo(context.Background(), params)
			if err != nil {
				return fmt.Errorf("failed to get systemInfo: %w", err)
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
