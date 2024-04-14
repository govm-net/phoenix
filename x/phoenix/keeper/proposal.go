package keeper

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/mempool"
	"github.com/govm-net/phoenix/x/phoenix/types"
)

const virtualBlockInterval = 10

type ProposalHandler struct {
	defaultHandler *baseapp.DefaultProposalHandler
	logger         log.Logger
	keeper         Keeper
}

func NewProposalHandler(logger log.Logger, keeper Keeper, mp mempool.Mempool, txVerifier baseapp.ProposalTxVerifier) *ProposalHandler {
	return &ProposalHandler{
		defaultHandler: baseapp.NewDefaultProposalHandler(mp, txVerifier),
		logger:         logger,
		keeper:         keeper,
	}
}

func (h *ProposalHandler) PrepareProposal() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		h.logger.Debug("PrepareProposal", "height", req.Height)

		dfHandler := h.defaultHandler.PrepareProposalHandler()
		resp, err := dfHandler(ctx, req)
		if err != nil {
			return nil, err
		}

		last, found := h.keeper.GetLastVirtualBlock(ctx)
		if !found {
			return nil, errors.New("fail to get vitrual block")
		}
		now := req.Time.Unix()
		// one vitrual block per 10s
		vbTime := last.Time + virtualBlockInterval
		if now <= vbTime {
			fmt.Println("skip", now, vbTime, now-vbTime)
			return resp, nil
		}
		var vb types.VirtualBlock

		vb.Time = vbTime

		vb.Previous, err = last.Hash()
		if err != nil {
			return nil, err
		}
		vb.Header = h.keeper.GetLastHeader(ctx)
		// todo other attributes

		bz, err := json.Marshal(vb)
		if err != nil {
			return nil, errors.New("failed to encode injected vote extension tx")
		}

		resp.Txs = append(resp.Txs, bz)

		return resp, nil
	}
}

func (h *ProposalHandler) ProcessProposal() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		h.logger.Debug("ProcessProposal", "height", req.Height)

		last, found := h.keeper.GetLastVirtualBlock(ctx)
		if !found {
			h.logger.Error("not found virtual block", "height", req.Height)
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, errors.New("fail to get vitrual block")
		}

		if len(req.Txs) == 0 {
			if last.Time+virtualBlockInterval < req.Time.Unix() {
				h.logger.Error("not found tx, virtual block time out", "height", req.Height)
				return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
			}
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
		}

		var vb types.VirtualBlock
		if err := json.Unmarshal(req.Txs[0], &vb); err != nil {
			if last.Time+virtualBlockInterval < req.Time.Unix() {
				return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
			}
			h.logger.Error("failed to decode injected vote extension tx", "err", err)
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
		}
		if vb.Time != last.Time+virtualBlockInterval {
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
		}
		if !bytes.Equal(vb.Header, h.keeper.GetLastHeader(ctx)) {
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
		}
		if h.keeper.GetParams(ctx).ShardId == 1 {
			if len(vb.Parent) != 0 {
				return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
			}
		}
		// todo check other attributes

		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}
}

func (h *ProposalHandler) PreBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	h.logger.Debug("PreBlocker", "height", req.Height)

	if len(req.Txs) == 0 {
		return h.checkBlock(ctx, req)
	}

	var vb types.VirtualBlock
	if err := json.Unmarshal(req.Txs[0], &vb); err != nil {
		return h.checkBlock(ctx, req)
	}
	if !bytes.Equal(vb.Header, h.keeper.GetLastHeader(ctx)) {
		return nil, errors.New("wrong header of virtual block")
	}
	// todo check other attributes

	h.keeper.AppendVirtualBlock(ctx, vb)

	return h.checkBlock(ctx, req)
}

func (h *ProposalHandler) checkBlock(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	res := &sdk.ResponsePreBlock{}
	vb, found := h.keeper.GetLastVirtualBlock(ctx)
	if !found {
		fmt.Println("not found the virtual block")
		return nil, errors.New("not found vitrual block")
	}
	now := req.Time.Unix()
	if now > vb.Time+virtualBlockInterval {
		return nil, errors.New("lost vitrual block")
	}
	return res, nil
}

func (h *ProposalHandler) ExtendVoteHandler() sdk.ExtendVoteHandler {
	return func(ctx sdk.Context, req *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
		return &abci.ResponseExtendVote{}, nil
	}
}

func (h *ProposalHandler) VerifyVoteExtensionHandler() sdk.VerifyVoteExtensionHandler {
	return func(ctx sdk.Context, req *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
	}
}

// GetVirtualBlockCount get the total number of virtualBlock
func (k Keeper) GetLastHeader(ctx context.Context) []byte {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LastHeaderKey)
	return store.Get(byteKey)
}

// SetVirtualBlockCount set the total number of virtualBlock
func (k Keeper) SetLastHeader(ctx context.Context, header []byte) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LastHeaderKey)
	store.Set(byteKey, header)
}
