package engine

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

func (e *Engine) verifyERC20ForwardSwapFillEvent(height uint64, requestTxHash, chainID string) (bool, error) {
	if _, ok := e.deps.Client[chainID]; !ok {
		return false, errors.Errorf("[Engine.verifyERC20ForwardSwapFillEvent]: client for chain id %s is not supported", chainID)
	}

	ctx := context.Background()
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	txHash := [32]byte(common.HexToHash(requestTxHash))

	iter, err := e.deps.ERC20SwapAgent[chainID].FilterSwapFilled(&opts, [][32]byte{txHash}, nil, nil)
	if err != nil {
		return false, errors.Wrap(err, "[Engine.verifyERC20ForwardSwapFillEvent]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Engine.verifyERC20ForwardSwapFillEvent]: failed to close iterator, %s", err.Error())
		}
	}()

	return iter.Next(), nil
}

func (e *Engine) verifyERC20BackwardSwapFillEvent(height uint64, requestTxHash, chainID string) (bool, error) {
	if _, ok := e.deps.Client[chainID]; !ok {
		return false, errors.Errorf("[Engine.verifyERC20BackwardSwapFillEvent]: client for chain id %s is not supported", chainID)
	}

	ctx := context.Background()
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	txHash := [32]byte(common.HexToHash(requestTxHash))
	iter, err := e.deps.ERC20SwapAgent[chainID].FilterBackwardSwapFilled(&opts, [][32]byte{txHash}, nil, nil)
	if err != nil {
		return false, errors.Wrap(err, "[Engine.verifyERC20BackwardSwapFillEvent]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Engine.verifyERC20BackwardSwapFillEvent]: failed to close iterator, %s", err.Error())
		}
	}()

	return iter.Next(), nil
}
