// Code generated by MockGen. DO NOT EDIT.
// Source: app/application/comment_app.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	params "github.com/jumpei00/board/backend/app/application/params"
	domain "github.com/jumpei00/board/backend/app/domain"
)

// MockCommentApplication is a mock of CommentApplication interface.
type MockCommentApplication struct {
	ctrl     *gomock.Controller
	recorder *MockCommentApplicationMockRecorder
}

// MockCommentApplicationMockRecorder is the mock recorder for MockCommentApplication.
type MockCommentApplicationMockRecorder struct {
	mock *MockCommentApplication
}

// NewMockCommentApplication creates a new mock instance.
func NewMockCommentApplication(ctrl *gomock.Controller) *MockCommentApplication {
	mock := &MockCommentApplication{ctrl: ctrl}
	mock.recorder = &MockCommentApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentApplication) EXPECT() *MockCommentApplicationMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockCommentApplication) CreateComment(param *params.CreateCommentAppLayerParam) (*[]domain.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", param)
	ret0, _ := ret[0].(*[]domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockCommentApplicationMockRecorder) CreateComment(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockCommentApplication)(nil).CreateComment), param)
}

// DeleteComment mocks base method.
func (m *MockCommentApplication) DeleteComment(param *params.DeleteCommentAppLayerParam) (*[]domain.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", param)
	ret0, _ := ret[0].(*[]domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockCommentApplicationMockRecorder) DeleteComment(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockCommentApplication)(nil).DeleteComment), param)
}

// EditComment mocks base method.
func (m *MockCommentApplication) EditComment(param *params.EditCommentAppLayerParam) (*[]domain.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditComment", param)
	ret0, _ := ret[0].(*[]domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditComment indicates an expected call of EditComment.
func (mr *MockCommentApplicationMockRecorder) EditComment(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditComment", reflect.TypeOf((*MockCommentApplication)(nil).EditComment), param)
}

// GetAllByThreadKey mocks base method.
func (m *MockCommentApplication) GetAllByThreadKey(threadKey string) (*[]domain.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByThreadKey", threadKey)
	ret0, _ := ret[0].(*[]domain.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByThreadKey indicates an expected call of GetAllByThreadKey.
func (mr *MockCommentApplicationMockRecorder) GetAllByThreadKey(threadKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByThreadKey", reflect.TypeOf((*MockCommentApplication)(nil).GetAllByThreadKey), threadKey)
}
