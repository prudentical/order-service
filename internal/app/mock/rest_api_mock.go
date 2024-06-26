// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/rest_api.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/rest_api.go -destination=internal/app/mock/rest_api_mock.go
//

// Package mock_app is a generated GoMock package.
package mock_app

import (
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockRESTApp is a mock of RESTApp interface.
type MockRESTApp struct {
	ctrl     *gomock.Controller
	recorder *MockRESTAppMockRecorder
}

// MockRESTAppMockRecorder is the mock recorder for MockRESTApp.
type MockRESTAppMockRecorder struct {
	mock *MockRESTApp
}

// NewMockRESTApp creates a new mock instance.
func NewMockRESTApp(ctrl *gomock.Controller) *MockRESTApp {
	mock := &MockRESTApp{ctrl: ctrl}
	mock.recorder = &MockRESTAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRESTApp) EXPECT() *MockRESTAppMockRecorder {
	return m.recorder
}

// server mocks base method.
func (m *MockRESTApp) server() *echo.Echo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "server")
	ret0, _ := ret[0].(*echo.Echo)
	return ret0
}

// server indicates an expected call of server.
func (mr *MockRESTAppMockRecorder) server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "server", reflect.TypeOf((*MockRESTApp)(nil).server))
}

// setup mocks base method.
func (m *MockRESTApp) setup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setup")
	ret0, _ := ret[0].(error)
	return ret0
}

// setup indicates an expected call of setup.
func (mr *MockRESTAppMockRecorder) setup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setup", reflect.TypeOf((*MockRESTApp)(nil).setup))
}
