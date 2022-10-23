package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
	"github.com/spf13/cobra"
)

func CmdListBetChart() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-bet-chart",
		Short: "list all betChart",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllBetChartRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.BetChartAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowBetChart() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-bet-chart [account-name]",
		Short: "shows a betChart",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAccountName := args[0]

			params := &types.QueryGetBetChartRequest{
				AccountName: argAccountName,
			}

			res, err := queryClient.BetChart(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
