// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/db/repository/pages.go
//
// Generated by this command:
//
//	mockgen -source=internal/storage/db/repository/pages.go -destination=internal/storage/mocks/pages_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	sqlx "github.com/jmoiron/sqlx"
	gomock "go.uber.org/mock/gomock"
)

// MockPagesRepositoryInterface is a mock of PagesRepositoryInterface interface.
type MockPagesRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPagesRepositoryInterfaceMockRecorder
	isgomock struct{}
}

// MockPagesRepositoryInterfaceMockRecorder is the mock recorder for MockPagesRepositoryInterface.
type MockPagesRepositoryInterfaceMockRecorder struct {
	mock *MockPagesRepositoryInterface
}

// NewMockPagesRepositoryInterface creates a new mock instance.
func NewMockPagesRepositoryInterface(ctrl *gomock.Controller) *MockPagesRepositoryInterface {
	mock := &MockPagesRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockPagesRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPagesRepositoryInterface) EXPECT() *MockPagesRepositoryInterfaceMockRecorder {
	return m.recorder
}

// InnerDB mocks base method.
func (m *MockPagesRepositoryInterface) InnerDB() *sqlx.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InnerDB")
	ret0, _ := ret[0].(*sqlx.DB)
	return ret0
}

// InnerDB indicates an expected call of InnerDB.
func (mr *MockPagesRepositoryInterfaceMockRecorder) InnerDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InnerDB", reflect.TypeOf((*MockPagesRepositoryInterface)(nil).InnerDB))
}

// MigrateContext mocks base method.
func (m *MockPagesRepositoryInterface) MigrateContext(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateContext", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// MigrateContext indicates an expected call of MigrateContext.
func (mr *MockPagesRepositoryInterfaceMockRecorder) MigrateContext(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateContext", reflect.TypeOf((*MockPagesRepositoryInterface)(nil).MigrateContext), ctx)
}

// Ping mocks base method.
func (m *MockPagesRepositoryInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockPagesRepositoryInterfaceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockPagesRepositoryInterface)(nil).Ping))
}