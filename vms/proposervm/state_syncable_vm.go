// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/vms/proposervm/summary"
)

var (
	errUnknownLastSummaryBlockID = errors.New("could not retrieve blockID associated with last summary")
	errBadLastSummaryBlock       = errors.New("could not parse last summary block")
)

func (vm *VM) StateSyncEnabled() (bool, error) {
	if vm.coreStateSyncVM == nil {
		return false, nil
	}

	return vm.coreStateSyncVM.StateSyncEnabled()
}

func (vm *VM) StateSyncGetOngoingSummary() (common.Summary, error) {
	if vm.coreStateSyncVM == nil {
		return nil, common.ErrStateSyncableVMNotImplemented
	}

	coreSummary, err := vm.coreStateSyncVM.StateSyncGetOngoingSummary()
	if err != nil {
		return nil, err // including common.ErrNoStateSyncOngoing case
	}

	// retrieve ProBlkID
	proBlkID, err := vm.GetBlockIDAtHeight(coreSummary.Key())
	if err != nil {
		// this should never happen, it's proVM being out of sync with coreVM
		vm.ctx.Log.Warn("core summary unknown to proposer VM. Block height index missing: %s", err)
		return nil, common.ErrUnknownStateSummary
	}

	return summary.New(proBlkID, coreSummary)
}

func (vm *VM) StateSyncGetLastSummary() (common.Summary, error) {
	if vm.coreStateSyncVM == nil {
		return nil, common.ErrStateSyncableVMNotImplemented
	}

	// Extract core last state summary
	coreSummary, err := vm.coreStateSyncVM.StateSyncGetLastSummary()
	if err != nil {
		return nil, err // including common.ErrUnknownStateSummary case
	}

	// retrieve ProBlkID
	proBlkID, err := vm.GetBlockIDAtHeight(coreSummary.Key())
	if err != nil {
		// this should never happen, it's proVM being out of sync with coreVM
		vm.ctx.Log.Warn("core summary unknown to proposer VM. Block height index missing: %s", err)
		return nil, common.ErrUnknownStateSummary
	}

	return summary.New(proBlkID, coreSummary)
}

// Note: it's important that StateSyncParseSummary do not use any index or state
// to allow summaries being parsed also by freshly started node with no previous state.
func (vm *VM) StateSyncParseSummary(summaryBytes []byte) (common.Summary, error) {
	if vm.coreStateSyncVM == nil {
		return nil, common.ErrStateSyncableVMNotImplemented
	}

	proContent, err := summary.Parse(summaryBytes)
	if err != nil {
		return nil, err
	}

	coreSummary, err := vm.coreStateSyncVM.StateSyncParseSummary(proContent.CoreSummaryBytes())
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal coreSummaryContent due to: %w", err)
	}

	return summary.New(proContent.ProposerBlockID(), coreSummary)
}

func (vm *VM) StateSyncGetSummary(key uint64) (common.Summary, error) {
	if vm.coreStateSyncVM == nil {
		return nil, common.ErrStateSyncableVMNotImplemented
	}

	coreSummary, err := vm.coreStateSyncVM.StateSyncGetSummary(key)
	if err != nil {
		return nil, err // including common.ErrUnknownStateSummary case
	}

	// retrieve ProBlkID
	proBlkID, err := vm.GetBlockIDAtHeight(coreSummary.Key())
	if err != nil {
		// this should never happen, it's proVM being out of sync with coreVM
		vm.ctx.Log.Warn("core summary unknown to proposer VM. Block height index missing: %s", err)
		return nil, common.ErrUnknownStateSummary
	}

	return summary.New(proBlkID, coreSummary)
}

func (vm *VM) StateSync(accepted []common.Summary) error {
	if vm.coreStateSyncVM == nil {
		return common.ErrStateSyncableVMNotImplemented
	}

	coreSummaries := make([]common.Summary, 0, len(accepted))
	for _, s := range accepted {
		proContent, err := summary.Parse(s.Bytes())
		if err != nil {
			return err
		}

		coreSummary, err := vm.coreStateSyncVM.StateSyncParseSummary(proContent.CoreSummaryBytes())
		if err != nil {
			return fmt.Errorf("could not parse coreSummaryContent due to: %w", err)
		}

		coreSummaries = append(coreSummaries, coreSummary)

		// Following state sync introduction, we update height -> blockID index
		// with summaries content in order to support resuming state sync in case
		// of shutdown. The height index allows to retrieve the proposerBlkID
		// of any state sync passed down to coreVM, so that the proposerVM state summary
		// information of any coreVM summary can be rebuilt and pass to the engine, even
		// following a shutdown.
		// Note that we won't download all the blocks associated with state summaries,
		// so proposerVM may not not all the full blocks indexed into height index. Same
		// is true for coreVM.
		if err := vm.updateHeightIndex(s.Key(), proContent.ProposerBlockID()); err != nil {
			return err
		}
	}

	if err := vm.db.Commit(); err != nil {
		return nil
	}

	return vm.coreStateSyncVM.StateSync(coreSummaries)
}

func (vm *VM) StateSyncGetResult() (ids.ID, uint64, error) {
	if vm.coreStateSyncVM == nil {
		return ids.Empty, 0, common.ErrStateSyncableVMNotImplemented
	}

	_, height, err := vm.coreStateSyncVM.StateSyncGetResult()
	if err != nil {
		return ids.Empty, 0, err
	}
	proBlkID, err := vm.GetBlockIDAtHeight(height)
	if err != nil {
		return ids.Empty, 0, errUnknownLastSummaryBlockID
	}
	return proBlkID, height, nil
}

func (vm *VM) StateSyncSetLastSummaryBlockID(blkID ids.ID) error {
	if vm.coreStateSyncVM == nil {
		return common.ErrStateSyncableVMNotImplemented
	}

	var (
		blk Block
		err error
	)

	blk, err = vm.getPostForkBlock(blkID)
	if err != nil {
		blk, err = vm.getPreForkBlock(blkID)
		if err != nil {
			return errBadLastSummaryBlock
		}
	}

	if err := blk.acceptOuterBlk(); err != nil {
		return err
	}

	return vm.coreStateSyncVM.StateSyncSetLastSummaryBlockID(blk.getInnerBlk().ID())
}
