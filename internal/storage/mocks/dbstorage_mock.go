// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/dbstorage.go
//
// Generated by this command:
//
//	mockgen -source=internal/storage/dbstorage.go -destination=internal/storage/mocks/dbstorage_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	sqlx "github.com/jmoiron/sqlx"
	repository "github.com/ole-larsen/plutonium/internal/storage/db/repository"
	gomock "go.uber.org/mock/gomock"
)

// MockDBStorageInterface is a mock of DBStorageInterface interface.
type MockDBStorageInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDBStorageInterfaceMockRecorder
	isgomock struct{}
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

// ConnectRepository mocks base method.
func (m *MockDBStorageInterface) ConnectRepository(name string, sqlxDB *sqlx.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectRepository", name, sqlxDB)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectRepository indicates an expected call of ConnectRepository.
func (mr *MockDBStorageInterfaceMockRecorder) ConnectRepository(name, sqlxDB any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).ConnectRepository), name, sqlxDB)
}

// GetAuthorsRepository mocks base method.
func (m *MockDBStorageInterface) GetAuthorsRepository() *repository.AuthorsRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthorsRepository")
	ret0, _ := ret[0].(*repository.AuthorsRepository)
	return ret0
}

// GetAuthorsRepository indicates an expected call of GetAuthorsRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetAuthorsRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorsRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetAuthorsRepository))
}

// GetBlogsRepository mocks base method.
func (m *MockDBStorageInterface) GetBlogsRepository() *repository.BlogsRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlogsRepository")
	ret0, _ := ret[0].(*repository.BlogsRepository)
	return ret0
}

// GetBlogsRepository indicates an expected call of GetBlogsRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetBlogsRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlogsRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetBlogsRepository))
}

// GetCategoriesRepository mocks base method.
func (m *MockDBStorageInterface) GetCategoriesRepository() *repository.CategoriesRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoriesRepository")
	ret0, _ := ret[0].(*repository.CategoriesRepository)
	return ret0
}

// GetCategoriesRepository indicates an expected call of GetCategoriesRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetCategoriesRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoriesRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetCategoriesRepository))
}

// GetContactFormsRepository mocks base method.
func (m *MockDBStorageInterface) GetContactFormsRepository() *repository.ContactFormRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactFormsRepository")
	ret0, _ := ret[0].(*repository.ContactFormRepository)
	return ret0
}

// GetContactFormsRepository indicates an expected call of GetContactFormsRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetContactFormsRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactFormsRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetContactFormsRepository))
}

// GetContactsRepository mocks base method.
func (m *MockDBStorageInterface) GetContactsRepository() *repository.ContactsRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactsRepository")
	ret0, _ := ret[0].(*repository.ContactsRepository)
	return ret0
}

// GetContactsRepository indicates an expected call of GetContactsRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetContactsRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactsRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetContactsRepository))
}

// GetContractsRepository mocks base method.
func (m *MockDBStorageInterface) GetContractsRepository() *repository.ContractsRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractsRepository")
	ret0, _ := ret[0].(*repository.ContractsRepository)
	return ret0
}

// GetContractsRepository indicates an expected call of GetContractsRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetContractsRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractsRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetContractsRepository))
}

// GetCreateAndSellRepository mocks base method.
func (m *MockDBStorageInterface) GetCreateAndSellRepository() *repository.CreateAndSellRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreateAndSellRepository")
	ret0, _ := ret[0].(*repository.CreateAndSellRepository)
	return ret0
}

// GetCreateAndSellRepository indicates an expected call of GetCreateAndSellRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetCreateAndSellRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreateAndSellRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetCreateAndSellRepository))
}

// GetFaqsRepository mocks base method.
func (m *MockDBStorageInterface) GetFaqsRepository() *repository.FaqsRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFaqsRepository")
	ret0, _ := ret[0].(*repository.FaqsRepository)
	return ret0
}

// GetFaqsRepository indicates an expected call of GetFaqsRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetFaqsRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFaqsRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetFaqsRepository))
}

// GetFilesRepository mocks base method.
func (m *MockDBStorageInterface) GetFilesRepository() *repository.FilesRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilesRepository")
	ret0, _ := ret[0].(*repository.FilesRepository)
	return ret0
}

// GetFilesRepository indicates an expected call of GetFilesRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetFilesRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilesRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetFilesRepository))
}

// GetHelpCenterRepository mocks base method.
func (m *MockDBStorageInterface) GetHelpCenterRepository() *repository.HelpCenterRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHelpCenterRepository")
	ret0, _ := ret[0].(*repository.HelpCenterRepository)
	return ret0
}

// GetHelpCenterRepository indicates an expected call of GetHelpCenterRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetHelpCenterRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHelpCenterRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetHelpCenterRepository))
}

// GetMenusRepository mocks base method.
func (m *MockDBStorageInterface) GetMenusRepository() *repository.MenusRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenusRepository")
	ret0, _ := ret[0].(*repository.MenusRepository)
	return ret0
}

// GetMenusRepository indicates an expected call of GetMenusRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetMenusRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenusRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetMenusRepository))
}

// GetPagesRepository mocks base method.
func (m *MockDBStorageInterface) GetPagesRepository() *repository.PagesRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPagesRepository")
	ret0, _ := ret[0].(*repository.PagesRepository)
	return ret0
}

// GetPagesRepository indicates an expected call of GetPagesRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetPagesRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPagesRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetPagesRepository))
}

// GetSlidersRepository mocks base method.
func (m *MockDBStorageInterface) GetSlidersRepository() *repository.SlidersRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlidersRepository")
	ret0, _ := ret[0].(*repository.SlidersRepository)
	return ret0
}

// GetSlidersRepository indicates an expected call of GetSlidersRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetSlidersRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlidersRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetSlidersRepository))
}

// GetTagsRepository mocks base method.
func (m *MockDBStorageInterface) GetTagsRepository() *repository.TagsRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagsRepository")
	ret0, _ := ret[0].(*repository.TagsRepository)
	return ret0
}

// GetTagsRepository indicates an expected call of GetTagsRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetTagsRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagsRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetTagsRepository))
}

// GetUsersRepository mocks base method.
func (m *MockDBStorageInterface) GetUsersRepository() *repository.UsersRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersRepository")
	ret0, _ := ret[0].(*repository.UsersRepository)
	return ret0
}

// GetUsersRepository indicates an expected call of GetUsersRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetUsersRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetUsersRepository))
}

// GetWalletsRepository mocks base method.
func (m *MockDBStorageInterface) GetWalletsRepository() *repository.WalletsRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletsRepository")
	ret0, _ := ret[0].(*repository.WalletsRepository)
	return ret0
}

// GetWalletsRepository indicates an expected call of GetWalletsRepository.
func (mr *MockDBStorageInterfaceMockRecorder) GetWalletsRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletsRepository", reflect.TypeOf((*MockDBStorageInterface)(nil).GetWalletsRepository))
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
func (mr *MockDBStorageInterfaceMockRecorder) Init(ctx, sqlxDB any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockDBStorageInterface)(nil).Init), ctx, sqlxDB)
}
