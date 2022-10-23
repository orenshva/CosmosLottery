package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
)

// SetBetChart set a specific betChart in the store from its index
func (k Keeper) SetBetChart(ctx sdk.Context, betChart types.BetChart) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetChartKeyPrefix))
	b := k.cdc.MustMarshal(&betChart)
	store.Set(types.BetChartKey(
		betChart.AccountName,
	), b)
}

// GetBetChart returns a betChart from its index
func (k Keeper) GetBetChart(
	ctx sdk.Context,
	accountName string,

) (val types.BetChart, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetChartKeyPrefix))

	b := store.Get(types.BetChartKey(
		accountName,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBetChart removes a betChart from the store
func (k Keeper) RemoveBetChart(
	ctx sdk.Context,
	accountName string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetChartKeyPrefix))
	store.Delete(types.BetChartKey(
		accountName,
	))
}

// GetAllBetChart returns all betChart
func (k Keeper) GetAllBetChart(ctx sdk.Context) (list []types.BetChart) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BetChartKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BetChart
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
