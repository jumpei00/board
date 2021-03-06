// Code generated by MockGen. DO NOT EDIT.
// Source: app/application/thread_app.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	params "github.com/jumpei00/board/backend/app/application/params"
	domain "github.com/jumpei00/board/backend/app/domain"
)

// MockThreadApplication is a mock of ThreadApplication interface.
type MockThreadApplication struct {
	ctrl     *gomock.Controller
	recorder *MockThreadApplicationMockRecorder
}

// MockThreadApplicationMockRecorder is the mock recorder for MockThreadApplication.
type MockThreadApplicationMockRecorder struct {
	mock *MockThreadApplication
}

// NewMockThreadApplication creates a new mock instance.
func NewMockThreadApplication(ctrl *gomock.Controller) *MockThreadApplication {
	mock := &MockThreadApplication{ctrl: ctrl}
	mock.recorder = &MockThreadApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThreadApplication) EXPECT() *MockThreadApplicationMockRecorder {
	return m.recorder
}

// CreateThread mocks base method.
func (m *MockThreadApplication) CreateThread(param *params.CreateThreadAppLayerParam) (*domain.Thread, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateThread", param)
	ret0, _ := ret[0].(*domain.Thread)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateThread indicates an expected call of CreateThread.
func (mr *MockThreadApplicationMockRecorder) CreateThread(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateThread", reflect.TypeOf((*MockThreadApplication)(nil).CreateThread), param)
}

// DeleteThread mocks base method.
func (m *MockThreadApplication) DeleteThread(param *params.DeleteThreadAppLayerParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteThread", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThread indicates an expected call of DeleteThread.
func (mr *MockThreadApplicationMockRecorder) DeleteThread(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThread", reflect.TypeOf((*MockThreadApplication)(nil).DeleteThread), param)
}

// EditThread mocks base method.
func (m *MockThreadApplication) EditThread(param *params.EditThreadAppLayerParam) (*domain.Thread, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditThread", param)
	ret0, _ := ret[0].(*domain.Thread)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditThread indicates an expected call of EditThread.
func (mr *MockThreadApplicationMockRecorder) EditThread(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditThread", reflect.TypeOf((*MockThreadApplication)(nil).EditThread), param)
}

// GetAllThread mocks base method.
func (m *MockThreadApplication) GetAllThread() (*[]domain.Thread, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllThread")
	ret0, _ := ret[0].(*[]domain.Thread)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllThread indicates an expected call of GetAllThread.
func (mr *MockThreadApplicationMockRecorder) GetAllThread() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllThread", reflect.TypeOf((*MockThreadApplication)(nil).GetAllThread))
}

// GetByThreadKey mocks base method.
func (m *MockThreadApplication) GetByThreadKey(threadKey string) (*domain.Thread, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByThreadKey", threadKey)
	ret0, _ := ret[0].(*domain.Thread)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByThreadKey indicates an expected call of GetByThreadKey.
func (mr *MockThreadApplicationMockRecorder) GetByThreadKey(threadKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByThreadKey", reflect.TypeOf((*MockThreadApplication)(nil).GetByThreadKey), threadKey)
}
