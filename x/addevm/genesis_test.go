package addevm_test

import (
	"testing"

	keepertest "add-evm/testutil/keeper"
	"add-evm/testutil/nullify"
	"add-evm/x/addevm"
	"add-evm/x/addevm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AddevmKeeper(t)
	addevm.InitGenesis(ctx, *k, genesisState)
	got := addevm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
