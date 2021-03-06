// Code generated by MockGen. DO NOT EDIT.
// Source: inventory.go

// Package inventory is a generated GoMock package.
package inventory

import (
	config "github.com/choria-io/go-choria/config"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
)

// MockChoriaFramework is a mock of ChoriaFramework interface
type MockChoriaFramework struct {
	ctrl     *gomock.Controller
	recorder *MockChoriaFrameworkMockRecorder
}

// MockChoriaFrameworkMockRecorder is the mock recorder for MockChoriaFramework
type MockChoriaFrameworkMockRecorder struct {
	mock *MockChoriaFramework
}

// NewMockChoriaFramework creates a new mock instance
func NewMockChoriaFramework(ctrl *gomock.Controller) *MockChoriaFramework {
	mock := &MockChoriaFramework{ctrl: ctrl}
	mock.recorder = &MockChoriaFrameworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChoriaFramework) EXPECT() *MockChoriaFrameworkMockRecorder {
	return m.recorder
}

// Logger mocks base method
func (m *MockChoriaFramework) Logger(arg0 string) *logrus.Entry {
	ret := m.ctrl.Call(m, "Logger", arg0)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

// Logger indicates an expected call of Logger
func (mr *MockChoriaFrameworkMockRecorder) Logger(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockChoriaFramework)(nil).Logger), arg0)
}

// Configuration mocks base method
func (m *MockChoriaFramework) Configuration() *config.Config {
	ret := m.ctrl.Call(m, "Configuration")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// Configuration indicates an expected call of Configuration
func (mr *MockChoriaFrameworkMockRecorder) Configuration() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configuration", reflect.TypeOf((*MockChoriaFramework)(nil).Configuration))
}
