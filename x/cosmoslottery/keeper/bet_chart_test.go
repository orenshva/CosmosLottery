package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/orenshva/CosmosLottery/testutil/keeper"
	"github.com/orenshva/CosmosLottery/testutil/nullify"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/keeper"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBetChart(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.BetChart {
	items := make([]types.BetChart, n)
	for i := range items {
		items[i].AccountName = strconv.Itoa(i)

		keeper.SetBetChart(ctx, items[i])
	}
	return items
}

func TestBetChartGet(t *testing.T) {
	keeper, ctx := keepertest.CosmoslotteryKeeper(t)
	items := createNBetChart(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBetChart(ctx,
			item.AccountName,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBetChartRemove(t *testing.T) {
	keeper, ctx := keepertest.CosmoslotteryKeeper(t)
	items := createNBetChart(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBetChart(ctx,
			item.AccountName,
		)
		_, found := keeper.GetBetChart(ctx,
			item.AccountName,
		)
		require.False(t, found)
	}
}

func TestBetChartGetAll(t *testing.T) {
	keeper, ctx := keepertest.CosmoslotteryKeeper(t)
	items := createNBetChart(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBetChart(ctx)),
	)
}
