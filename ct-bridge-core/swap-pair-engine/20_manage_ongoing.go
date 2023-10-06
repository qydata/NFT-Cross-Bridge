package engine

import (
	"github.com/pkg/errors"

	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc20"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC20OngoingRegistration() {
	fromChainID := e.chainID()
	ss, err := e.queryERC20SwapPair(fromChainID, []erc20.SwapPairState{
		erc20.SwapPairStateRegistrationOngoing,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC20OngoingRegistration]: failed to query onging SwapPairs"))
		return
	}

	ss, err = e.filterERC20ConfirmedRegisterEvents(ss)
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC20OngoingRegistration]: failed to filter confirmed SwapPairs"))
		return
	}

	if len(ss) == 0 {
		return
	}

	ids := make([]string, len(ss))
	for idx, s := range ss {
		ids[idx] = s.ID
	}

	if err := e.deps.DB.Model(&ss).Where("id in ?", ids).Updates(map[string]interface{}{
		"state": erc20.SwapPairStateRegistrationConfirmed,
	}).Error; err != nil {
		util.Logger.Error(
			errors.Wrapf(err, "[Engine.manageERC20OngoingRegistration]: failed to update state '%s'", erc20.SwapPairStateRegistrationConfirmed),
		)
	}

	for _, s := range ss {
		util.Logger.Infof("[Engine.manageERC20OngoingRegistration]: updated SwapPair %s state to '%s'", s.ID, s.State)
	}
}
