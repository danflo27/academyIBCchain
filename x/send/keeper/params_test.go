package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ibcFinal/testutil/keeper"
	"ibcFinal/x/send/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SendKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
