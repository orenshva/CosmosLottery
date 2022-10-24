package cosmoslottery

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/client/cli"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/keeper"
	"github.com/orenshva/CosmosLottery/x/cosmoslottery/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface for the capability module.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the capability module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns the capability module's default genesis state.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis performs genesis state validation for the capability module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterRESTRoutes registers the capability module's REST service handlers.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the capability module's root tx command.
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the capability module's root query command.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface for the capability module.
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// Name returns the capability module's name.
func (am AppModule) Name() string {
	return am.AppModuleBasic.Name()
}

// Route returns the capability module's message routing key.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
}

// QuerierRoute returns the capability module's query routing key.
func (AppModule) QuerierRoute() string { return types.QuerierRoute }

// LegacyQuerierHandler returns the capability module's Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the capability module's invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the capability module's genesis initialization It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the capability module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion implements ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 2 }

// BeginBlock executes all ABCI BeginBlock logic respective to the capability module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock executes all ABCI EndBlock logic respective to the capability module. It
// returns no validator updates.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	// get TxCounter
	// initTxCounter := types.TxCounter{Count: 0}
	// am.keeper.SetTxCounter(ctx, initTxCounter)
	currentTxCount, found := am.keeper.GetTxCounter(ctx)
	if found == false {
		panic("No TX counter. The lottery can't operate without one")
	}

	// get fee counter
	feeCounter, found := am.keeper.GetFeeCounter(ctx)
	if found == false {
		panic("No fee counter, can't calculate award")
	}

	// get the bet chart
	completeBetChart := am.keeper.GetAllBetChart(ctx)

	// check TxCounter
	if currentTxCount.Count >= 10 {
		concatBets := ""
		var totalBetsInThisRound uint64 = 0

		/* iterate over all the bet chart entries and then:
		1. save the total amount of bets in this round
		2. append the placed bets to get the value to be hashed
		3. remove the bet entry from the store (so it won't persist to the next lottery) */
		for _, betChartEntry := range completeBetChart {
			tempBet := betChartEntry.GetBet()
			totalBetsInThisRound += tempBet
			concatBets += strconv.Itoa(int(tempBet))
			am.keeper.RemoveBetChart(ctx, betChartEntry.GetAccountName())
		}

		// hash the result
		hashFunc := sha256.New()
		hashFunc.Write([]byte(concatBets))
		hashResult := binary.BigEndian.Uint64(hashFunc.Sum(nil))

		// decide the winner
		winnerIndex := (hashResult ^ 0xFFFF) % currentTxCount.Count

		// get the winner's address
		winnerBetChartEntry := completeBetChart[int(winnerIndex)]
		winnerAccountAddress, err := sdk.AccAddressFromBech32(winnerBetChartEntry.GetAccountName())
		if err != nil {
			fmt.Printf("Couldn't find winner's address. No payment performed")
			goto funcend
		}

		// determine whether the winner had the largest/lowest bet
		var userHasLargestBet bool = true
		var userHasLowestBet bool = true

		for index, betChartEntry := range completeBetChart {
			if index == int(winnerIndex) {
				continue
			}

			if betChartEntry.GetBet() > winnerBetChartEntry.GetBet() {
				userHasLargestBet = false
			} else if betChartEntry.GetBet() < winnerBetChartEntry.GetBet() {
				userHasLowestBet = false
			}
		}

		// get the amount of money in the lottery pool
		moduleAddress, err := sdk.AccAddressFromBech32(types.ModuleName)
		if err != nil {
			panic("No module account (no lottery pool)")
		}
		lotteryPoolFunds := am.bankKeeper.GetBalance(ctx, moduleAddress, "token")

		if userHasLargestBet == true {
			// the winner has the highest bet
			payoutAmount := lotteryPoolFunds.Amount
			am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winnerAccountAddress, sdk.NewCoins(sdk.NewCoin("token", payoutAmount)))
		} else if userHasLargestBet == false && userHasLowestBet == false {
			// the winner has a mid-sized bet
			feeCount := feeCounter.GetCount()
			payoutAmount := totalBetsInThisRound - feeCount*types.LotteryFee.Amount.Uint64()
			am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winnerAccountAddress, sdk.NewCoins(sdk.NewCoin("token", sdk.NewInt(int64(payoutAmount)))))
		}

		// Zero out TxCounter
		currentTxCount.Count = 0
		am.keeper.SetTxCounter(ctx, currentTxCount)

		// Zero out FeeCounter
		feeCounter.Count = 0
		am.keeper.SetFeeCounter(ctx, feeCounter)
	}

funcend:
	return []abci.ValidatorUpdate{}
}
