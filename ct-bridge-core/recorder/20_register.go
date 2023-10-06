package recorder

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"

	"github.com/qydata/ct-evm-compatible-bridge-core/model/block"
	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc20"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

func (r *Recorder) recordERC20RegisterTx(tx *gorm.DB, b *block.Log) error {
	ctx, cancel := context.WithTimeout(context.Background(), registerFilterLogsTimeout)
	defer cancel()

	height := uint64(b.Height)
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	iter, err := r.deps.ERC20SwapAgent[r.ChainID()].FilterSwapPairRegister(&opts, nil, nil)
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC20RegisterTx]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Recorder.recordERC20RegisterTx]: failed to close iterator, %s", err.Error())
		}
	}()

	var ss []erc20.SwapPair
	for iter.Next() {
		s := erc20.SwapPair{
			SrcChainID:   r.ChainID(),
			DstChainID:   iter.Event.ToChainId.String(),
			SrcTokenAddr: iter.Event.TokenAddress.String(),
			DstTokenAddr: "",
			Sponsor:      iter.Event.Sponsor.String(),
			Available:    false,
			Signature:    "",
			NAME:         "",
			SYMBOL:       "",
			DECIMALS:     0,

			State: erc20.SwapPairStateRegistrationOngoing,

			RegisterTxHash:     iter.Event.Raw.TxHash.String(),
			RegisterHeight:     int64(iter.Event.Raw.BlockNumber),
			RegisterBlockHash:  iter.Event.Raw.BlockHash.String(),
			RegisterBlockLog:   nil,
			RegisterBlockLogID: &b.ID,

			CreateTxHash:     "",
			CreateHeight:     math.MaxInt64,
			CreateBlockHash:  "",
			CreateBlockLog:   nil,
			CreateBlockLogID: nil,
		}

		name, err := r.retrieveERC20Name(s.SrcTokenAddr)
		if err != nil {
			return errors.Wrap(err, "[Recorder.recordERC20RegisterTx]: failed to get name")
		}

		symbol, err1 := r.retrieveERC20Symbol(s.SrcTokenAddr)
		if err1 != nil {
			return errors.Wrap(err1, "[Recorder.recordERC20RegisterTx]: failed to get symbol")
		}

		decimals, err2 := r.retrieveERC20Decimals(s.SrcTokenAddr)
		if err2 != nil {
			return errors.Wrap(err1, "[Recorder.recordERC20RegisterTx]: failed to get decimals")
		}

		s.NAME = name
		s.SYMBOL = symbol
		s.DECIMALS = decimals
		ss = append(ss, s)
	}

	if err := iter.Error(); err != nil {
		return errors.Wrap(err, "[Recorder.recordERC20RegisterTx]: failed to iterate events")
	}

	err = tx.Clauses(
		clause.OnConflict{DoNothing: true},
	).Omit(
		clause.Associations,
	).CreateInBatches(
		&ss, 100,
	).Error
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC20RegisterTx]: failed to bulk create")
	}

	return nil
}

func (r *Recorder) retrieveERC20Name(tokenAddr string) (string, error) {
	token, ok := r.deps.ERC20Token[r.ChainID()]
	if !ok {
		return "", errors.Errorf("[Recorder.retrieveERC20Metadata]: unsupported chain id %s", r.ChainID())
	}

	opts := &bind.CallOpts{
		Pending: true,
	}
	name, err := token.NAME(opts, tokenAddr)
	if err != nil {
		return "", errors.Wrap(err, "[Recorder.retrieveERC20Metadata]: failed to retrieve name")
	}

	return name, nil
}

func (r *Recorder) retrieveERC20Symbol(tokenAddr string) (string, error) {
	token, ok := r.deps.ERC20Token[r.ChainID()]
	if !ok {
		return "", errors.Errorf("[Recorder.retrieveERC20Metadata]: unsupported chain id %s", r.ChainID())
	}

	opts := &bind.CallOpts{
		Pending: true,
	}
	symbol, err := token.SYMBOL(opts, tokenAddr)
	if err != nil {
		return "", errors.Wrap(err, "[Recorder.retrieveERC20Symbol]: failed to retrieve symbol")
	}

	return symbol, nil
}

func (r *Recorder) retrieveERC20Decimals(tokenAddr string) (uint8, error) {
	token, ok := r.deps.ERC20Token[r.ChainID()]
	if !ok {
		return 0, errors.Errorf("[Recorder.retrieveERC20Metadata]: unsupported chain id %s", r.ChainID())
	}

	opts := &bind.CallOpts{
		Pending: true,
	}
	decimals, err := token.DECIMALS(opts, tokenAddr)
	if err != nil {
		return 0, errors.Wrap(err, "[Recorder.retrieveERC20Symbol]: failed to retrieve decimals")
	}

	return decimals, nil
}
