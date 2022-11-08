package hickory_test

import (
	"testing"

	keepertest "Hickory/testutil/keeper"
	"Hickory/testutil/nullify"
	"Hickory/x/hickory"
	"Hickory/x/hickory/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HickoryKeeper(t)
	hickory.InitGenesis(ctx, *k, genesisState)
	got := hickory.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
