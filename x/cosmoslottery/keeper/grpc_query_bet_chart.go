package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BetChartAll(c context.Context, req *types.QueryAllBetChartRequest) (*types.QueryAllBetChartResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var betCharts []types.BetChart
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	betChartStore := prefix.NewStore(store, types.KeyPrefix(types.BetChartKeyPrefix))

	pageRes, err := query.Paginate(betChartStore, req.Pagination, func(key []byte, value []byte) error {
		var betChart types.BetChart
		if err := k.cdc.Unmarshal(value, &betChart); err != nil {
			return err
		}

		betCharts = append(betCharts, betChart)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBetChartResponse{BetChart: betCharts, Pagination: pageRes}, nil
}

func (k Keeper) BetChart(c context.Context, req *types.QueryGetBetChartRequest) (*types.QueryGetBetChartResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBetChart(
		ctx,
		req.AccountName,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBetChartResponse{BetChart: val}, nil
}
