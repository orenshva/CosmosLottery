package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
	"github.com/spf13/cobra"
)

func CmdShowFeeCounter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-fee-counter",
		Short: "shows feeCounter",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetFeeCounterRequest{}

			res, err := queryClient.FeeCounter(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
