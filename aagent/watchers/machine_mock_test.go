// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/choria-io/go-choria/aagent/watchers (interfaces: Machine)
//
// Generated by this command:
//
//	mockgen -write_generate_directive -package watchers -destination machine_mock_test.go github.com/choria-io/go-choria/aagent/watchers Machine
//

// Package watchers is a generated GoMock package.
package watchers

import (
	json "encoding/json"
	reflect "reflect"

	lifecycle "github.com/choria-io/go-choria/lifecycle"
	jsm "github.com/nats-io/jsm.go"
	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -write_generate_directive -package watchers -destination machine_mock_test.go github.com/choria-io/go-choria/aagent/watchers Machine

// MockMachine is a mock of Machine interface.
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
	isgomock struct{}
}

// MockMachineMockRecorder is the mock recorder for MockMachine.
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance.
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// ChoriaStatusFile mocks base method.
func (m *MockMachine) ChoriaStatusFile() (string, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChoriaStatusFile")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// ChoriaStatusFile indicates an expected call of ChoriaStatusFile.
func (mr *MockMachineMockRecorder) ChoriaStatusFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChoriaStatusFile", reflect.TypeOf((*MockMachine)(nil).ChoriaStatusFile))
}

// Data mocks base method.
func (m *MockMachine) Data() map[string]any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].(map[string]any)
	return ret0
}

// Data indicates an expected call of Data.
func (mr *MockMachineMockRecorder) Data() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockMachine)(nil).Data))
}

// DataDelete mocks base method.
func (m *MockMachine) DataDelete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataDelete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DataDelete indicates an expected call of DataDelete.
func (mr *MockMachineMockRecorder) DataDelete(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataDelete", reflect.TypeOf((*MockMachine)(nil).DataDelete), key)
}

// DataGet mocks base method.
func (m *MockMachine) DataGet(key string) (any, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataGet", key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// DataGet indicates an expected call of DataGet.
func (mr *MockMachineMockRecorder) DataGet(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataGet", reflect.TypeOf((*MockMachine)(nil).DataGet), key)
}

// DataPut mocks base method.
func (m *MockMachine) DataPut(key string, val any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataPut", key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// DataPut indicates an expected call of DataPut.
func (mr *MockMachineMockRecorder) DataPut(key, val any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataPut", reflect.TypeOf((*MockMachine)(nil).DataPut), key, val)
}

// Debugf mocks base method.
func (m *MockMachine) Debugf(name, format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{name, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockMachineMockRecorder) Debugf(name, format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockMachine)(nil).Debugf), varargs...)
}

// Directory mocks base method.
func (m *MockMachine) Directory() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Directory")
	ret0, _ := ret[0].(string)
	return ret0
}

// Directory indicates an expected call of Directory.
func (mr *MockMachineMockRecorder) Directory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Directory", reflect.TypeOf((*MockMachine)(nil).Directory))
}

// Errorf mocks base method.
func (m *MockMachine) Errorf(name, format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{name, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockMachineMockRecorder) Errorf(name, format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockMachine)(nil).Errorf), varargs...)
}

// Facts mocks base method.
func (m *MockMachine) Facts() json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Facts")
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// Facts indicates an expected call of Facts.
func (mr *MockMachineMockRecorder) Facts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Facts", reflect.TypeOf((*MockMachine)(nil).Facts))
}

// Identity mocks base method.
func (m *MockMachine) Identity() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identity")
	ret0, _ := ret[0].(string)
	return ret0
}

// Identity indicates an expected call of Identity.
func (mr *MockMachineMockRecorder) Identity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identity", reflect.TypeOf((*MockMachine)(nil).Identity))
}

// Infof mocks base method.
func (m *MockMachine) Infof(name, format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{name, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockMachineMockRecorder) Infof(name, format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockMachine)(nil).Infof), varargs...)
}

// InstanceID mocks base method.
func (m *MockMachine) InstanceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// InstanceID indicates an expected call of InstanceID.
func (mr *MockMachineMockRecorder) InstanceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceID", reflect.TypeOf((*MockMachine)(nil).InstanceID))
}

