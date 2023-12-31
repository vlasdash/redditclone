// Code generated by MockGen. DO NOT EDIT.
// Source: comment.go

// Package comment is a generated GoMock package.
package mock

import (
	"github.com/vlasdash/redditclone/internal/comment"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCommentRepo is a mock of CommentRepo interface.
type MockCommentRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepoMockRecorder
}

// MockCommentRepoMockRecorder is the mock recorder for MockCommentRepo.
type MockCommentRepoMockRecorder struct {
	mock *MockCommentRepo
}

// NewMockCommentRepo creates a new mock instance.
func NewMockCommentRepo(ctrl *gomock.Controller) *MockCommentRepo {
	mock := &MockCommentRepo{ctrl: ctrl}
	mock.recorder = &MockCommentRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentRepo) EXPECT() *MockCommentRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCommentRepo) Add(userID uint, body string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", userID, body)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockCommentRepoMockRecorder) Add(userID, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCommentRepo)(nil).Add), userID, body)
}

// Delete mocks base method.
func (m *MockCommentRepo) Delete(id string, userID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCommentRepoMockRecorder) Delete(id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommentRepo)(nil).Delete), id, userID)
}

// GetByID mocks base method.
func (m *MockCommentRepo) GetByID(id string) (*comment.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*comment.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockCommentRepoMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCommentRepo)(nil).GetByID), id)
}
