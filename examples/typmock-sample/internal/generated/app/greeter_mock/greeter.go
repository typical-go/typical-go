// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/typical-go/typical-go/examples/typmock-sample/internal/app/greeter (interfaces: Greeter)

// Package greeter_mock is a generated GoMock package.
package greeter_mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGreeter is a mock of Greeter interface
type MockGreeter struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterMockRecorder
}

// MockGreeterMockRecorder is the mock recorder for MockGreeter
type MockGreeterMockRecorder struct {
	mock *MockGreeter
}

// NewMockGreeter creates a new mock instance
func NewMockGreeter(ctrl *gomock.Controller) *MockGreeter {
	mock := &MockGreeter{ctrl: ctrl}
	mock.recorder = &MockGreeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGreeter) EXPECT() *MockGreeterMockRecorder {
	return m.recorder
}

// Greet mocks base method
func (m *MockGreeter) Greet() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Greet")
	ret0, _ := ret[0].(string)
	return ret0
}

// Greet indicates an expected call of Greet
func (mr *MockGreeterMockRecorder) Greet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Greet", reflect.TypeOf((*MockGreeter)(nil).Greet))
}
