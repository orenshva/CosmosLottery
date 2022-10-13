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

func createTestTxCounter(keeper *keeper.Keeper, ctx sdk.Context) types.TxCounter {
	item := types.TxCounter{}
	keeper.SetTxCounter(ctx, item)
	return item
}

func TestTxCounterGet(t *testing.T) {
	keeper, ctx := keepertest.CosmoslotteryKeeper(t)
	item := createTestTxCounter(keeper, ctx)
	rst, found := keeper.GetTxCounter(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTxCounterRemove(t *testing.T) {
	keeper, ctx := keepertest.CosmoslotteryKeeper(t)
	createTestTxCounter(keeper, ctx)
	keeper.RemoveTxCounter(ctx)
	_, found := keeper.GetTxCounter(ctx)
	require.False(t, found)
}
