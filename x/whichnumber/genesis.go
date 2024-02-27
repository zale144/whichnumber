package whichnumber

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zale144/whichnumber/x/whichnumber/keeper"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set if defined
	k.SetStoredSystemInfo(ctx, genState.SystemInfo)
	// Set all the storedGame
	for _, elem := range genState.StoredGameList {
		k.SetStoredGame(ctx, elem)
	}

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetStoredParams(ctx)

	// Get all systemInfo
	systemInfo, found := k.GetStoredSystemInfo(ctx)
	if found {
		genesis.SystemInfo = systemInfo
	}
	genesis.StoredGameList = k.GetStoredGames(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
