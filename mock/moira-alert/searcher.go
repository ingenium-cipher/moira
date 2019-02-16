// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moira-alert/moira (interfaces: Searcher)

// Package mock_moira_alert is a generated GoMock package.
package mock_moira_alert

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSearcher is a mock of Searcher interface
type MockSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockSearcherMockRecorder
}

// MockSearcherMockRecorder is the mock recorder for MockSearcher
type MockSearcherMockRecorder struct {
	mock *MockSearcher
}

// NewMockSearcher creates a new mock instance
func NewMockSearcher(ctrl *gomock.Controller) *MockSearcher {
	mock := &MockSearcher{ctrl: ctrl}
	mock.recorder = &MockSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearcher) EXPECT() *MockSearcherMockRecorder {
	return m.recorder
}

// IsReady mocks base method
func (m *MockSearcher) IsReady() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsReady")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsReady indicates an expected call of IsReady
func (mr *MockSearcherMockRecorder) IsReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReady", reflect.TypeOf((*MockSearcher)(nil).IsReady))
}

// SearchTriggers mocks base method
func (m *MockSearcher) SearchTriggers(arg0 []string, arg1 string, arg2 bool, arg3, arg4 int64) ([]string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTriggers", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchTriggers indicates an expected call of SearchTriggers
func (mr *MockSearcherMockRecorder) SearchTriggers(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTriggers", reflect.TypeOf((*MockSearcher)(nil).SearchTriggers), arg0, arg1, arg2, arg3, arg4)
}

// Start mocks base method
func (m *MockSearcher) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockSearcherMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSearcher)(nil).Start))
}

// Stop mocks base method
func (m *MockSearcher) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockSearcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockSearcher)(nil).Stop))
}