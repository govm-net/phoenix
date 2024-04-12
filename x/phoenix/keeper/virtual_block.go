package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/govm-net/phoenix/x/phoenix/types"
)

// GetVirtualBlockCount get the total number of virtualBlock
func (k Keeper) GetVirtualBlockCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.VirtualBlockCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVirtualBlockCount set the total number of virtualBlock
func (k Keeper) SetVirtualBlockCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.VirtualBlockCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVirtualBlock appends a virtualBlock in the store with a new id and update the count
func (k Keeper) AppendVirtualBlock(
	ctx context.Context,
	virtualBlock types.VirtualBlock,
) uint64 {
	// Create the virtualBlock
	count := k.GetVirtualBlockCount(ctx)

	// Set the ID of the appended value
	virtualBlock.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualBlockKey))
	appendedValue := k.cdc.MustMarshal(&virtualBlock)
	store.Set(GetVirtualBlockIDBytes(virtualBlock.Id), appendedValue)

	// Update virtualBlock count
	k.SetVirtualBlockCount(ctx, count+1)

	return count
}

// SetVirtualBlock set a specific virtualBlock in the store
func (k Keeper) SetVirtualBlock(ctx context.Context, virtualBlock types.VirtualBlock) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualBlockKey))
	b := k.cdc.MustMarshal(&virtualBlock)
	store.Set(GetVirtualBlockIDBytes(virtualBlock.Id), b)
}

// GetVirtualBlock returns a virtualBlock from its id
func (k Keeper) GetVirtualBlock(ctx context.Context, id uint64) (val types.VirtualBlock, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualBlockKey))
	b := store.Get(GetVirtualBlockIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVirtualBlock removes a virtualBlock from the store
func (k Keeper) RemoveVirtualBlock(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualBlockKey))
	store.Delete(GetVirtualBlockIDBytes(id))
}

// GetAllVirtualBlock returns all virtualBlock
func (k Keeper) GetAllVirtualBlock(ctx context.Context) (list []types.VirtualBlock) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VirtualBlockKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VirtualBlock
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVirtualBlockIDBytes returns the byte representation of the ID
func GetVirtualBlockIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.VirtualBlockKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
