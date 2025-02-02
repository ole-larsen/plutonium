// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/db/repository/wallets.go
//
// Generated by this command:
//
//	mockgen -source=internal/storage/db/repository/wallets.go -destination=internal/storage/mocks/wallets_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	sqlx "github.com/jmoiron/sqlx"
	models "github.com/ole-larsen/plutonium/models"
	gomock "go.uber.org/mock/gomock"
)

// MockWalletsRepositoryInterface is a mock of WalletsRepositoryInterface interface.
type MockWalletsRepositoryInterface struct {
	isgomock struct{}
	ctrl     *gomock.Controller
	recorder *MockWalletsRepositoryInterfaceMockRecorder
}

// MockWalletsRepositoryInterfaceMockRecorder is the mock recorder for MockWalletsRepositoryInterface.
type MockWalletsRepositoryInterfaceMockRecorder struct {
	mock *MockWalletsRepositoryInterface
}

// NewMockWalletsRepositoryInterface creates a new mock instance.
func NewMockWalletsRepositoryInterface(ctrl *gomock.Controller) *MockWalletsRepositoryInterface {
	mock := &MockWalletsRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockWalletsRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletsRepositoryInterface) EXPECT() *MockWalletsRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWalletsRepositoryInterface) Create(ctx context.Context, walletMap map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, walletMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockWalletsRepositoryInterfaceMockRecorder) Create(ctx, walletMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWalletsRepositoryInterface)(nil).Create), ctx, walletMap)
}

// GetPublicWalletConnect mocks base method.
func (m *MockWalletsRepositoryInterface) GetPublicWalletConnect(ctx context.Context) ([]*models.PublicWalletConnectItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicWalletConnect", ctx)
	ret0, _ := ret[0].([]*models.PublicWalletConnectItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicWalletConnect indicates an expected call of GetPublicWalletConnect.
func (mr *MockWalletsRepositoryInterfaceMockRecorder) GetPublicWalletConnect(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicWalletConnect", reflect.TypeOf((*MockWalletsRepositoryInterface)(nil).GetPublicWalletConnect), ctx)
}

// GetWalletByID mocks base method.
func (m *MockWalletsRepositoryInterface) GetWalletByID(ctx context.Context, id int64) (*models.WalletConnect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletByID", ctx, id)
	ret0, _ := ret[0].(*models.WalletConnect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletByID indicates an expected call of GetWalletByID.
func (mr *MockWalletsRepositoryInterfaceMockRecorder) GetWalletByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletByID", reflect.TypeOf((*MockWalletsRepositoryInterface)(nil).GetWalletByID), ctx, id)
}

// GetWallets mocks base method.
func (m *MockWalletsRepositoryInterface) GetWallets(ctx context.Context) ([]*models.WalletConnect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallets", ctx)
	ret0, _ := ret[0].([]*models.WalletConnect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallets indicates an expected call of GetWallets.
func (mr *MockWalletsRepositoryInterfaceMockRecorder) GetWallets(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallets", reflect.TypeOf((*MockWalletsRepositoryInterface)(nil).GetWallets), ctx)
}

// InnerDB mocks base method.
func (m *MockWalletsRepositoryInterface) InnerDB() *sqlx.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InnerDB")
	ret0, _ := ret[0].(*sqlx.DB)
	return ret0
}

// InnerDB indicates an expected call of InnerDB.
func (mr *MockWalletsRepositoryInterfaceMockRecorder) InnerDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InnerDB", reflect.TypeOf((*MockWalletsRepositoryInterface)(nil).InnerDB))
}

// Ping mocks base method.
func (m *MockWalletsRepositoryInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockWalletsRepositoryInterfaceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockWalletsRepositoryInterface)(nil).Ping))
}

// Update mocks base method.
func (m *MockWalletsRepositoryInterface) Update(ctx context.Context, walletMap map[string]any) ([]*models.WalletConnect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, walletMap)
	ret0, _ := ret[0].([]*models.WalletConnect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockWalletsRepositoryInterfaceMockRecorder) Update(ctx, walletMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWalletsRepositoryInterface)(nil).Update), ctx, walletMap)
}
