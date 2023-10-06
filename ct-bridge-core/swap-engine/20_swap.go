package engine

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc20"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
	"gorm.io/gorm"
)

func (e *Engine) generateERC20TxHash(s *erc20.Swap) (string, error) {
	request, err := e.sendERC20FillSwapRequest(s, true)
	if err != nil {
		return "", errors.Wrap(err, "[Engine.generateERC20TxHash]: failed to dry run sending a request")
	}

	return request.Hash().String(), nil
}

// sendERC20FillSwapRequest sends transaction to fill a swap on destination chain
func (e *Engine) sendERC20FillSwapRequest(s *erc20.Swap, dryRun bool) (*types.Transaction, error) {
	dstChainID := s.DstChainID
	dstChainIDInt := util.StrToBigInt(dstChainID)
	if _, ok := e.deps.Client[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC20FillSwapRequest]: client for chain id %s is not supported", dstChainID)
	}
	if _, ok := e.deps.ERC20SwapAgent[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC20FillSwapRequest]: swap agent for chain id %s is not supported", dstChainID)
	}

	ctx := context.Background()
	txOpts, err := util.TxOpts(ctx, e.deps.Client[dstChainID], e.conf.PrivateKey, dstChainIDInt)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC20FillSwapRequest]: failed to create tx opts")
	}

	tokenAddr := s.SrcTokenAddr
	tokenChainID := s.SrcChainID
	if s.SwapDirection == erc20.SwapDirectionBackward {
		tokenAddr = s.DstTokenAddr
		tokenChainID = s.SrcChainID
	}

	var amount = s.Amount

	txOpts.NoSend = dryRun
	tx, err := e.deps.ERC20SwapAgent[dstChainID].Fill(
		txOpts,
		common.HexToHash(s.RequestTxHash),
		common.HexToAddress(tokenAddr),
		common.HexToAddress(s.Recipient),
		util.StrToBigInt(tokenChainID),
		util.StrToBigInt(amount),
	)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC20FillSwapRequest]: failed to send swap pair creation tx")
	}

	return tx, nil
}

// queryERC20Swap queries Swap this engine is responsible
func (e *Engine) queryERC20Swap(fromChainID string, states []erc20.SwapState) ([]*erc20.Swap, error) {
	// TODO: check the index
	var ss []*erc20.Swap
	err := e.deps.DB.Where(
		"state in ? and src_chain_id = ?",
		states,
		fromChainID,
	).Order(
		"request_height asc",
	).Limit(
		querySwapLimit,
	).Find(&ss).Error
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.queryERC20Swap]: failed to query Swap")
	}

	return ss, nil
}

// fillERC20RequiredInfo fills swap destination tokens
func (e *Engine) fillERC20RequiredInfo(ss []*erc20.Swap) error {
	for _, s := range ss {
		if s.IsRequiredInfoValid() {
			continue
		}

		s.RequestTrackRetry += 1

		var skip bool
		var err error
		if s.SwapDirection == erc20.SwapDirectionForward {
			skip, err = e.fillERC20Forward(s)
		} else {
			skip, err = e.fillERC20Backward(s)
		}

		if skip {
			continue
		}
		if err != nil {
			return errors.Wrap(err, "[Engine.fillERC20RequiredInfo]: failed to fill required information")
		}
	}

	return nil
}

func (e *Engine) fillERC20Forward(s *erc20.Swap) (skip bool, err error) {
	// TODO: check db index
	var sp erc20.SwapPair
	err = e.deps.DB.Where(
		"src_token_addr = ? and src_chain_id = ? and dst_chain_id = ? and available = ?",
		s.SrcTokenAddr,
		s.SrcChainID,
		s.DstChainID,
		true,
	).First(&sp).Error

	if err == gorm.ErrRecordNotFound {
		return true, nil
	}
	if err != nil {
		return false, errors.Wrap(err, "[Engine.fillERC20Forward]: failed to query available Swaps")
	}

	s.SetRequiredInfo(sp.DstTokenAddr)

	return false, nil
}

func (e *Engine) fillERC20Backward(s *erc20.Swap) (skip bool, err error) {
	// TODO: check db index
	var sp erc20.SwapPair
	err = e.deps.DB.Where(
		"dst_token_addr = ? and dst_chain_id = ? and src_chain_id = ? and available = ?",
		s.SrcTokenAddr,
		s.SrcChainID,
		s.DstChainID,
		true,
	).First(&sp).Error

	if err == gorm.ErrRecordNotFound {
		return true, nil
	}
	if err != nil {
		return false, errors.Wrap(err, "[Engine.fillERC20Backward]: failed to query available Swaps")
	}

	s.SetRequiredInfo(sp.SrcTokenAddr)

	return false, nil
}

func (e *Engine) separateERC20SwapEvents(ss []*erc20.Swap) (pass []*erc20.Swap, pending []*erc20.Swap, rejected []*erc20.Swap) {
	for _, s := range ss {
		if !s.IsRequiredInfoValid() {
			if s.RequestTrackRetry > e.conf.MaxTrackRetry {
				rejected = append(rejected, s)
			} else {
				pending = append(pending, s)
			}

			continue
		}

		pass = append(pass, s)
	}

	return
}

// filterERC20ConfirmedSwapEvents checks block confirmation of the chain this engine is responsible
func (e *Engine) filterERC20ConfirmedSwapEvents(ss []*erc20.Swap) (events []*erc20.Swap, err error) {
	for _, s := range ss {
		confirmed, err := e.hasBlockConfirmed(s.RequestTxHash, e.chainID())
		if err != nil {
			util.Logger.Warning(errors.Wrap(err, "[Engine.filterERC20ConfirmedSwapEvents]: failed to check block confirmation"))
			continue
		}
		if confirmed {
			events = append(events, s)
		}
	}

	return events, nil
}
