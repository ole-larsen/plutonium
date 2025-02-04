// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/db/repository/authors.go
//
// Generated by this command:
//
//	mockgen -source=internal/storage/db/repository/authors.go -destination=internal/storage/mocks/authors_mock.go -package=mocks
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

// MockAuthorsRepositoryInterface is a mock of AuthorsRepositoryInterface interface.
type MockAuthorsRepositoryInterface struct {
	isgomock struct{}
	ctrl     *gomock.Controller
	recorder *MockAuthorsRepositoryInterfaceMockRecorder
}

// MockAuthorsRepositoryInterfaceMockRecorder is the mock recorder for MockAuthorsRepositoryInterface.
type MockAuthorsRepositoryInterfaceMockRecorder struct {
	mock *MockAuthorsRepositoryInterface
}

// NewMockAuthorsRepositoryInterface creates a new mock instance.
func NewMockAuthorsRepositoryInterface(ctrl *gomock.Controller) *MockAuthorsRepositoryInterface {
	mock := &MockAuthorsRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockAuthorsRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorsRepositoryInterface) EXPECT() *MockAuthorsRepositoryInterfaceMockRecorder {
	return m.recorder
}

// BindSocial mocks base method.
func (m *MockAuthorsRepositoryInterface) BindSocial(ctx context.Context, socialMap map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindSocial", ctx, socialMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindSocial indicates an expected call of BindSocial.
func (mr *MockAuthorsRepositoryInterfaceMockRecorder) BindSocial(ctx, socialMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindSocial", reflect.TypeOf((*MockAuthorsRepositoryInterface)(nil).BindSocial), ctx, socialMap)
}

// BindWallet mocks base method.
func (m *MockAuthorsRepositoryInterface) BindWallet(ctx context.Context, walletMap map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindWallet", ctx, walletMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindWallet indicates an expected call of BindWallet.
func (mr *MockAuthorsRepositoryInterfaceMockRecorder) BindWallet(ctx, walletMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindWallet", reflect.TypeOf((*MockAuthorsRepositoryInterface)(nil).BindWallet), ctx, walletMap)
}

// Create mocks base method.
func (m *MockAuthorsRepositoryInterface) Create(ctx context.Context, authorMap map[string]any, socials []*models.Social, wallets []*models.Wallet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, authorMap, socials, wallets)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAuthorsRepositoryInterfaceMockRecorder) Create(ctx, authorMap, socials, wallets any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthorsRepositoryInterface)(nil).Create), ctx, authorMap, socials, wallets)
}

// GetAuthorByID mocks base method.
func (m *MockAuthorsRepositoryInterface) GetAuthorByID(ctx context.Context, id int64) (*models.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorByID", ctx, id)
	ret0, _ := ret[0].(*models.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorByID indicates an expected call of GetAuthorByID.
func (mr *MockAuthorsRepositoryInterfaceMockRecorder) GetAuthorByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorByID", reflect.TypeOf((*MockAuthorsRepositoryInterface)(nil).GetAuthorByID), ctx, id)
}

// GetAuthors mocks base method.
func (m *MockAuthorsRepositoryInterface) GetAuthors(ctx context.Context) ([]*models.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthors", ctx)
	ret0, _ := ret[0].([]*models.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthors indicates an expected call of GetAuthors.
func (mr *MockAuthorsRepositoryInterfaceMockRecorder) GetAuthors(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthors", reflect.TypeOf((*MockAuthorsRepositoryInterface)(nil).GetAuthors), ctx)
}

// GetPublicAuthor mocks base method.
func (m *MockAuthorsRepositoryInterface) GetPublicAuthor(ctx context.Context, slug string) (*models.PublicAuthorItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicAuthor", ctx, slug)
	ret0, _ := ret[0].(*models.PublicAuthorItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicAuthor indicates an expected call of GetPublicAuthor.
func (mr *MockAuthorsRepositoryInterfaceMockRecorder) GetPublicAuthor(ctx, slug any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicAuthor", reflect.TypeOf((*MockAuthorsRepositoryInterface)(nil).GetPublicAuthor), ctx, slug)
}

// GetPublicAuthors mocks base method.
func (m *MockAuthorsRepositoryInterface) GetPublicAuthors(ctx context.Context) ([]*models.PublicAuthorItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicAuthors", ctx)
	ret0, _ := ret[0].([]*models.PublicAuthorItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicAuthors indicates an expected call of GetPublicAuthors.
func (mr *MockAuthorsRepositoryInterfaceMockRecorder) GetPublicAuthors(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicAuthors", reflect.TypeOf((*MockAuthorsRepositoryInterface)(nil).GetPublicAuthors), ctx)
}

// InnerDB mocks base method.
func (m *MockAuthorsRepositoryInterface) InnerDB() *sqlx.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InnerDB")
	ret0, _ := ret[0].(*sqlx.DB)
	return ret0
}

// InnerDB indicates an expected call of InnerDB.
func (mr *MockAuthorsRepositoryInterfaceMockRecorder) InnerDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InnerDB", reflect.TypeOf((*MockAuthorsRepositoryInterface)(nil).InnerDB))
}

// Ping mocks base method.
func (m *MockAuthorsRepositoryInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockAuthorsRepositoryInterfaceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockAuthorsRepositoryInterface)(nil).Ping))
}

// Update mocks base method.
func (m *MockAuthorsRepositoryInterface) Update(ctx context.Context, authorMap map[string]any, socials []*models.Social, wallets []*models.Wallet) ([]*models.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, authorMap, socials, wallets)
	ret0, _ := ret[0].([]*models.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAuthorsRepositoryInterfaceMockRecorder) Update(ctx, authorMap, socials, wallets any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthorsRepositoryInterface)(nil).Update), ctx, authorMap, socials, wallets)
}
