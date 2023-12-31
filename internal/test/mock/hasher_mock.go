// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package user is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPasswordHasher is a mock of PasswordHasher interface.
type MockPasswordHasher struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordHasherMockRecorder
}

// MockPasswordHasherMockRecorder is the mock recorder for MockPasswordHasher.
type MockPasswordHasherMockRecorder struct {
	mock *MockPasswordHasher
}

// NewMockPasswordHasher creates a new mock instance.
func NewMockPasswordHasher(ctrl *gomock.Controller) *MockPasswordHasher {
	mock := &MockPasswordHasher{ctrl: ctrl}
	mock.recorder = &MockPasswordHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordHasher) EXPECT() *MockPasswordHasherMockRecorder {
	return m.recorder
}

// GetHashPassword mocks base method.
func (m *MockPasswordHasher) GetHashPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHashPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHashPassword indicates an expected call of GetHashPassword.
func (mr *MockPasswordHasherMockRecorder) GetHashPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashPassword", reflect.TypeOf((*MockPasswordHasher)(nil).GetHashPassword), password)
}

// IsPassword mocks base method.
func (m *MockPasswordHasher) IsPassword(hash, password string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPassword", hash, password)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPassword indicates an expected call of IsPassword.
func (mr *MockPasswordHasherMockRecorder) IsPassword(hash, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPassword", reflect.TypeOf((*MockPasswordHasher)(nil).IsPassword), hash, password)
}
