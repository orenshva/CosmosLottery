package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/orenshva/CosmosLottery/testutil/keeper"
	"github.com/orenshva/CosmosLottery/testutil/nullify"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
)

func TestTxCounterQuery(t *testing.T) {
	keeper, ctx := keepertest.CosmoslotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTxCounter(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTxCounterRequest
		response *types.QueryGetTxCounterResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTxCounterRequest{},
			response: &types.QueryGetTxCounterResponse{TxCounter: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TxCounter(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
