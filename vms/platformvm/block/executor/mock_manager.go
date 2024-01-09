// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/vms/platformvm/block/executor (interfaces: Manager)

// Package executor is a generated GoMock package.
package executor

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	snowman "github.com/ava-labs/avalanchego/snow/consensus/snowman"
	block "github.com/ava-labs/avalanchego/vms/platformvm/block"
	state "github.com/ava-labs/avalanchego/vms/platformvm/state"
	txs "github.com/ava-labs/avalanchego/vms/platformvm/txs"
	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// GetBlock mocks base method.
func (m *MockManager) GetBlock(arg0 ids.ID) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockManagerMockRecorder) GetBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockManager)(nil).GetBlock), arg0)
}

// GetState mocks base method.
func (m *MockManager) GetState(arg0 ids.ID) (state.Chain, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState", arg0)
	ret0, _ := ret[0].(state.Chain)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetState indicates an expected call of GetState.
func (mr *MockManagerMockRecorder) GetState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockManager)(nil).GetState), arg0)
}

// GetStatelessBlock mocks base method.
func (m *MockManager) GetStatelessBlock(arg0 ids.ID) (block.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatelessBlock", arg0)
	ret0, _ := ret[0].(block.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatelessBlock indicates an expected call of GetStatelessBlock.
func (mr *MockManagerMockRecorder) GetStatelessBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatelessBlock", reflect.TypeOf((*MockManager)(nil).GetStatelessBlock), arg0)
}

// LastAccepted mocks base method.
func (m *MockManager) LastAccepted() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastAccepted")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// LastAccepted indicates an expected call of LastAccepted.
func (mr *MockManagerMockRecorder) LastAccepted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastAccepted", reflect.TypeOf((*MockManager)(nil).LastAccepted))
}

// NewBlock mocks base method.
func (m *MockManager) NewBlock(arg0 block.Block) snowman.Block {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBlock", arg0)
	ret0, _ := ret[0].(snowman.Block)
	return ret0
}

// NewBlock indicates an expected call of NewBlock.
func (mr *MockManagerMockRecorder) NewBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBlock", reflect.TypeOf((*MockManager)(nil).NewBlock), arg0)
}

// Preferred mocks base method.
func (m *MockManager) Preferred() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Preferred")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// Preferred indicates an expected call of Preferred.
func (mr *MockManagerMockRecorder) Preferred() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Preferred", reflect.TypeOf((*MockManager)(nil).Preferred))
}

// SetPreference mocks base method.
func (m *MockManager) SetPreference(arg0 ids.ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPreference", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetPreference indicates an expected call of SetPreference.
func (mr *MockManagerMockRecorder) SetPreference(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPreference", reflect.TypeOf((*MockManager)(nil).SetPreference), arg0)
}

// VerifyTx mocks base method.
func (m *MockManager) VerifyTx(arg0 *txs.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyTx", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyTx indicates an expected call of VerifyTx.
func (mr *MockManagerMockRecorder) VerifyTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyTx", reflect.TypeOf((*MockManager)(nil).VerifyTx), arg0)
}
