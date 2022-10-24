package cosmoslottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/keeper"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.TxCounter != nil {
		k.SetTxCounter(ctx, *genState.TxCounter)
	}
	// Set all the betChart
	for _, elem := range genState.BetChartList {
		k.SetBetChart(ctx, elem)
	}
	// Set if defined
	if genState.FeeCounter != nil {
		k.SetFeeCounter(ctx, *genState.FeeCounter)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all txCounter
	txCounter, found := k.GetTxCounter(ctx)
	if found {
		genesis.TxCounter = &txCounter
	}
	genesis.BetChartList = k.GetAllBetChart(ctx)
	// Get all feeCounter
	feeCounter, found := k.GetFeeCounter(ctx)
	if found {
		genesis.FeeCounter = &feeCounter
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
