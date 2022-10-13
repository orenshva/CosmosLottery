package keeper

import (
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
)

var _ types.QueryServer = Keeper{}
