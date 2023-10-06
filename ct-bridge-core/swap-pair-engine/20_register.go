package engine

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc20"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

// queryERC20SwapPair queries SwapPair this engine is responsible
func (e *Engine) queryERC20SwapPair(fromChainID string, states []erc20.SwapPairState) ([]*erc20.SwapPair, error) {
	// TODO: check the index
	var ss []*erc20.SwapPair
	err := e.deps.DB.Where(
		"state in ? and src_chain_id = ?",
		states,
		fromChainID,
	).Order(
		"register_height asc",
	).Limit(
		querySwapPairLimit,
	).Find(&ss).Error
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.queryERC20SwapPair]: failed to query SwapPair")
	}

	return ss, nil
}

// filterERC20ConfirmedRegisterEvents checks block confirmation of the chain this engine is responsible
func (e *Engine) filterERC20ConfirmedRegisterEvents(ss []*erc20.SwapPair) (events []*erc20.SwapPair, err error) {
	for _, s := range ss {
		confirmed, err := e.hasBlockConfirmed(s.RegisterTxHash, e.chainID())
		if err != nil {
			util.Logger.Warning(errors.Wrap(err, "[Engine.filterERC20ConfirmedRegisterEvents]: failed to check block confirmation"))
			continue
		}
		if confirmed {
			events = append(events, s)
		}
	}

	return events, nil
}

func (e *Engine) generateERC20TxHash(s *erc20.SwapPair) (string, error) {
	request, err := e.sendERC20CreatePairRequest(s, true)
	if err != nil {
		return "", errors.Wrap(err, "[Engine.generateERC20TxHash]: failed to dry run sending a request")
	}

	return request.Hash().String(), nil
}

// sendERC20CreatePairRequest sends transaction to create a swap pair on destination chain
func (e *Engine) sendERC20CreatePairRequest(s *erc20.SwapPair, dryRun bool) (*types.Transaction, error) {
	dstChainID := s.DstChainID
	dstChainIDInt := util.StrToBigInt(dstChainID)
	if _, ok := e.deps.Client[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC20CreatePairRequest]: client for chain id %s is not supported", dstChainID)
	}
	if _, ok := e.deps.ERC20SwapAgent[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC20CreatePairRequest]: swap agent for chain id %s is not supported", dstChainID)
	}

	ctx := context.Background()
	txOpts, err := util.TxOpts(ctx, e.deps.Client[dstChainID], e.conf.PrivateKey, dstChainIDInt)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC20CreatePairRequest]: failed to create tx opts")
	}

	txOpts.NoSend = dryRun

	tx, err := e.deps.ERC20SwapAgent[dstChainID].CreateSwapPair(
		txOpts,
		common.HexToHash(s.RegisterTxHash),
		common.HexToAddress(s.SrcTokenAddr),
		util.StrToBigInt(s.SrcChainID),
		s.NAME,
		s.SYMBOL,
		s.DECIMALS,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC20CreatePairRequest]: failed to send swap pair creation tx")
	}

	return tx, nil
}
