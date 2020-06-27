// Code generated by MockGen. DO NOT EDIT.
// Source: ./receiver/device.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDeviceReceiverInterface is a mock of DeviceReceiverInterface interface
type MockDeviceReceiverInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceReceiverInterfaceMockRecorder
}

// MockDeviceReceiverInterfaceMockRecorder is the mock recorder for MockDeviceReceiverInterface
type MockDeviceReceiverInterfaceMockRecorder struct {
	mock *MockDeviceReceiverInterface
}

// NewMockDeviceReceiverInterface creates a new mock instance
func NewMockDeviceReceiverInterface(ctrl *gomock.Controller) *MockDeviceReceiverInterface {
	mock := &MockDeviceReceiverInterface{ctrl: ctrl}
	mock.recorder = &MockDeviceReceiverInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceReceiverInterface) EXPECT() *MockDeviceReceiverInterfaceMockRecorder {
	return m.recorder
}

// SetA mocks base method
func (m *MockDeviceReceiverInterface) SetA(a int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetA", a)
}

// SetA indicates an expected call of SetA
func (mr *MockDeviceReceiverInterfaceMockRecorder) SetA(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetA", reflect.TypeOf((*MockDeviceReceiverInterface)(nil).SetA), a)
}

// SetB mocks base method
func (m *MockDeviceReceiverInterface) SetB(b int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetB", b)
}

// SetB indicates an expected call of SetB
func (mr *MockDeviceReceiverInterfaceMockRecorder) SetB(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetB", reflect.TypeOf((*MockDeviceReceiverInterface)(nil).SetB), b)
}

// SetN mocks base method
func (m *MockDeviceReceiverInterface) SetN(n int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetN", n)
}

// SetN indicates an expected call of SetN
func (mr *MockDeviceReceiverInterfaceMockRecorder) SetN(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetN", reflect.TypeOf((*MockDeviceReceiverInterface)(nil).SetN), n)
}

// Sum mocks base method
func (m *MockDeviceReceiverInterface) Sum() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sum")
	ret0, _ := ret[0].(int)
	return ret0
}

// Sum indicates an expected call of Sum
func (mr *MockDeviceReceiverInterfaceMockRecorder) Sum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sum", reflect.TypeOf((*MockDeviceReceiverInterface)(nil).Sum))
}

// Multiply mocks base method
func (m *MockDeviceReceiverInterface) Multiply() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Multiply")
	ret0, _ := ret[0].(int)
	return ret0
}

// Multiply indicates an expected call of Multiply
func (mr *MockDeviceReceiverInterfaceMockRecorder) Multiply() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Multiply", reflect.TypeOf((*MockDeviceReceiverInterface)(nil).Multiply))
}

// Prime mocks base method
func (m *MockDeviceReceiverInterface) Prime() []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prime")
	ret0, _ := ret[0].([]int)
	return ret0
}

// Prime indicates an expected call of Prime
func (mr *MockDeviceReceiverInterfaceMockRecorder) Prime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prime", reflect.TypeOf((*MockDeviceReceiverInterface)(nil).Prime))
}

// Fibonacci mocks base method
func (m *MockDeviceReceiverInterface) Fibonacci() []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fibonacci")
	ret0, _ := ret[0].([]int)
	return ret0
}

// Fibonacci indicates an expected call of Fibonacci
func (mr *MockDeviceReceiverInterfaceMockRecorder) Fibonacci() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fibonacci", reflect.TypeOf((*MockDeviceReceiverInterface)(nil).Fibonacci))
}