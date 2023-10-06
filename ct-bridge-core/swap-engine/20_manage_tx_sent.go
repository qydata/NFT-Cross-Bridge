package engine

import (
	"github.com/pkg/errors"

	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc20"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC20TxSentSwap() {
	fromChainID := e.chainID()
	ss, err := e.queryERC20Swap(fromChainID, []erc20.SwapState{
		erc20.SwapStateFillTxSent,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC20TxSentSwap]: failed to query tx_sent Swaps"))
		return
	}

	var ids []string
	for _, s := range ss {
		confirmed, err := e.hasBlockConfirmed(s.FillTxHash, s.DstChainID)
		if err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC20TxSentSwap]: failed to check block confirmation for Swap %s", s.ID),
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
		"state": erc20.SwapStateFillTxConfirmed,
	}).Error; err != nil {
		util.Logger.Error(
			errors.Wrapf(err, "[Engine.manageERC20TxSentSwap]: failed to update state '%s'", erc20.SwapStateFillTxConfirmed),
		)
	}

	for _, s := range ss {
		util.Logger.Infof("[Engine.manageERC20TxSentSwap]: updated Swap %s state to '%s'", s.ID, s.State)
	}
}
