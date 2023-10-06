package recorder

import (
	"context"
	"math"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/qydata/ct-evm-compatible-bridge-core/model/block"
	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc20"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

func (r *Recorder) recordERC20SwapTx(tx *gorm.DB, b *block.Log) error {
	ctx, cancel := context.WithTimeout(context.Background(), swapFilterLogsTimeout)
	defer cancel()

	height := uint64(b.Height)
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	iter, err := r.deps.ERC20SwapAgent[r.ChainID()].FilterSwapStarted(&opts, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC20SwapTx]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Recorder.recordERC20SwapTx]: failed to close iterator, %s", err.Error())
		}
	}()

	var ss []erc20.Swap
	for iter.Next() {
		amount := iter.Event.Amount

		s := erc20.Swap{
			SrcChainID:            r.ChainID(),
			DstChainID:            iter.Event.DstChainId.String(),
			SrcTokenAddr:          iter.Event.TokenAddr.String(),
			DstTokenAddr:          "",
			Sender:                iter.Event.Sender.String(),
			Recipient:             iter.Event.Recipient.String(),
			Amount:                amount.String(),
			Signature:             "",
			State:                 erc20.SwapStateRequestOngoing,
			SwapDirection:         erc20.SwapDirectionForward,
			RequestTxHash:         iter.Event.Raw.TxHash.String(),
			RequestHeight:         int64(iter.Event.Raw.BlockNumber),
			RequestBlockHash:      iter.Event.Raw.BlockHash.String(),
			RequestBlockLogID:     &b.ID,
			RequestBlockLog:       nil,
			RequestTrackRetry:     0,
			FillConsumedFeeAmount: "",
			FillGasPrice:          "",
			FillGasUsed:           0,
			FillHeight:            math.MaxInt64,
			FillTxHash:            "",
			FillTrackRetry:        0,
			FillBlockHash:         "",
			FillBlockLogID:        nil,
			FillBlockLog:          nil,
			MessageLog:            "",
		}

		ss = append(ss, s)
	}

	if err := iter.Error(); err != nil {
		return errors.Wrap(err, "[Recorder.recordERC20SwapTx]: failed to iterate events")
	}

	err = tx.Clauses(
		clause.OnConflict{DoNothing: true},
	).Omit(
		clause.Associations,
	).CreateInBatches(
		&ss, 100,
	).Error
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC20SwapTx]: failed to bulk create")
	}

	return nil
}

func (r *Recorder) recordERC20BackwardSwapTx(tx *gorm.DB, b *block.Log) error {
	ctx, cancel := context.WithTimeout(context.Background(), swapFilterLogsTimeout)
	defer cancel()

	height := uint64(b.Height)
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	iter, err := r.deps.ERC20SwapAgent[r.ChainID()].FilterBackwardSwapStarted(&opts, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC20BackwardSwapTx]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Recorder.recordERC20BackwardSwapTx]: failed to close iterator, %s", err.Error())
		}
	}()

	var ss []erc20.Swap
	for iter.Next() {
		amount := iter.Event.Amount

		if err != nil {
			return errors.Wrap(err, "[Recorder.recordERC20BackwardSwapTx]: failed to marshal amount")
		}

		s := erc20.Swap{
			SrcChainID:            r.ChainID(),
			DstChainID:            iter.Event.DstChainId.String(),
			SrcTokenAddr:          iter.Event.MirroredTokenAddr.String(),
			DstTokenAddr:          "",
			Sender:                iter.Event.Sender.String(),
			Recipient:             iter.Event.Recipient.String(),
			Amount:                amount.String(),
			Signature:             "",
			State:                 erc20.SwapStateRequestOngoing,
			SwapDirection:         erc20.SwapDirectionBackward,
			RequestTxHash:         iter.Event.Raw.TxHash.String(),
			RequestHeight:         int64(iter.Event.Raw.BlockNumber),
			RequestBlockHash:      iter.Event.Raw.BlockHash.String(),
			RequestBlockLogID:     &b.ID,
			RequestBlockLog:       nil,
			RequestTrackRetry:     0,
			FillConsumedFeeAmount: "",
			FillGasPrice:          "",
			FillGasUsed:           0,
			FillHeight:            math.MaxInt64,
			FillTxHash:            "",
			FillTrackRetry:        0,
			FillBlockHash:         "",
			FillBlockLogID:        nil,
			FillBlockLog:          nil,
			MessageLog:            "",
		}

		ss = append(ss, s)
	}

	if err := iter.Error(); err != nil {
		return errors.Wrap(err, "[Recorder.recordERC20BackwardSwapTx]: failed to iterate events")
	}

	err = tx.Clauses(
		clause.OnConflict{DoNothing: true},
	).Omit(
		clause.Associations,
	).CreateInBatches(
		&ss, 100,
	).Error
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC20BackwardSwapTx]: failed to bulk create")
	}

	return nil
}
