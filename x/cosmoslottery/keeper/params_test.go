package keeper_test

import (
	"testing"

	testkeeper "github.com/orenshva/CosmosLottery/testutil/keeper"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmoslotteryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
