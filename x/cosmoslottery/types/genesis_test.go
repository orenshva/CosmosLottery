package types_test

import (
	"testing"

	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				TxCounter: &types.TxCounter{
					Count: 2,
				},
				BetChartList: []types.BetChart{
					{
						AccountName: "0",
					},
					{
						AccountName: "1",
					},
				},
				FeeCounter: &types.FeeCounter{
					Count: 13,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated betChart",
			genState: &types.GenesisState{
				BetChartList: []types.BetChart{
					{
						AccountName: "0",
					},
					{
						AccountName: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
