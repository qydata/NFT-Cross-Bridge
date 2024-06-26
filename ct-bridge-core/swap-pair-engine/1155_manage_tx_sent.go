package engine

import (
	"github.com/pkg/errors"
	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc1155"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC1155TxSentRegistration() {
	fromChainID := e.chainID()
	ss, err := e.queryERC1155SwapPair(fromChainID, []erc1155.SwapPairState{
		erc1155.SwapPairStateCreationTxSent,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC1155TxSentRegistration]: failed to query tx_sent SwapPairs"))
		return
	}

	var ids []string
	for _, s := range ss {
		confirmed, err := e.hasBlockConfirmed(s.CreateTxHash, s.DstChainID)
		if err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC1155TxSentRegistration]: failed to check block confirmation for SwapPair %s", s.ID),
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
		"state":     erc1155.SwapPairStateCreationTxConfirmed,
		"available": true,
	}).Error; err != nil {
		util.Logger.Error(
			errors.Wrapf(err, "[Engine.manageERC1155TxSentRegistration]: failed to update state '%s'", erc1155.SwapPairStateCreationTxConfirmed),
		)
	}

	for _, s := range ss {
		util.Logger.Infof("[Engine.manageERC1155TxSentRegistration]: updated SwapPair %s state to '%s'", s.ID, s.State)
	}
}
