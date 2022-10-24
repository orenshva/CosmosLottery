package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
)

// SetFeeCounter set feeCounter in the store
func (k Keeper) SetFeeCounter(ctx sdk.Context, feeCounter types.FeeCounter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeCounterKey))
	b := k.cdc.MustMarshal(&feeCounter)
	store.Set([]byte{0}, b)
}

// GetFeeCounter returns feeCounter
func (k Keeper) GetFeeCounter(ctx sdk.Context) (val types.FeeCounter, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeCounterKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFeeCounter removes feeCounter from the store
func (k Keeper) RemoveFeeCounter(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeCounterKey))
	store.Delete([]byte{0})
}
