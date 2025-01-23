// Code generated by MockGen. DO NOT EDIT.
// Source: google.golang.org/grpc/balancer (interfaces: ClientConn,SubConn)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	balancer "google.golang.org/grpc/balancer"
	resolver "google.golang.org/grpc/resolver"
)

// MockClientConn is a mock of ClientConn interface.
type MockClientConn struct {
	// The ClientConn interface below is embedded by a manual edit to comply 
	// with grpc's requirement to embed an existing ClientConn to allow grpc to
	// add new methods to the interface easily.
	balancer.ClientConn
	ctrl     *gomock.Controller
	recorder *MockClientConnMockRecorder
}

// MockClientConnMockRecorder is the mock recorder for MockClientConn.
type MockClientConnMockRecorder struct {
	mock *MockClientConn
}

// NewMockClientConn creates a new mock instance.
func NewMockClientConn(ctrl *gomock.Controller) *MockClientConn {
	mock := &MockClientConn{ctrl: ctrl}
	mock.recorder = &MockClientConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientConn) EXPECT() *MockClientConnMockRecorder {
	return m.recorder
}

// NewSubConn mocks base method.
func (m *MockClientConn) NewSubConn(arg0 []resolver.Address, arg1 balancer.NewSubConnOptions) (balancer.SubConn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSubConn", arg0, arg1)
	ret0, _ := ret[0].(balancer.SubConn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSubConn indicates an expected call of NewSubConn.
func (mr *MockClientConnMockRecorder) NewSubConn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSubConn", reflect.TypeOf((*MockClientConn)(nil).NewSubConn), arg0, arg1)
}

// RemoveSubConn mocks base method.
func (m *MockClientConn) RemoveSubConn(arg0 balancer.SubConn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveSubConn", arg0)
}

// RemoveSubConn indicates an expected call of RemoveSubConn.
func (mr *MockClientConnMockRecorder) RemoveSubConn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSubConn", reflect.TypeOf((*MockClientConn)(nil).RemoveSubConn), arg0)
}

// ResolveNow mocks base method.
func (m *MockClientConn) ResolveNow(arg0 resolver.ResolveNowOptions) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResolveNow", arg0)
}

// ResolveNow indicates an expected call of ResolveNow.
func (mr *MockClientConnMockRecorder) ResolveNow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveNow", reflect.TypeOf((*MockClientConn)(nil).ResolveNow), arg0)
}

// Target mocks base method.
func (m *MockClientConn) Target() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Target")
	ret0, _ := ret[0].(string)
	return ret0
}

// Target indicates an expected call of Target.
func (mr *MockClientConnMockRecorder) Target() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Target", reflect.TypeOf((*MockClientConn)(nil).Target))
}

// UpdateAddresses mocks base method.
func (m *MockClientConn) UpdateAddresses(arg0 balancer.SubConn, arg1 []resolver.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateAddresses", arg0, arg1)
}

// UpdateAddresses indicates an expected call of UpdateAddresses.
func (mr *MockClientConnMockRecorder) UpdateAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddresses", reflect.TypeOf((*MockClientConn)(nil).UpdateAddresses), arg0, arg1)
}

// UpdateState mocks base method.
func (m *MockClientConn) UpdateState(arg0 balancer.State) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateState", arg0)
}

// UpdateState indicates an expected call of UpdateState.
func (mr *MockClientConnMockRecorder) UpdateState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateState", reflect.TypeOf((*MockClientConn)(nil).UpdateState), arg0)
}

// MockSubConn is a mock of SubConn interface.
type MockSubConn struct {
	// The SubConn interface is added by a manual edit to 
	// comply with the requirements in the SubConn interface
	// doc comment.
	balancer.SubConn
	ctrl     *gomock.Controller
	recorder *MockSubConnMockRecorder
}

// MockSubConnMockRecorder is the mock recorder for MockSubConn.
type MockSubConnMockRecorder struct {
	mock *MockSubConn
}

// NewMockSubConn creates a new mock instance.
func NewMockSubConn(ctrl *gomock.Controller) *MockSubConn {
	mock := &MockSubConn{ctrl: ctrl}
	mock.recorder = &MockSubConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubConn) EXPECT() *MockSubConnMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockSubConn) Connect() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Connect")
}

// Connect indicates an expected call of Connect.
func (mr *MockSubConnMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockSubConn)(nil).Connect))
}

// GetOrBuildProducer mocks base method.
func (m *MockSubConn) GetOrBuildProducer(arg0 balancer.ProducerBuilder) (balancer.Producer, func()) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrBuildProducer", arg0)
	ret0, _ := ret[0].(balancer.Producer)
	ret1, _ := ret[1].(func())
	return ret0, ret1
}

// GetOrBuildProducer indicates an expected call of GetOrBuildProducer.
func (mr *MockSubConnMockRecorder) GetOrBuildProducer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrBuildProducer", reflect.TypeOf((*MockSubConn)(nil).GetOrBuildProducer), arg0)
}

// Shutdown mocks base method.
func (m *MockSubConn) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockSubConnMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockSubConn)(nil).Shutdown))
}

// UpdateAddresses mocks base method.
func (m *MockSubConn) UpdateAddresses(arg0 []resolver.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateAddresses", arg0)
}

// UpdateAddresses indicates an expected call of UpdateAddresses.
func (mr *MockSubConnMockRecorder) UpdateAddresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddresses", reflect.TypeOf((*MockSubConn)(nil).UpdateAddresses), arg0)
}
