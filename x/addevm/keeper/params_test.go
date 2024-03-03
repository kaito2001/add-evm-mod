package keeper_test

import (
	"testing"

	testkeeper "add-evm/testutil/keeper"
	"add-evm/x/addevm/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AddevmKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
