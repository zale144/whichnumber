package whichnumber_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zale144/whichnumber/testutil/keeper"
	"github.com/zale144/whichnumber/testutil/nullify"
	"github.com/zale144/whichnumber/x/whichnumber"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WhichnumberKeeper(t)
	whichnumber.InitGenesis(ctx, *k, genesisState)
	got := whichnumber.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
