// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eneskzlcn/softdare/internal/home (interfaces: HomeService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	home "github.com/eneskzlcn/softdare/internal/home"
	gomock "github.com/golang/mock/gomock"
)

// MockHomeService is a mock of HomeService interface.
type MockHomeService struct {
	ctrl     *gomock.Controller
	recorder *MockHomeServiceMockRecorder
}

// MockHomeServiceMockRecorder is the mock recorder for MockHomeService.
type MockHomeServiceMockRecorder struct {
	mock *MockHomeService
}

// NewMockHomeService creates a new mock instance.
func NewMockHomeService(ctrl *gomock.Controller) *MockHomeService {
	mock := &MockHomeService{ctrl: ctrl}
	mock.recorder = &MockHomeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHomeService) EXPECT() *MockHomeServiceMockRecorder {
	return m.recorder
}

// GetPosts mocks base method.
func (m *MockHomeService) GetPosts(arg0 context.Context) ([]home.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", arg0)
	ret0, _ := ret[0].([]home.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockHomeServiceMockRecorder) GetPosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockHomeService)(nil).GetPosts), arg0)
}
