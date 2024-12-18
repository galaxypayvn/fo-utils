// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/broker/mqtt/producer.go

// Package mqtt is a generated GoMock package.
package mqtt

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockISender is a mock of ISender interface.
type MockISender struct {
	ctrl     *gomock.Controller
	recorder *MockISenderMockRecorder
}

// MockISenderMockRecorder is the mock recorder for MockISender.
type MockISenderMockRecorder struct {
	mock *MockISender
}

// NewMockISender creates a new mock instance.
func NewMockISender(ctrl *gomock.Controller) *MockISender {
	mock := &MockISender{ctrl: ctrl}
	mock.recorder = &MockISenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISender) EXPECT() *MockISenderMockRecorder {
	return m.recorder
}

// SendBytes mocks base method.
func (m *MockISender) SendBytes(msg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendBytes", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendBytes indicates an expected call of SendBytes.
func (mr *MockISenderMockRecorder) SendBytes(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBytes", reflect.TypeOf((*MockISender)(nil).SendBytes), msg)
}

// SendString mocks base method.
func (m *MockISender) SendString(msg string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendString", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendString indicates an expected call of SendString.
func (mr *MockISenderMockRecorder) SendString(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendString", reflect.TypeOf((*MockISender)(nil).SendString), msg)
}

// MockIProducer is a mock of IProducer interface.
type MockIProducer struct {
	ctrl     *gomock.Controller
	recorder *MockIProducerMockRecorder
}

// MockIProducerMockRecorder is the mock recorder for MockIProducer.
type MockIProducerMockRecorder struct {
	mock *MockIProducer
}

// NewMockIProducer creates a new mock instance.
func NewMockIProducer(ctrl *gomock.Controller) *MockIProducer {
	mock := &MockIProducer{ctrl: ctrl}
	mock.recorder = &MockIProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProducer) EXPECT() *MockIProducerMockRecorder {
	return m.recorder
}

// GetSender mocks base method.
func (m *MockIProducer) GetSender(topic string, qos byte) ISender {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSender", topic, qos)
	ret0, _ := ret[0].(ISender)
	return ret0
}

// GetSender indicates an expected call of GetSender.
func (mr *MockIProducerMockRecorder) GetSender(topic, qos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSender", reflect.TypeOf((*MockIProducer)(nil).GetSender), topic, qos)
}
