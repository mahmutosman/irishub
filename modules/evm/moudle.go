package evm

import (
	"encoding/json"
	"math/big"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	ethermint "github.com/evmos/ethermint/x/evm"
	"github.com/evmos/ethermint/x/evm/keeper"
	"github.com/evmos/ethermint/x/evm/types"

	iristypes "github.com/irisnet/irishub/types"
)

var (
	_ module.AppModule = AppModule{}
)

// ____________________________________________________________________________

// AppModule implements an application module for the evm module.
type AppModule struct {
	ethermint.AppModule
	eip155ChainID *big.Int
}

// NewAppModule creates a new AppModule object
func NewAppModule(k *keeper.Keeper, ak types.AccountKeeper, eip155ChainID *big.Int) AppModule {
	return AppModule{
		AppModule:     ethermint.NewAppModule(k, ak),
		eip155ChainID: eip155ChainID,
	}
}

// BeginBlock returns the begin block for the evm module.
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	ethChainID := iristypes.BuildEthChainID(ctx.ChainID(), am.eip155ChainID)
	am.AppModule.BeginBlock(ctx.WithChainID(ethChainID), req)
}

// InitGenesis performs genesis initialization for the evm module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	ethChainID := iristypes.BuildEthChainID(ctx.ChainID(), am.eip155ChainID)
	return am.AppModule.InitGenesis(ctx.WithChainID(ethChainID), cdc, data)
}
