// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/dbstorage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	dgoogauth "github.com/dgryski/dgoogauth"
	gomock "go.uber.org/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	repository "github.com/ole-larsen/plutonium/internal/storage/db/repository"
)

// MockDBStorageInterface is a mock of DBStorageInterface interface.
type MockDBStorageInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDBStorageInterfaceMockRecorder
}

// MockDBStorageInterfaceMockRecorder is the mock recorder for MockDBStorageInterface.
type MockDBStorageInterfaceMockRecorder struct {
	mock *MockDBStorageInterface
}

// NewMockDBStorageInterface creates a new mock instance.
func NewMockDBStorageInterface(ctrl *gomock.Controller) *MockDBStorageInterface {
	mock := &MockDBStorageInterface{ctrl: ctrl}
	mock.recorder = &MockDBStorageInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBStorageInterface) EXPECT() *MockDBStorageInterfaceMockRecorder {
	return m.recorder
}

// ConnectUsersRepository mocks base method.
func (m *MockDBStorageInterface) ConnectUsersRepository(ctx context.Context, sqlxDB *sqlx.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectUsersRepository", ctx, sqlxDB)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectUsersRepository indicates an expected call of ConnectUsersRepository.
func (mr *MockDBStorageInterfaceMockRecorder) ConnectUsersRepository(ctx, sqlxDB interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectUsersRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).ConnectUsersRepository), ctx, sqlxDB)
}

// CreateUser mocks base method.
func (m *MockDBStorageInterface) CreateUser(ctx context.Context, userMap map[string]interface{}) (*dgoogauth.OTPConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, userMap)
	ret0, _ := ret[0].(*dgoogauth.OTPConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDBStorageInterfaceMockRecorder) CreateUser(ctx, userMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDBStorageInterface)(nil).CreateUser), ctx, userMap)
}

// GetUser mocks base method.
func (m *MockDBStorageInterface) GetUser(ctx context.Context, email string) (*repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email)
	ret0, _ := ret[0].(*repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockDBStorageInterfaceMockRecorder) GetUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockDBStorageInterface)(nil).GetUser), ctx, email)
}

// Init mocks base method.
func (m *MockDBStorageInterface) Init(ctx context.Context, sqlxDB *sqlx.DB) (*sqlx.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", ctx, sqlxDB)
	ret0, _ := ret[0].(*sqlx.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init.
func (mr *MockDBStorageInterfaceMockRecorder) Init(ctx, sqlxDB interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockDBStorageInterface)(nil).Init), ctx, sqlxDB)
}

// Ping mocks base method.
func (m *MockDBStorageInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockDBStorageInterfaceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDBStorageInterface)(nil).Ping))
}
