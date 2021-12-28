// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/olezhek28/items-keeper/internal/pkg/api/telegram (interfaces: ITelegramClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITelegramClient is a mock of ITelegramClient interface.
type MockITelegramClient struct {
	ctrl     *gomock.Controller
	recorder *MockITelegramClientMockRecorder
}

// MockITelegramClientMockRecorder is the mock recorder for MockITelegramClient.
type MockITelegramClientMockRecorder struct {
	mock *MockITelegramClient
}

// NewMockITelegramClient creates a new mock instance.
func NewMockITelegramClient(ctrl *gomock.Controller) *MockITelegramClient {
	mock := &MockITelegramClient{ctrl: ctrl}
	mock.recorder = &MockITelegramClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITelegramClient) EXPECT() *MockITelegramClientMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockITelegramClient) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockITelegramClientMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockITelegramClient)(nil).Start))
}
