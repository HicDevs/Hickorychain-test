package keeper_test

import (
	"context"
	"testing"

	keepertest "Hickory/testutil/keeper"
	"Hickory/x/hickory/keeper"
	"Hickory/x/hickory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HickoryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
