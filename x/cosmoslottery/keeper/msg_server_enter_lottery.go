package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
)

func (k msgServer) EnterLottery(goCtx context.Context, msg *types.MsgEnterLottery) (*types.MsgEnterLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check that the user isn't already participating in the lottery
	_, found := k.GetBetChart(ctx, msg.GetCreator())
	if found == true {
		return nil, sdkerrors.Wrapf(types.ErrUserAlreadyBetted, "%s", msg.GetCreator())
	}

	// Check that the bet and lottery fee are valid
	if msg.GetLotteryFee() != types.LotteryFee.Amount.Uint64() {
		return nil, sdkerrors.Wrapf(types.ErrBadLotteryFee, "%s", msg.GetLotteryFee())
	}

	if msg.GetBet() < types.MinBet.Amount.Uint64() {
		return nil, sdkerrors.Wrapf(types.ErrBetTooSmall, "%s", msg.GetBet())
	}

	if msg.GetBet() > types.MaxBet.Amount.Uint64() {
		return nil, sdkerrors.Wrapf(types.ErrBetTooBig, "%s", msg.GetBet())
	}

	// Get the user's account address
	bidderAcc, err := sdk.AccAddressFromBech32(msg.GetCreator())
	if err != nil {
		return nil, err
	}

	// Get the module's account address
	moduleAccAddr, err := sdk.AccAddressFromBech32(types.ModuleName)
	if err != nil {
		return nil, err
	}

	// Check that the user has enough balance to pay the fee and place the bet
	minAmount := msg.GetLotteryFee() + msg.GetBet()
	if k.bankKeeper.HasBalance(ctx, bidderAcc, sdk.NewCoin("token", sdk.NewInt(int64(minAmount)))) == false {
		return nil, sdkerrors.Wrapf(types.ErrUserDoesNotHaveEnoughFunds, "%s", msg.GetCreator())
	}

	// Deduct the lottery fee
	err = k.bankKeeper.DelegateCoins(ctx, bidderAcc, moduleAccAddr, sdk.NewCoins(types.LotteryFee))
	if err != nil {
		return nil, err
	}

	// Transfer the bet to the lottery pool
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidderAcc, types.ModuleName, sdk.NewCoins(sdk.NewCoin("token", sdk.NewInt(int64(msg.GetBet())))))
	if err != nil {
		return nil, err
	}

	// Create a new bet chart entry
	betChartEntry := types.BetChart{
		AccountName: msg.Creator,
		Bet:         msg.Bet,
	}

	// update the bet chart
	k.SetBetChart(ctx, betChartEntry)
	_ = ctx

	return &types.MsgEnterLotteryResponse{}, nil
}
