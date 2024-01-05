// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/hypersdk/state (interfaces: View)

// Package state is a generated GoMock package.
package state

import (
	context "context"
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	merkledb "github.com/ava-labs/avalanchego/x/merkledb"
	gomock "go.uber.org/mock/gomock"
)

// MockView is a mock of View interface.
type MockView struct {
	ctrl     *gomock.Controller
	recorder *MockViewMockRecorder
}

// MockViewMockRecorder is the mock recorder for MockView.
type MockViewMockRecorder struct {
	mock *MockView
}

// NewMockView creates a new mock instance.
func NewMockView(ctrl *gomock.Controller) *MockView {
	mock := &MockView{ctrl: ctrl}
	mock.recorder = &MockViewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockView) EXPECT() *MockViewMockRecorder {
	return m.recorder
}

// GetMerkleRoot mocks base method.
func (m *MockView) GetMerkleRoot(arg0 context.Context) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMerkleRoot", arg0)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMerkleRoot indicates an expected call of GetMerkleRoot.
func (mr *MockViewMockRecorder) GetMerkleRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMerkleRoot", reflect.TypeOf((*MockView)(nil).GetMerkleRoot), arg0)
}

// GetValue mocks base method.
func (m *MockView) GetValue(arg0 context.Context, arg1 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValue indicates an expected call of GetValue.
func (mr *MockViewMockRecorder) GetValue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockView)(nil).GetValue), arg0, arg1)
}

// NewView mocks base method.
func (m *MockView) NewView(arg0 context.Context, arg1 merkledb.ViewChanges) (merkledb.TrieView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewView", arg0, arg1)
	ret0, _ := ret[0].(merkledb.TrieView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewView indicates an expected call of NewView.
func (mr *MockViewMockRecorder) NewView(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewView", reflect.TypeOf((*MockView)(nil).NewView), arg0, arg1)
}
