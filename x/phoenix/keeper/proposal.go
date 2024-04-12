package keeper

import (
	"encoding/json"
	"errors"
	"fmt"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto/tmhash"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/govm-net/phoenix/x/phoenix/types"
)

// StakeWeightedPrices defines the structure a proposer should use to calculate
// and submit the stake-weighted prices for a given set of supported currency
// pairs, in addition to the vote extensions used to calculate them. This is so
// validators can verify the proposer's calculations.
type StakeWeightedPrices struct {
	StakeWeightedPrices map[string]math.LegacyDec
	ExtendedCommitInfo  abci.ExtendedCommitInfo
}

type ProposalHandler struct {
	logger log.Logger
	keeper Keeper
}

func NewProposalHandler(logger log.Logger, keeper Keeper) *ProposalHandler {
	return &ProposalHandler{
		logger: logger,
		keeper: keeper,
	}
}

func getHeaderHash(ctx sdk.Context) []byte {
	h := ctx.BlockHeader()
	hd, err := h.Marshal()
	if err != nil {
		return nil
	}
	return tmhash.Sum(hd)
}

func (h *ProposalHandler) PrepareProposal() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		proposalTxs := req.Txs
		fmt.Println("--PrepareProposal--", req.Height)

		bid := ctx.BlockHeight()

		if bid == 1 {
			d, _ := json.MarshalIndent(ctx.BlockHeader(), "", "  ")
			fmt.Println("first header:", string(d))
			var vb types.VirtualBlock

			vb.Header = getHeaderHash(ctx)
			vb.Time = req.Time.Unix()
			if h.keeper.GetParams(ctx).ShardId > 1 {
				vb.Parent = ctx.BlockHeader().AppHash
			}
			bz, err := json.Marshal(vb)
			if err != nil {
				return nil, errors.New("failed to encode injected vote extension tx")
			}

			proposalTxs = append(proposalTxs, bz)

			return &abci.ResponsePrepareProposal{
				Txs: proposalTxs,
			}, nil
		}

		lid := h.keeper.GetVirtualBlockCount(ctx)
		last, found := h.keeper.GetVirtualBlock(ctx, lid-1)
		if !found {
			fmt.Println("not found the virtual block", lid-1)
			return &abci.ResponsePrepareProposal{
				Txs: proposalTxs,
			}, nil
		}
		now := ctx.BlockTime().Unix()
		// one vitrual block per 10s
		vbTime := last.Time + 10
		if now < vbTime {
			fmt.Println("skip", now, vbTime, now-vbTime)
			return &abci.ResponsePrepareProposal{
				Txs: proposalTxs,
			}, nil
		}
		var vb types.VirtualBlock

		vb.Time = vbTime

		hd, err := last.Marshal()
		if err != nil {
			return nil, err
		}
		vb.Previous = tmhash.Sum(hd)

		bz, err := json.Marshal(vb)
		if err != nil {
			return nil, errors.New("failed to encode injected vote extension tx")
		}

		proposalTxs = append(proposalTxs, bz)

		return &abci.ResponsePrepareProposal{
			Txs: proposalTxs,
		}, nil
	}
}

func (h *ProposalHandler) ProcessProposal() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		fmt.Println("--ProcessProposal--", req.Height)
		if len(req.Txs) == 0 {
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
		}

		var vb types.VirtualBlock
		if err := json.Unmarshal(req.Txs[0], &vb); err != nil {
			h.logger.Error("failed to decode injected vote extension tx", "err", err)
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
		}

		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}
}

func (h *ProposalHandler) PreBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	fmt.Println("--PreBlocker--", req.Height)
	res := &sdk.ResponsePreBlock{}
	var err error
	defer func() {
		res, err = h.checkBlock(ctx, req)
	}()
	if len(req.Txs) == 0 {
		return res, err
	}

	var vb types.VirtualBlock
	if err := json.Unmarshal(req.Txs[0], &vb); err != nil {
		return res, err
	}

	h.keeper.AppendVirtualBlock(ctx, vb)

	return res, err
}

func (h *ProposalHandler) checkBlock(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	res := &sdk.ResponsePreBlock{}
	lid := h.keeper.GetVirtualBlockCount(ctx)
	vb, found := h.keeper.GetVirtualBlock(ctx, lid-1)
	if !found {
		fmt.Println("not found the virtual block", lid-1)
		return nil, errors.New("not found vitrual block")
	}
	now := req.Time.Unix()
	if now > vb.Time+10 {
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
