// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eneskzlcn/softdare/internal/login (interfaces: LoginRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	login "github.com/eneskzlcn/softdare/internal/login"
	gomock "github.com/golang/mock/gomock"
)

// MockLoginRepository is a mock of LoginRepository interface.
type MockLoginRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLoginRepositoryMockRecorder
}

// MockLoginRepositoryMockRecorder is the mock recorder for MockLoginRepository.
type MockLoginRepositoryMockRecorder struct {
	mock *MockLoginRepository
}

// NewMockLoginRepository creates a new mock instance.
func NewMockLoginRepository(ctrl *gomock.Controller) *MockLoginRepository {
	mock := &MockLoginRepository{ctrl: ctrl}
	mock.recorder = &MockLoginRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginRepository) EXPECT() *MockLoginRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockLoginRepository) CreateUser(arg0 context.Context, arg1 login.CreateUserRequest) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockLoginRepositoryMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockLoginRepository)(nil).CreateUser), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockLoginRepository) GetUserByEmail(arg0 context.Context, arg1 string) (*login.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*login.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockLoginRepositoryMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockLoginRepository)(nil).GetUserByEmail), arg0, arg1)
}

// IsUserExistsByEmail mocks base method.
func (m *MockLoginRepository) IsUserExistsByEmail(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserExistsByEmail", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserExistsByEmail indicates an expected call of IsUserExistsByEmail.
func (mr *MockLoginRepositoryMockRecorder) IsUserExistsByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserExistsByEmail", reflect.TypeOf((*MockLoginRepository)(nil).IsUserExistsByEmail), arg0, arg1)
}

// IsUserExistsByUsername mocks base method.
func (m *MockLoginRepository) IsUserExistsByUsername(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserExistsByUsername", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserExistsByUsername indicates an expected call of IsUserExistsByUsername.
func (mr *MockLoginRepositoryMockRecorder) IsUserExistsByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserExistsByUsername", reflect.TypeOf((*MockLoginRepository)(nil).IsUserExistsByUsername), arg0, arg1)
}