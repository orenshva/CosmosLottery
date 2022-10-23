package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

var LotteryFee sdk.Coin = sdk.NewCoin("token", sdk.NewInt(5))
var MinBet sdk.Coin = sdk.NewCoin("token", sdk.NewInt(10))
var MaxBet sdk.Coin = sdk.NewCoin("token", sdk.NewInt(100))

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TxCounter:    nil,
		BetChartList: []BetChart{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in betChart
	betChartIndexMap := make(map[string]struct{})

	for _, elem := range gs.BetChartList {
		index := string(BetChartKey(elem.AccountName))
		if _, ok := betChartIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for betChart")
		}
		betChartIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
