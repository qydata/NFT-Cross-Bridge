package engine

import (
	"github.com/pkg/errors"
	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc20"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC20TxSentRegistration() {
	fromChainID := e.chainID()
	ss, err := e.queryERC20SwapPair(fromChainID, []erc20.SwapPairState{
		erc20.SwapPairStateCreationTxSent,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC20TxSentRegistration]: failed to query tx_sent SwapPairs"))
		return
	}

	var ids []string
	for _, s := range ss {
		confirmed, err := e.hasBlockConfirmed(s.CreateTxHash, s.DstChainID)
		if err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC20TxSentRegistration]: failed to check block confirmation for SwapPair %s", s.ID),
			)

			continue
		}

		if confirmed {
			ids = append(ids, s.ID)
		}
	}

	if len(ids) == 0 {
		return
	}

	if err := e.deps.DB.Model(&ss).Where("id in ?", ids).Updates(map[string]interface{}{
		"state":     erc20.SwapPairStateCreationTxConfirmed,
		"available": true,
	}).Error; err != nil {
		util.Logger.Error(
			errors.Wrapf(err, "[Engine.manageERC20TxSentRegistration]: failed to update state '%s'", erc20.SwapPairStateCreationTxConfirmed),
		)
	}

	for _, s := range ss {
		util.Logger.Infof("[Engine.manageERC20TxSentRegistration]: updated SwapPair %s state to '%s'", s.ID, s.State)
	}
}