// JetStreamConnection mocks base method.
func (m *MockMachine) JetStreamConnection() (*jsm.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JetStreamConnection")
	ret0, _ := ret[0].(*jsm.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JetStreamConnection indicates an expected call of JetStreamConnection.
func (mr *MockMachineMockRecorder) JetStreamConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JetStreamConnection", reflect.TypeOf((*MockMachine)(nil).JetStreamConnection))
}

// MainCollective mocks base method.
func (m *MockMachine) MainCollective() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MainCollective")
	ret0, _ := ret[0].(string)
	return ret0
}

// MainCollective indicates an expected call of MainCollective.
func (mr *MockMachineMockRecorder) MainCollective() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MainCollective", reflect.TypeOf((*MockMachine)(nil).MainCollective))
}

// Name mocks base method.
func (m *MockMachine) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockMachineMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockMachine)(nil).Name))
}

// NotifyWatcherState mocks base method.
func (m *MockMachine) NotifyWatcherState(arg0 string, arg1 any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyWatcherState", arg0, arg1)
}

// NotifyWatcherState indicates an expected call of NotifyWatcherState.
func (mr *MockMachineMockRecorder) NotifyWatcherState(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyWatcherState", reflect.TypeOf((*MockMachine)(nil).NotifyWatcherState), arg0, arg1)
}

// OverrideData mocks base method.
func (m *MockMachine) OverrideData() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OverrideData")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OverrideData indicates an expected call of OverrideData.
func (mr *MockMachineMockRecorder) OverrideData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OverrideData", reflect.TypeOf((*MockMachine)(nil).OverrideData))
}

// PublishLifecycleEvent mocks base method.
func (m *MockMachine) PublishLifecycleEvent(t lifecycle.Type, opts ...lifecycle.Option) {
	m.ctrl.T.Helper()
	varargs := []any{t}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "PublishLifecycleEvent", varargs...)
}

// PublishLifecycleEvent indicates an expected call of PublishLifecycleEvent.
func (mr *MockMachineMockRecorder) PublishLifecycleEvent(t any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{t}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishLifecycleEvent", reflect.TypeOf((*MockMachine)(nil).PublishLifecycleEvent), varargs...)
}

// SignerKey mocks base method.
func (m *MockMachine) SignerKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignerKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// SignerKey indicates an expected call of SignerKey.
func (mr *MockMachineMockRecorder) SignerKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignerKey", reflect.TypeOf((*MockMachine)(nil).SignerKey))
}

// State mocks base method.
func (m *MockMachine) State() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(string)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockMachineMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockMachine)(nil).State))
}

// TextFileDirectory mocks base method.
func (m *MockMachine) TextFileDirectory() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TextFileDirectory")
	ret0, _ := ret[0].(string)
	return ret0
}

// TextFileDirectory indicates an expected call of TextFileDirectory.
func (mr *MockMachineMockRecorder) TextFileDirectory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TextFileDirectory", reflect.TypeOf((*MockMachine)(nil).TextFileDirectory))
}

// TimeStampSeconds mocks base method.
func (m *MockMachine) TimeStampSeconds() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeStampSeconds")
	ret0, _ := ret[0].(int64)
	return ret0
}

// TimeStampSeconds indicates an expected call of TimeStampSeconds.
func (mr *MockMachineMockRecorder) TimeStampSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeStampSeconds", reflect.TypeOf((*MockMachine)(nil).TimeStampSeconds))
}

// Transition mocks base method.
func (m *MockMachine) Transition(t string, args ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{t}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transition", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transition indicates an expected call of Transition.
func (mr *MockMachineMockRecorder) Transition(t any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{t}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transition", reflect.TypeOf((*MockMachine)(nil).Transition), varargs...)
}

// Version mocks base method.
func (m *MockMachine) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockMachineMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockMachine)(nil).Version))
}

// Warnf mocks base method.
func (m *MockMachine) Warnf(name, format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{name, format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockMachineMockRecorder) Warnf(name, format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockMachine)(nil).Warnf), varargs...)
}

// Watchers mocks base method.
func (m *MockMachine) Watchers() []*WatcherDef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watchers")
	ret0, _ := ret[0].([]*WatcherDef)
	return ret0
}

// Watchers indicates an expected call of Watchers.
func (mr *MockMachineMockRecorder) Watchers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watchers", reflect.TypeOf((*MockMachine)(nil).Watchers))
}
