// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/position_service.go
//
// Generated by this command:
//
//	mockgen -source=internal/service/position_service.go -destination=internal/service/mock/position_service_mock.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "order-service/internal/model"
	persistence "order-service/internal/persistence"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPositionService is a mock of PositionService interface.
type MockPositionService struct {
	ctrl     *gomock.Controller
	recorder *MockPositionServiceMockRecorder
}

// MockPositionServiceMockRecorder is the mock recorder for MockPositionService.
type MockPositionServiceMockRecorder struct {
	mock *MockPositionService
}

// NewMockPositionService creates a new mock instance.
func NewMockPositionService(ctrl *gomock.Controller) *MockPositionService {
	mock := &MockPositionService{ctrl: ctrl}
	mock.recorder = &MockPositionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPositionService) EXPECT() *MockPositionServiceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPositionService) Close(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPositionServiceMockRecorder) Close(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPositionService)(nil).Close), id)
}

// Create mocks base method.
func (m *MockPositionService) Create(position model.Position) (model.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", position)
	ret0, _ := ret[0].(model.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPositionServiceMockRecorder) Create(position any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPositionService)(nil).Create), position)
}

// DeleteByBotId mocks base method.
func (m *MockPositionService) DeleteByBotId(botId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByBotId", botId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByBotId indicates an expected call of DeleteByBotId.
func (mr *MockPositionServiceMockRecorder) DeleteByBotId(botId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByBotId", reflect.TypeOf((*MockPositionService)(nil).DeleteByBotId), botId)
}

// GetByBotId mocks base method.
func (m *MockPositionService) GetByBotId(botId int64, page, size int, status string) (persistence.Page[model.Position], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBotId", botId, page, size, status)
	ret0, _ := ret[0].(persistence.Page[model.Position])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBotId indicates an expected call of GetByBotId.
func (mr *MockPositionServiceMockRecorder) GetByBotId(botId, page, size, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBotId", reflect.TypeOf((*MockPositionService)(nil).GetByBotId), botId, page, size, status)
}

// GetById mocks base method.
func (m *MockPositionService) GetById(id int64) (model.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(model.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockPositionServiceMockRecorder) GetById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPositionService)(nil).GetById), id)
}