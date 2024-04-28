// Code generated by MockGen. DO NOT EDIT.
// Source: internal/api/health_check.go
//
// Generated by this command:
//
//	mockgen -source=internal/api/health_check.go -destination=internal/api/mock/health_check_mock.go
//

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockHealthCheck is a mock of HealthCheck interface.
type MockHealthCheck struct {
	ctrl     *gomock.Controller
	recorder *MockHealthCheckMockRecorder
}

// MockHealthCheckMockRecorder is the mock recorder for MockHealthCheck.
type MockHealthCheckMockRecorder struct {
	mock *MockHealthCheck
}

// NewMockHealthCheck creates a new mock instance.
func NewMockHealthCheck(ctrl *gomock.Controller) *MockHealthCheck {
	mock := &MockHealthCheck{ctrl: ctrl}
	mock.recorder = &MockHealthCheckMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthCheck) EXPECT() *MockHealthCheckMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockHealthCheck) Check(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockHealthCheckMockRecorder) Check(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockHealthCheck)(nil).Check), c)
}

// HandleRoutes mocks base method.
func (m *MockHealthCheck) HandleRoutes(e *echo.Echo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleRoutes", e)
}

// HandleRoutes indicates an expected call of HandleRoutes.
func (mr *MockHealthCheckMockRecorder) HandleRoutes(e any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRoutes", reflect.TypeOf((*MockHealthCheck)(nil).HandleRoutes), e)
}