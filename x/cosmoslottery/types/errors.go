package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cosmoslottery module sentinel errors
var (
	ErrUserAlreadyBetted          = sdkerrors.Register(ModuleName, 1104, "player is already participating in the lottery")
	ErrBadLotteryFee              = sdkerrors.Register(ModuleName, 1105, "player did not pay the right amount of lottery fee")
	ErrBetTooSmall                = sdkerrors.Register(ModuleName, 1106, "player's bet is lower than the minimum bet")
	ErrBetTooBig                  = sdkerrors.Register(ModuleName, 1107, "player's bet is bigger than the maximum bet")
	ErrUserDoesNotHaveEnoughFunds = sdkerrors.Register(ModuleName, 1108, "player does not have enough funds to pay the lottery fee and the desired bet")
	ErrUserIsValidator            = sdkerrors.Register(ModuleName, 1109, "validators can't participate in the lottery")
)
