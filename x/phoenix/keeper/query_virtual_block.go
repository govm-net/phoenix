package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/govm-net/phoenix/x/phoenix/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VirtualBlockAll(ctx context.Context, req *types.QueryAllVirtualBlockRequest) (*types.QueryAllVirtualBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virtualBlocks []types.VirtualBlock

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	virtualBlockStore := prefix.NewStore(store, types.KeyPrefix(types.VirtualBlockKey))

	pageRes, err := query.Paginate(virtualBlockStore, req.Pagination, func(key []byte, value []byte) error {
		var virtualBlock types.VirtualBlock
		if err := k.cdc.Unmarshal(value, &virtualBlock); err != nil {
			return err
		}

		virtualBlocks = append(virtualBlocks, virtualBlock)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVirtualBlockResponse{VirtualBlock: virtualBlocks, Pagination: pageRes}, nil
}

func (k Keeper) VirtualBlock(ctx context.Context, req *types.QueryGetVirtualBlockRequest) (*types.QueryGetVirtualBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	virtualBlock, found := k.GetVirtualBlock(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVirtualBlockResponse{VirtualBlock: virtualBlock}, nil
}
