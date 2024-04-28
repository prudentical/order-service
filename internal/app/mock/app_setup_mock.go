// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/app_setup.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/app_setup.go -destination=internal/app/mock/app_setup_mock.go
//

// Package mock_app is a generated GoMock package.
package mock_app

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAppSetupManager is a mock of AppSetupManager interface.
type MockAppSetupManager struct {
	ctrl     *gomock.Controller
	recorder *MockAppSetupManagerMockRecorder
}

// MockAppSetupManagerMockRecorder is the mock recorder for MockAppSetupManager.
type MockAppSetupManagerMockRecorder struct {
	mock *MockAppSetupManager
}

// NewMockAppSetupManager creates a new mock instance.
func NewMockAppSetupManager(ctrl *gomock.Controller) *MockAppSetupManager {
	mock := &MockAppSetupManager{ctrl: ctrl}
	mock.recorder = &MockAppSetupManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppSetupManager) EXPECT() *MockAppSetupManagerMockRecorder {
	return m.recorder
}

// Setup mocks base method.
func (m *MockAppSetupManager) Setup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup.
func (mr *MockAppSetupManagerMockRecorder) Setup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockAppSetupManager)(nil).Setup))
}

// Shutdown mocks base method.
func (m *MockAppSetupManager) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockAppSetupManagerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockAppSetupManager)(nil).Shutdown))
}