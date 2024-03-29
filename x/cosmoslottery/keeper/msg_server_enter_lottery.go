package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
)

func (k msgServer) EnterLottery(goCtx context.Context, msg *types.MsgEnterLottery) (*types.MsgEnterLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var userAlreadyBetted bool = false

	// Check that if the user isn't already participating in the lottery
	userBetChart, found := k.GetBetChart(ctx, msg.GetCreator())
	if found == true {
		userAlreadyBetted = true
	}

	// Prevent the block proposer from participating
	if hex.EncodeToString(ctx.BlockHeader().ProposerAddress) == msg.GetCreator() {
		return nil, sdkerrors.Wrapf(types.ErrUserIsBlockProposer, "%s", msg.GetCreator())
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

	// Check that the user has enough balance to pay the fee and place the bet
	minAmount := msg.GetLotteryFee() + msg.GetBet()
	if k.bankKeeper.HasBalance(ctx, bidderAcc, sdk.NewCoin("token", sdk.NewInt(int64(minAmount)))) == false {
		return nil, sdkerrors.Wrapf(types.ErrUserDoesNotHaveEnoughFunds, "%s", msg.GetCreator())
	}

	// Deduct the lottery fee
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidderAcc, types.ModuleName, sdk.NewCoins(types.LotteryFee))
	if err != nil {
		return nil, err
	}

	// update the fee counter
	currentFeeCount, found := k.GetFeeCounter(ctx)
	if found == false {
		panic("No fee counter. The lottery can't operate without one")
	}
	currentFeeCount.Count += 1
	k.SetFeeCounter(ctx, currentFeeCount)

	// Transfer the bet to the lottery pool
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, bidderAcc, types.ModuleName, sdk.NewCoins(sdk.NewCoin("token", sdk.NewInt(int64(msg.GetBet())))))
	if err != nil {
		return nil, err
	}

	if userAlreadyBetted == false {
		// Bet from new user

		// Create a new bet chart entry
		newBetChartEntry := types.BetChart{
			AccountName: msg.Creator,
			Bet:         msg.Bet,
		}

		// update the bet chart
		k.SetBetChart(ctx, newBetChartEntry)

		// update the TX counter
		currentTxCount, found := k.GetTxCounter(ctx)
		if found == false {
			panic("No TX counter. The lottery can't operate without one")
		}
		currentTxCount.Count += 1
		k.SetTxCounter(ctx, currentTxCount)
	} else {
		// Bet from old user (bet substitution)
		userBetChart.Bet = msg.GetBet()
		k.SetBetChart(ctx, userBetChart)
	}

	_ = ctx
	return &types.MsgEnterLotteryResponse{}, nil
}
