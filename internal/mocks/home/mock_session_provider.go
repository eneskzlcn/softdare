// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eneskzlcn/softdare/internal/home (interfaces: SessionProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSessionProvider is a mock of SessionProvider interface.
type MockSessionProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSessionProviderMockRecorder
}

// MockSessionProviderMockRecorder is the mock recorder for MockSessionProvider.
type MockSessionProviderMockRecorder struct {
	mock *MockSessionProvider
}

// NewMockSessionProvider creates a new mock instance.
func NewMockSessionProvider(ctrl *gomock.Controller) *MockSessionProvider {
	mock := &MockSessionProvider{ctrl: ctrl}
	mock.recorder = &MockSessionProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionProvider) EXPECT() *MockSessionProviderMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockSessionProvider) Exists(arg0 *http.Request, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockSessionProviderMockRecorder) Exists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockSessionProvider)(nil).Exists), arg0, arg1)
}

// Get mocks base method.
func (m *MockSessionProvider) Get(arg0 *http.Request, arg1 string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockSessionProviderMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSessionProvider)(nil).Get), arg0, arg1)
}