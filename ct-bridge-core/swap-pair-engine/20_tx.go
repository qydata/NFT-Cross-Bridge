package engine

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

func (e *Engine) retrieveERC20DstTokenAddr(height uint64, fromTokenAddr, registerTxHash, chainID string) (string, error) {
	if _, ok := e.deps.Client[chainID]; !ok {
		return "", errors.Errorf("[Engine.retrieveERC20DstTokenAddr]: client for chain id %s is not supported", chainID)
	}

	ctx := context.Background()
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	txHash := [32]byte(common.HexToHash(registerTxHash))
	iter, err := e.deps.ERC20SwapAgent[chainID].FilterSwapPairCreated(&opts, [][32]byte{txHash}, nil, nil)
	if err != nil {
		return "", errors.Wrap(err, "[Engine.retrieveERC20DstTokenAddr]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Engine.retrieveERC20DstTokenAddr]: failed to close iterator, %s", err.Error())
		}
	}()

	for iter.Next() {
		if iter.Event.FromTokenAddr.String() == fromTokenAddr {
			return iter.Event.MirroredTokenAddr.String(), nil
		}
	}

	if err := iter.Error(); err != nil {
		return "", errors.Wrap(err, "[Recorder.retrieveERC20DstTokenAddr]: failed to iterate events")
	}

	return "", nil
}
