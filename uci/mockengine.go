// Code generated by MockGen. DO NOT EDIT.
// Source: ../engine/engine.go
//
// Generated by this command:
//
//	mockgen -source=../engine/engine.go -destination=mockengine.go -package uci
//

// Package uci is a generated GoMock package.
package uci

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockEngine) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init.
func (mr *MockEngineMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockEngine)(nil).Init))
}

// IsRunning mocks base method.
func (m *MockEngine) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning.
func (mr *MockEngineMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockEngine)(nil).IsRunning))
}

// Play mocks base method.
func (m *MockEngine) Play(moveNotation string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Play", moveNotation)
}

// Play indicates an expected call of Play.
func (mr *MockEngineMockRecorder) Play(moveNotation any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Play", reflect.TypeOf((*MockEngine)(nil).Play), moveNotation)
}

// Quit mocks base method.
func (m *MockEngine) Quit() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Quit")
}

// Quit indicates an expected call of Quit.
func (mr *MockEngineMockRecorder) Quit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockEngine)(nil).Quit))
}

// Reset mocks base method.
func (m *MockEngine) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockEngineMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockEngine)(nil).Reset))
}

// SetupPosition mocks base method.
func (m *MockEngine) SetupPosition() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetupPosition")
}

// SetupPosition indicates an expected call of SetupPosition.
func (mr *MockEngineMockRecorder) SetupPosition() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupPosition", reflect.TypeOf((*MockEngine)(nil).SetupPosition))
}

// Stop mocks base method.
func (m *MockEngine) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockEngineMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockEngine)(nil).Stop))
}
