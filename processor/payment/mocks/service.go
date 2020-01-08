// Code generated by MockGen. DO NOT EDIT.
// Source: processor/payment/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	models "processor/payment/models"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// ProcessPayment mocks base method
func (m *MockService) ProcessPayment(req models.Request, errors *models.Error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPayment", req, errors)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ProcessPayment indicates an expected call of ProcessPayment
func (mr *MockServiceMockRecorder) ProcessPayment(req, errors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPayment", reflect.TypeOf((*MockService)(nil).ProcessPayment), req, errors)
}