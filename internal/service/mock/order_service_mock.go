// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/order_service.go
//
// Generated by this command:
//
//	mockgen -source=internal/service/order_service.go -destination=internal/service/mock/order_service_mock.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	dto "order-service/internal/dto"
	model "order-service/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockOrderService is a mock of OrderService interface.
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService.
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance.
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderService) Create(order dto.OrderDTO) (model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", order)
	ret0, _ := ret[0].(model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrderServiceMockRecorder) Create(order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderService)(nil).Create), order)
}
