package cosmoslottery_test

import (
	"testing"

	keepertest "github.com/orenshva/CosmosLottery/testutil/keeper"
	"github.com/orenshva/CosmosLottery/testutil/nullify"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TxCounter: &types.TxCounter{
			Count: 43,
		},
		BetChartList: []types.BetChart{
			{
				AccountName: "0",
			},
			{
				AccountName: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmoslotteryKeeper(t)
	cosmoslottery.InitGenesis(ctx, *k, genesisState)
	got := cosmoslottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.TxCounter, got.TxCounter)
	require.ElementsMatch(t, genesisState.BetChartList, got.BetChartList)
	// this line is used by starport scaffolding # genesis/test/assert
}
