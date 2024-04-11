package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/govm-net/phoenix/testutil/keeper"
	"github.com/govm-net/phoenix/testutil/nullify"
	"github.com/govm-net/phoenix/x/phoenix/keeper"
	"github.com/govm-net/phoenix/x/phoenix/types"
	"github.com/stretchr/testify/require"
)

func createNVirtualBlock(keeper keeper.Keeper, ctx context.Context, n int) []types.VirtualBlock {
	items := make([]types.VirtualBlock, n)
	for i := range items {
		items[i].Id = keeper.AppendVirtualBlock(ctx, items[i])
	}
	return items
}

func TestVirtualBlockGet(t *testing.T) {
	keeper, ctx := keepertest.PhoenixKeeper(t)
	items := createNVirtualBlock(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetVirtualBlock(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestVirtualBlockRemove(t *testing.T) {
	keeper, ctx := keepertest.PhoenixKeeper(t)
	items := createNVirtualBlock(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVirtualBlock(ctx, item.Id)
		_, found := keeper.GetVirtualBlock(ctx, item.Id)
		require.False(t, found)
	}
}

func TestVirtualBlockGetAll(t *testing.T) {
	keeper, ctx := keepertest.PhoenixKeeper(t)
	items := createNVirtualBlock(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVirtualBlock(ctx)),
	)
}

func TestVirtualBlockCount(t *testing.T) {
	keeper, ctx := keepertest.PhoenixKeeper(t)
	items := createNVirtualBlock(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetVirtualBlockCount(ctx))
}
