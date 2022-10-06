// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wordict/w-api/internal/handler (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/wordict/w-api/internal/entity"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Signup mocks base method.
func (m *MockService) Signup(arg0 context.Context, arg1 entity.SignupRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signup indicates an expected call of Signup.
func (mr *MockServiceMockRecorder) Signup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockService)(nil).Signup), arg0, arg1)
}
