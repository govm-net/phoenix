package phoenix

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/govm-net/phoenix/x/phoenix/keeper"
	"github.com/govm-net/phoenix/x/phoenix/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the virtualBlock
	for _, elem := range genState.VirtualBlockList {
		k.SetVirtualBlock(ctx, elem)
	}

	// Set virtualBlock count
	k.SetVirtualBlockCount(ctx, genState.VirtualBlockCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.VirtualBlockList = k.GetAllVirtualBlock(ctx)
	genesis.VirtualBlockCount = k.GetVirtualBlockCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
