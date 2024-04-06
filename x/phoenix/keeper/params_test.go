package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/govm-net/phoenix/testutil/keeper"
	"github.com/govm-net/phoenix/x/phoenix/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.PhoenixKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
