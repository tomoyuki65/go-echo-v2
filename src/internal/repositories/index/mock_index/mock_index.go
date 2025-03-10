// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repositories/index/index.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repositories/index/index.go -destination=./internal/repositories/index/mock_index/mock_index.go
//

// Package mock_index is a generated GoMock package.
package mock_index

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIndexRepository is a mock of IndexRepository interface.
type MockIndexRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIndexRepositoryMockRecorder
	isgomock struct{}
}

// MockIndexRepositoryMockRecorder is the mock recorder for MockIndexRepository.
type MockIndexRepositoryMockRecorder struct {
	mock *MockIndexRepository
}

// NewMockIndexRepository creates a new mock instance.
func NewMockIndexRepository(ctrl *gomock.Controller) *MockIndexRepository {
	mock := &MockIndexRepository{ctrl: ctrl}
	mock.recorder = &MockIndexRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexRepository) EXPECT() *MockIndexRepositoryMockRecorder {
	return m.recorder
}

// Hello mocks base method.
func (m *MockIndexRepository) Hello() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hello")
	ret0, _ := ret[0].(string)
	return ret0
}

// Hello indicates an expected call of Hello.
func (mr *MockIndexRepositoryMockRecorder) Hello() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockIndexRepository)(nil).Hello))
}
