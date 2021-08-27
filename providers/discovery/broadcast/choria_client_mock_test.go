// Code generated by MockGen. DO NOT EDIT.
// Source: broadcast.go

// Package broadcast is a generated GoMock package.
package broadcast

import (
	context "context"
	reflect "reflect"

	client "github.com/choria-io/go-choria/client/client"
	inter "github.com/choria-io/go-choria/inter"
	gomock "github.com/golang/mock/gomock"
)

// MockChoriaClient is a mock of ChoriaClient interface.
type MockChoriaClient struct {
	ctrl     *gomock.Controller
	recorder *MockChoriaClientMockRecorder
}

// MockChoriaClientMockRecorder is the mock recorder for MockChoriaClient.
type MockChoriaClientMockRecorder struct {
	mock *MockChoriaClient
}

// NewMockChoriaClient creates a new mock instance.
func NewMockChoriaClient(ctrl *gomock.Controller) *MockChoriaClient {
	mock := &MockChoriaClient{ctrl: ctrl}
	mock.recorder = &MockChoriaClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChoriaClient) EXPECT() *MockChoriaClientMockRecorder {
	return m.recorder
}

// Request mocks base method.
func (m *MockChoriaClient) Request(ctx context.Context, msg inter.Message, handler client.Handler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", ctx, msg, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// Request indicates an expected call of Request.
func (mr *MockChoriaClientMockRecorder) Request(ctx, msg, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockChoriaClient)(nil).Request), ctx, msg, handler)
}
