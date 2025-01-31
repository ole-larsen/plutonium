// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/db/repository/blogs.go
//
// Generated by this command:
//
//	mockgen -source=internal/storage/db/repository/blogs.go -destination=internal/storage/mocks/blogs_mock.go -package=mocks
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

// MockBlogsRepositoryInterface is a mock of BlogsRepositoryInterface interface.
type MockBlogsRepositoryInterface struct {
	isgomock struct{}
	ctrl     *gomock.Controller
	recorder *MockBlogsRepositoryInterfaceMockRecorder
}

// MockBlogsRepositoryInterfaceMockRecorder is the mock recorder for MockBlogsRepositoryInterface.
type MockBlogsRepositoryInterfaceMockRecorder struct {
	mock *MockBlogsRepositoryInterface
}

// NewMockBlogsRepositoryInterface creates a new mock instance.
func NewMockBlogsRepositoryInterface(ctrl *gomock.Controller) *MockBlogsRepositoryInterface {
	mock := &MockBlogsRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockBlogsRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlogsRepositoryInterface) EXPECT() *MockBlogsRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBlogsRepositoryInterface) Create(ctx context.Context, blogMap map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, blogMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockBlogsRepositoryInterfaceMockRecorder) Create(ctx, blogMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBlogsRepositoryInterface)(nil).Create), ctx, blogMap)
}

// GetBlogByID mocks base method.
func (m *MockBlogsRepositoryInterface) GetBlogByID(ctx context.Context, id int64) (*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlogByID", ctx, id)
	ret0, _ := ret[0].(*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlogByID indicates an expected call of GetBlogByID.
func (mr *MockBlogsRepositoryInterfaceMockRecorder) GetBlogByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlogByID", reflect.TypeOf((*MockBlogsRepositoryInterface)(nil).GetBlogByID), ctx, id)
}

// GetBlogs mocks base method.
func (m *MockBlogsRepositoryInterface) GetBlogs(ctx context.Context) ([]*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlogs", ctx)
	ret0, _ := ret[0].([]*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlogs indicates an expected call of GetBlogs.
func (mr *MockBlogsRepositoryInterfaceMockRecorder) GetBlogs(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlogs", reflect.TypeOf((*MockBlogsRepositoryInterface)(nil).GetBlogs), ctx)
}

// GetPublicBlogItem mocks base method.
func (m *MockBlogsRepositoryInterface) GetPublicBlogItem(ctx context.Context, slug string) (*models.PublicBlogItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicBlogItem", ctx, slug)
	ret0, _ := ret[0].(*models.PublicBlogItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicBlogItem indicates an expected call of GetPublicBlogItem.
func (mr *MockBlogsRepositoryInterfaceMockRecorder) GetPublicBlogItem(ctx, slug any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicBlogItem", reflect.TypeOf((*MockBlogsRepositoryInterface)(nil).GetPublicBlogItem), ctx, slug)
}

// GetPublicBlogs mocks base method.
func (m *MockBlogsRepositoryInterface) GetPublicBlogs(ctx context.Context) ([]*models.PublicBlogItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicBlogs", ctx)
	ret0, _ := ret[0].([]*models.PublicBlogItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicBlogs indicates an expected call of GetPublicBlogs.
func (mr *MockBlogsRepositoryInterfaceMockRecorder) GetPublicBlogs(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicBlogs", reflect.TypeOf((*MockBlogsRepositoryInterface)(nil).GetPublicBlogs), ctx)
}

// InnerDB mocks base method.
func (m *MockBlogsRepositoryInterface) InnerDB() *sqlx.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InnerDB")
	ret0, _ := ret[0].(*sqlx.DB)
	return ret0
}

// InnerDB indicates an expected call of InnerDB.
func (mr *MockBlogsRepositoryInterfaceMockRecorder) InnerDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InnerDB", reflect.TypeOf((*MockBlogsRepositoryInterface)(nil).InnerDB))
}

// Ping mocks base method.
func (m *MockBlogsRepositoryInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockBlogsRepositoryInterfaceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockBlogsRepositoryInterface)(nil).Ping))
}

// Update mocks base method.
func (m *MockBlogsRepositoryInterface) Update(ctx context.Context, blogMap map[string]any) ([]*models.Blog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, blogMap)
	ret0, _ := ret[0].([]*models.Blog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBlogsRepositoryInterfaceMockRecorder) Update(ctx, blogMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBlogsRepositoryInterface)(nil).Update), ctx, blogMap)
}
