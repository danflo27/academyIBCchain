package send_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ibcFinal/testutil/keeper"
	"ibcFinal/testutil/nullify"
	"ibcFinal/x/send"
	"ibcFinal/x/send/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SendKeeper(t)
	send.InitGenesis(ctx, *k, genesisState)
	got := send.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
