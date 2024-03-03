package keeper_test

import (
	"context"
	"testing"

	keepertest "add-evm/testutil/keeper"
	"add-evm/x/addevm/keeper"
	"add-evm/x/addevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AddevmKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
