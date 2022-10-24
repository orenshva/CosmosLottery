package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/orenshva/CosmosLottery/testutil/keeper"
	"github.com/orenshva/CosmosLottery/testutil/nullify"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/keeper"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
)

func createTestFeeCounter(keeper *keeper.Keeper, ctx sdk.Context) types.FeeCounter {
	item := types.FeeCounter{}
	keeper.SetFeeCounter(ctx, item)
	return item
}

func TestFeeCounterGet(t *testing.T) {
	keeper, ctx := keepertest.CosmoslotteryKeeper(t)
	item := createTestFeeCounter(keeper, ctx)
	rst, found := keeper.GetFeeCounter(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestFeeCounterRemove(t *testing.T) {
	keeper, ctx := keepertest.CosmoslotteryKeeper(t)
	createTestFeeCounter(keeper, ctx)
	keeper.RemoveFeeCounter(ctx)
	_, found := keeper.GetFeeCounter(ctx)
	require.False(t, found)
}
