// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/providers/tinkerbell/hardware/json.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	hardware "github.com/aws/eks-anywhere/pkg/providers/tinkerbell/hardware"
	gomock "github.com/golang/mock/gomock"
)

// MockTinkerbellHardwareJsonFactory is a mock of TinkerbellHardwareJsonFactory interface.
type MockTinkerbellHardwareJsonFactory struct {
	ctrl     *gomock.Controller
	recorder *MockTinkerbellHardwareJsonFactoryMockRecorder
}

// MockTinkerbellHardwareJsonFactoryMockRecorder is the mock recorder for MockTinkerbellHardwareJsonFactory.
type MockTinkerbellHardwareJsonFactoryMockRecorder struct {
	mock *MockTinkerbellHardwareJsonFactory
}

// NewMockTinkerbellHardwareJsonFactory creates a new mock instance.
func NewMockTinkerbellHardwareJsonFactory(ctrl *gomock.Controller) *MockTinkerbellHardwareJsonFactory {
	mock := &MockTinkerbellHardwareJsonFactory{ctrl: ctrl}
	mock.recorder = &MockTinkerbellHardwareJsonFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTinkerbellHardwareJsonFactory) EXPECT() *MockTinkerbellHardwareJsonFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTinkerbellHardwareJsonFactory) Create(name string) (*hardware.TinkerbellHardwareJson, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name)
	ret0, _ := ret[0].(*hardware.TinkerbellHardwareJson)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTinkerbellHardwareJsonFactoryMockRecorder) Create(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTinkerbellHardwareJsonFactory)(nil).Create), name)
}

// MockTinkerbellHardwarePusher is a mock of TinkerbellHardwarePusher interface.
type MockTinkerbellHardwarePusher struct {
	ctrl     *gomock.Controller
	recorder *MockTinkerbellHardwarePusherMockRecorder
}

// MockTinkerbellHardwarePusherMockRecorder is the mock recorder for MockTinkerbellHardwarePusher.
type MockTinkerbellHardwarePusherMockRecorder struct {
	mock *MockTinkerbellHardwarePusher
}

// NewMockTinkerbellHardwarePusher creates a new mock instance.
func NewMockTinkerbellHardwarePusher(ctrl *gomock.Controller) *MockTinkerbellHardwarePusher {
	mock := &MockTinkerbellHardwarePusher{ctrl: ctrl}
	mock.recorder = &MockTinkerbellHardwarePusherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTinkerbellHardwarePusher) EXPECT() *MockTinkerbellHardwarePusherMockRecorder {
	return m.recorder
}

// PushHardware mocks base method.
func (m *MockTinkerbellHardwarePusher) PushHardware(ctx context.Context, hardware []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushHardware", ctx, hardware)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushHardware indicates an expected call of PushHardware.
func (mr *MockTinkerbellHardwarePusherMockRecorder) PushHardware(ctx, hardware interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushHardware", reflect.TypeOf((*MockTinkerbellHardwarePusher)(nil).PushHardware), ctx, hardware)
}
