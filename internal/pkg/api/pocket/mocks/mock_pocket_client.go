// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/olezhek28/items-keeper/internal/pkg/api/pocket (interfaces: IPocketClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pocket "github.com/zhashkevych/go-pocket-sdk"
)

// MockIPocketClient is a mock of IPocketClient interface.
type MockIPocketClient struct {
	ctrl     *gomock.Controller
	recorder *MockIPocketClientMockRecorder
}

// MockIPocketClientMockRecorder is the mock recorder for MockIPocketClient.
type MockIPocketClientMockRecorder struct {
	mock *MockIPocketClient
}

// NewMockIPocketClient creates a new mock instance.
func NewMockIPocketClient(ctrl *gomock.Controller) *MockIPocketClient {
	mock := &MockIPocketClient{ctrl: ctrl}
	mock.recorder = &MockIPocketClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPocketClient) EXPECT() *MockIPocketClientMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIPocketClient) Add(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockIPocketClientMockRecorder) Add(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIPocketClient)(nil).Add), arg0, arg1, arg2)
}

// Authorize mocks base method.
func (m *MockIPocketClient) Authorize(arg0 context.Context, arg1 string) (*pocket.AuthorizeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", arg0, arg1)
	ret0, _ := ret[0].(*pocket.AuthorizeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorize indicates an expected call of Authorize.
func (mr *MockIPocketClientMockRecorder) Authorize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockIPocketClient)(nil).Authorize), arg0, arg1)
}

// GetAuthorizationLink mocks base method.
func (m *MockIPocketClient) GetAuthorizationLink(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorizationLink", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizationLink indicates an expected call of GetAuthorizationLink.
func (mr *MockIPocketClientMockRecorder) GetAuthorizationLink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizationLink", reflect.TypeOf((*MockIPocketClient)(nil).GetAuthorizationLink), arg0)
}

// GetRequestToken mocks base method.
func (m *MockIPocketClient) GetRequestToken(arg0 context.Context, arg1 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestToken indicates an expected call of GetRequestToken.
func (mr *MockIPocketClientMockRecorder) GetRequestToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestToken", reflect.TypeOf((*MockIPocketClient)(nil).GetRequestToken), arg0, arg1)
}
