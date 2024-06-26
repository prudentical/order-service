// Code generated by MockGen. DO NOT EDIT.
// Source: internal/persistence/position_dao.go
//
// Generated by this command:
//
//	mockgen -source=internal/persistence/position_dao.go -destination=internal/persistence/mock/position_dao_mock.go
//

// Package mock_persistence is a generated GoMock package.
package mock_persistence

import (
	model "order-service/internal/model"
	persistence "order-service/internal/persistence"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPositionDAO is a mock of PositionDAO interface.
type MockPositionDAO struct {
	ctrl     *gomock.Controller
	recorder *MockPositionDAOMockRecorder
}

// MockPositionDAOMockRecorder is the mock recorder for MockPositionDAO.
type MockPositionDAOMockRecorder struct {
	mock *MockPositionDAO
}

// NewMockPositionDAO creates a new mock instance.
func NewMockPositionDAO(ctrl *gomock.Controller) *MockPositionDAO {
	mock := &MockPositionDAO{ctrl: ctrl}
	mock.recorder = &MockPositionDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPositionDAO) EXPECT() *MockPositionDAOMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPositionDAO) Create(position model.Position) (model.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", position)
	ret0, _ := ret[0].(model.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPositionDAOMockRecorder) Create(position any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPositionDAO)(nil).Create), position)
}

// DeleteByBotId mocks base method.
func (m *MockPositionDAO) DeleteByBotId(botId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByBotId", botId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByBotId indicates an expected call of DeleteByBotId.
func (mr *MockPositionDAOMockRecorder) DeleteByBotId(botId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByBotId", reflect.TypeOf((*MockPositionDAO)(nil).DeleteByBotId), botId)
}

// GetByBotIdPaginated mocks base method.
func (m *MockPositionDAO) GetByBotIdPaginated(botId int64, page, size int, status *string) (persistence.Page[model.Position], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBotIdPaginated", botId, page, size, status)
	ret0, _ := ret[0].(persistence.Page[model.Position])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBotIdPaginated indicates an expected call of GetByBotIdPaginated.
func (mr *MockPositionDAOMockRecorder) GetByBotIdPaginated(botId, page, size, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBotIdPaginated", reflect.TypeOf((*MockPositionDAO)(nil).GetByBotIdPaginated), botId, page, size, status)
}

// GetById mocks base method.
func (m *MockPositionDAO) GetById(id int64) (model.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(model.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockPositionDAOMockRecorder) GetById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPositionDAO)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockPositionDAO) Update(position model.Position) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", position)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPositionDAOMockRecorder) Update(position any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPositionDAO)(nil).Update), position)
}
