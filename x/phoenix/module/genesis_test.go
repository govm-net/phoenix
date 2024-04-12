package phoenix_test

import (
	"testing"

	keepertest "github.com/govm-net/phoenix/testutil/keeper"
	"github.com/govm-net/phoenix/testutil/nullify"
	phoenix "github.com/govm-net/phoenix/x/phoenix/module"
	"github.com/govm-net/phoenix/x/phoenix/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		VirtualBlockList: []types.VirtualBlock{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VirtualBlockCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PhoenixKeeper(t)
	phoenix.InitGenesis(ctx, k, genesisState)
	got := phoenix.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.VirtualBlockList, got.VirtualBlockList)
	require.Equal(t, genesisState.VirtualBlockCount, got.VirtualBlockCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
