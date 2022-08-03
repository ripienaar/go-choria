// Code generated by MockGen. DO NOT EDIT.
// Source: ../request_signer.go

// Package imock is a generated GoMock package.
package imock

import (
	context "context"
	url "net/url"
	reflect "reflect"

	inter "github.com/choria-io/go-choria/inter"
	gomock "github.com/golang/mock/gomock"
)

// MockRequestSignerConfig is a mock of RequestSignerConfig interface.
type MockRequestSignerConfig struct {
	ctrl     *gomock.Controller
	recorder *MockRequestSignerConfigMockRecorder
}

// MockRequestSignerConfigMockRecorder is the mock recorder for MockRequestSignerConfig.
type MockRequestSignerConfigMockRecorder struct {
	mock *MockRequestSignerConfig
}

// NewMockRequestSignerConfig creates a new mock instance.
func NewMockRequestSignerConfig(ctrl *gomock.Controller) *MockRequestSignerConfig {
	mock := &MockRequestSignerConfig{ctrl: ctrl}
	mock.recorder = &MockRequestSignerConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestSignerConfig) EXPECT() *MockRequestSignerConfigMockRecorder {
	return m.recorder
}

// RemoteSignerToken mocks base method.
func (m *MockRequestSignerConfig) RemoteSignerToken() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteSignerToken")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteSignerToken indicates an expected call of RemoteSignerToken.
func (mr *MockRequestSignerConfigMockRecorder) RemoteSignerToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteSignerToken", reflect.TypeOf((*MockRequestSignerConfig)(nil).RemoteSignerToken))
}

// RemoteSignerURL mocks base method.
func (m *MockRequestSignerConfig) RemoteSignerURL() (*url.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteSignerURL")
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteSignerURL indicates an expected call of RemoteSignerURL.
func (mr *MockRequestSignerConfigMockRecorder) RemoteSignerURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteSignerURL", reflect.TypeOf((*MockRequestSignerConfig)(nil).RemoteSignerURL))
}

// MockRequestSigner is a mock of RequestSigner interface.
type MockRequestSigner struct {
	ctrl     *gomock.Controller
	recorder *MockRequestSignerMockRecorder
}

// MockRequestSignerMockRecorder is the mock recorder for MockRequestSigner.
type MockRequestSignerMockRecorder struct {
	mock *MockRequestSigner
}

// NewMockRequestSigner creates a new mock instance.
func NewMockRequestSigner(ctrl *gomock.Controller) *MockRequestSigner {
	mock := &MockRequestSigner{ctrl: ctrl}
	mock.recorder = &MockRequestSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestSigner) EXPECT() *MockRequestSignerMockRecorder {
	return m.recorder
}

// Kind mocks base method.
func (m *MockRequestSigner) Kind() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(string)
	return ret0
}

// Kind indicates an expected call of Kind.
func (mr *MockRequestSignerMockRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockRequestSigner)(nil).Kind))
}

// Sign mocks base method.
func (m *MockRequestSigner) Sign(ctx context.Context, request []byte, cfg inter.RequestSignerConfig) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", ctx, request, cfg)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockRequestSignerMockRecorder) Sign(ctx, request, cfg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockRequestSigner)(nil).Sign), ctx, request, cfg)
}
