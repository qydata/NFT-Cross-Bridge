package engine

import (
	"math/big"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/qydata/ct-evm-compatible-bridge-core/model/block"
	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc20"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC20TxCreatedRegistration() {
	fromChainID := e.chainID()
	ss, err := e.queryERC20SwapPair(fromChainID, []erc20.SwapPairState{
		erc20.SwapPairStateCreationTxCreated,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC20TxCreatedRegistration]: failed to query tx_created SwapPairs"))
		return
	}

	for _, s := range ss {
		ethTx, isPending, err := e.retrieveTx(s.CreateTxHash, s.DstChainID)
		if err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC20TxCreatedRegistration]: failed to get SwapPair creation tx %s", s.CreateTxHash),
			)

			continue
		}
		if isPending {
			util.Logger.Infof("[Engine.manageERC20TxCreatedRegistration]: the tx %s is pending in mempools, skip", s.CreateTxHash)
			continue
		}

		receipt, err := e.retrieveTxReceipt(s.CreateTxHash, s.DstChainID)
		if err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC20TxCreatedRegistration]: failed to get SwapPair creation receipt for tx %s", s.CreateTxHash),
			)

			continue
		}

		if ethTx == nil {
			util.Logger.Infof("[Engine.manageERC20TxCreatedRegistration]: the tx is not found while cheking tx %s", s.CreateTxHash)
		}

		if receipt == nil {
			util.Logger.Infof("[Engine.manageERC20TxCreatedRegistration]: the receipt is not found while cheking tx %s", s.CreateTxHash)
		}

		if ethTx == nil || receipt == nil {
			s.CreateTrackRetry += 1
			if err := e.deps.DB.Save(s).Error; err != nil {
				util.Logger.Error(
					errors.Wrapf(err, "[Engine.manageERC20TxCreatedRegistration]: failed to increase create track retry counter %s", s.ID),
				)

				continue
			}

			if s.CreateTrackRetry > e.conf.MaxTrackRetry {
				s.State = erc20.SwapPairStateCreationTxMissing
				s.MessageLog = "[Engine.manageERC20TxCreatedRegistration]: tx is missing"
				if err := e.deps.DB.Save(s).Error; err != nil {
					util.Logger.Error(
						errors.Wrapf(err, "[Engine.manageERC20TxCreatedRegistration]: failed to update SwapPair %s to '%s' state", s.ID, s.State),
					)

					continue
				}
			}

			continue
		}

		var b block.Log
		err = e.deps.DB.Where(
			"chain_id = ? and block_hash = ?",
			s.DstChainID,
			receipt.BlockHash.String(),
		).Select(
			"id",
		).First(
			&b,
		).Error
		if err == gorm.ErrRecordNotFound {
			util.Logger.Infof("[Engine.manageERC20TxCreatedRegistration]: wait for the system to catch up the block %s in chain id %s", receipt.BlockHash.String(), e.chainID())

			continue
		}
		if err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC20TxCreatedRegistration]: failed to update SwapPair %s to '%s' state", s.ID, s.State),
			)

			continue
		}

		createBlockHeight := receipt.BlockNumber.Int64()
		dstTokenAddr, err := e.retrieveERC20DstTokenAddr(uint64(createBlockHeight), s.SrcTokenAddr, s.RegisterTxHash, s.DstChainID)
		if err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC20TxCreatedRegistration]: failed to get destination token address for SwapPair %s", s.ID),
			)

			continue
		}

		if dstTokenAddr == "" {
			s.State = erc20.SwapPairStateCreationTxFailed
			s.MessageLog = "[Engine.manageERC20TxCreatedRegistration]: destination token address was not found"
			if err := e.deps.DB.Save(s).Error; err != nil {
				util.Logger.Error(
					errors.Wrapf(err, "[Engine.manageERC20TxCreatedRegistration]: failed to update SwapPair %s to '%s' state", s.ID, s.State),
				)

				continue
			}
		}

		gasPrice := big.NewInt(0)
		gasPrice.SetString(ethTx.GasPrice().String(), 10)
		s.DstTokenAddr = dstTokenAddr
		s.CreateGasPrice = gasPrice.String()
		s.CreateConsumedFeeAmount = big.NewInt(1).Mul(gasPrice, big.NewInt(s.CreateGasUsed)).String()
		s.CreateGasUsed = int64(receipt.GasUsed)
		s.CreateHeight = createBlockHeight
		s.CreateBlockHash = receipt.BlockHash.String()
		s.CreateBlockLogID = &b.ID
		s.State = erc20.SwapPairStateCreationTxSent
		if err := e.deps.DB.Save(s).Error; err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC20TxCreatedRegistration]: failed to update SwapPair %s basic info", s.ID),
			)

			continue
		}

		util.Logger.Infof("[Engine.manageERC20TxCreatedRegistration]: updated SwapPair %s after sending out with tx hash %s", s.ID, s.CreateTxHash)
	}
}
