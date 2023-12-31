// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase/usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	util "fiber-upload-file/internal/util"
	multipart "mime/multipart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecaseInterface is a mock of UsecaseInterface interface.
type MockUsecaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseInterfaceMockRecorder
}

// MockUsecaseInterfaceMockRecorder is the mock recorder for MockUsecaseInterface.
type MockUsecaseInterfaceMockRecorder struct {
	mock *MockUsecaseInterface
}

// NewMockUsecaseInterface creates a new mock instance.
func NewMockUsecaseInterface(ctrl *gomock.Controller) *MockUsecaseInterface {
	mock := &MockUsecaseInterface{ctrl: ctrl}
	mock.recorder = &MockUsecaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecaseInterface) EXPECT() *MockUsecaseInterfaceMockRecorder {
	return m.recorder
}

// HitAPIsConcurrently mocks base method.
func (m *MockUsecaseInterface) HitAPIsConcurrently() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HitAPIsConcurrently")
}

// HitAPIsConcurrently indicates an expected call of HitAPIsConcurrently.
func (mr *MockUsecaseInterfaceMockRecorder) HitAPIsConcurrently() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HitAPIsConcurrently", reflect.TypeOf((*MockUsecaseInterface)(nil).HitAPIsConcurrently))
}

// SaveFile mocks base method.
func (m *MockUsecaseInterface) SaveFile(arg0 util.WebContext, arg1 *multipart.FileHeader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockUsecaseInterfaceMockRecorder) SaveFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockUsecaseInterface)(nil).SaveFile), arg0, arg1)
}
