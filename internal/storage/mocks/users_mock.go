// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/db/repository/users.go
//
// Generated by this command:
//
//	mockgen -source=internal/storage/db/repository/users.go -destination=internal/storage/mocks/users_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	sqlx "github.com/jmoiron/sqlx"
	repository "github.com/ole-larsen/plutonium/internal/storage/db/repository"
	models "github.com/ole-larsen/plutonium/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUsersRepositoryInterface is a mock of UsersRepositoryInterface interface.
type MockUsersRepositoryInterface struct {
	isgomock struct{}
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryInterfaceMockRecorder
}

// MockUsersRepositoryInterfaceMockRecorder is the mock recorder for MockUsersRepositoryInterface.
type MockUsersRepositoryInterfaceMockRecorder struct {
	mock *MockUsersRepositoryInterface
}

// NewMockUsersRepositoryInterface creates a new mock instance.
func NewMockUsersRepositoryInterface(ctrl *gomock.Controller) *MockUsersRepositoryInterface {
	mock := &MockUsersRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepositoryInterface) EXPECT() *MockUsersRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsersRepositoryInterface) Create(ctx context.Context, userMap map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUsersRepositoryInterfaceMockRecorder) Create(ctx, userMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).Create), ctx, userMap)
}

// GetPublicUserByID mocks base method.
func (m *MockUsersRepositoryInterface) GetPublicUserByID(ctx context.Context, id int64) (*models.PublicUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicUserByID", ctx, id)
	ret0, _ := ret[0].(*models.PublicUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicUserByID indicates an expected call of GetPublicUserByID.
func (mr *MockUsersRepositoryInterfaceMockRecorder) GetPublicUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicUserByID", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).GetPublicUserByID), ctx, id)
}

// GetUserByAddress mocks base method.
func (m *MockUsersRepositoryInterface) GetUserByAddress(ctx context.Context, address string) (*repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAddress", ctx, address)
	ret0, _ := ret[0].(*repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAddress indicates an expected call of GetUserByAddress.
func (mr *MockUsersRepositoryInterfaceMockRecorder) GetUserByAddress(ctx, address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAddress", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).GetUserByAddress), ctx, address)
}

// GetUserByEmail mocks base method.
func (m *MockUsersRepositoryInterface) GetUserByEmail(ctx context.Context, email string) (*repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUsersRepositoryInterfaceMockRecorder) GetUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).GetUserByEmail), ctx, email)
}

// GetUserByID mocks base method.
func (m *MockUsersRepositoryInterface) GetUserByID(ctx context.Context, id int64) (*repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUsersRepositoryInterfaceMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).GetUserByID), ctx, id)
}

// InnerDB mocks base method.
func (m *MockUsersRepositoryInterface) InnerDB() *sqlx.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InnerDB")
	ret0, _ := ret[0].(*sqlx.DB)
	return ret0
}

// InnerDB indicates an expected call of InnerDB.
func (mr *MockUsersRepositoryInterfaceMockRecorder) InnerDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InnerDB", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).InnerDB))
}

// Ping mocks base method.
func (m *MockUsersRepositoryInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockUsersRepositoryInterfaceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).Ping))
}

// UpdateGravatar mocks base method.
func (m *MockUsersRepositoryInterface) UpdateGravatar(ctx context.Context, userMap map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGravatar", ctx, userMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGravatar indicates an expected call of UpdateGravatar.
func (mr *MockUsersRepositoryInterfaceMockRecorder) UpdateGravatar(ctx, userMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGravatar", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).UpdateGravatar), ctx, userMap)
}

// UpdateNonce mocks base method.
func (m *MockUsersRepositoryInterface) UpdateNonce(ctx context.Context, userMap map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNonce", ctx, userMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNonce indicates an expected call of UpdateNonce.
func (mr *MockUsersRepositoryInterfaceMockRecorder) UpdateNonce(ctx, userMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNonce", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).UpdateNonce), ctx, userMap)
}

// UpdateSecret mocks base method.
func (m *MockUsersRepositoryInterface) UpdateSecret(ctx context.Context, userMap map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", ctx, userMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockUsersRepositoryInterfaceMockRecorder) UpdateSecret(ctx, userMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).UpdateSecret), ctx, userMap)
}

// UpdateWallpaper mocks base method.
func (m *MockUsersRepositoryInterface) UpdateWallpaper(ctx context.Context, userMap map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWallpaper", ctx, userMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWallpaper indicates an expected call of UpdateWallpaper.
func (mr *MockUsersRepositoryInterfaceMockRecorder) UpdateWallpaper(ctx, userMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWallpaper", reflect.TypeOf((*MockUsersRepositoryInterface)(nil).UpdateWallpaper), ctx, userMap)
}
