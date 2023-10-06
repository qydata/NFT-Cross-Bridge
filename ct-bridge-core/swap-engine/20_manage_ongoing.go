package engine

import (
	"github.com/pkg/errors"

	"github.com/qydata/ct-evm-compatible-bridge-core/model/erc20"
	"github.com/qydata/ct-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC20OngoingRequest() {
	fromChainID := e.chainID()
	ss, err := e.queryERC20Swap(fromChainID, []erc20.SwapState{
		erc20.SwapStateRequestOngoing,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC20OngoingRequest]: failed to query onging Swaps"))
		return
	}

	// Fill required information without updating to DB
	if err := e.fillERC20RequiredInfo(ss); err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC20OngoingRequest]: failed to fill destination"))
		return
	}

	// Separate ready Swaps, pending Swaps, and rejected Swaps
	ss, pp, rr := e.separateERC20SwapEvents(ss)
	for _, r := range rr {
		r.State = erc20.SwapStateRequestRejected
		if err := e.deps.DB.Save(&r).Error; err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC20OngoingRequest]: failed to update Swap %s to state '%s'", r.ID, r.State),
			)
		}
	}
	for _, p := range pp {
		if err := e.deps.DB.Save(&p).Error; err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC20OngoingRequest]: failed to update Swap %s", p.ID),
			)
		}
	}

	ss, err = e.filterERC20ConfirmedSwapEvents(ss)
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC20OngoingRequest]: failed to filter confirmed Swaps"))
		return
	}

	for _, s := range ss {
		s.State = erc20.SwapStateRequestConfirmed
		if err := e.deps.DB.Save(&s).Error; err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC20OngoingRequest]: failed to update Swap %s to state '%s'", s.ID, s.State),
			)
		}
	}
}
