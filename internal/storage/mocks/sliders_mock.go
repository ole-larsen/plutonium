// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/db/repository/sliders.go
//
// Generated by this command:
//
//	mockgen -source=internal/storage/db/repository/sliders.go -destination=internal/storage/mocks/sliders_mock.go -package=mocks
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

// MockSlidersRepositoryInterface is a mock of SlidersRepositoryInterface interface.
type MockSlidersRepositoryInterface struct {
	isgomock struct{}
	ctrl     *gomock.Controller
	recorder *MockSlidersRepositoryInterfaceMockRecorder
}

// MockSlidersRepositoryInterfaceMockRecorder is the mock recorder for MockSlidersRepositoryInterface.
type MockSlidersRepositoryInterfaceMockRecorder struct {
	mock *MockSlidersRepositoryInterface
}

// NewMockSlidersRepositoryInterface creates a new mock instance.
func NewMockSlidersRepositoryInterface(ctrl *gomock.Controller) *MockSlidersRepositoryInterface {
	mock := &MockSlidersRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockSlidersRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlidersRepositoryInterface) EXPECT() *MockSlidersRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSlidersRepositoryInterface) Create(ctx context.Context, sliderMap map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, sliderMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSlidersRepositoryInterfaceMockRecorder) Create(ctx, sliderMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSlidersRepositoryInterface)(nil).Create), ctx, sliderMap)
}

// GetSliderByID mocks base method.
func (m *MockSlidersRepositoryInterface) GetSliderByID(ctx context.Context, id int64) (*models.Slider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSliderByID", ctx, id)
	ret0, _ := ret[0].(*models.Slider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSliderByID indicates an expected call of GetSliderByID.
func (mr *MockSlidersRepositoryInterfaceMockRecorder) GetSliderByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSliderByID", reflect.TypeOf((*MockSlidersRepositoryInterface)(nil).GetSliderByID), ctx, id)
}

// GetSliderByProvider mocks base method.
func (m *MockSlidersRepositoryInterface) GetSliderByProvider(ctx context.Context, provider string) (*models.PublicSlider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSliderByProvider", ctx, provider)
	ret0, _ := ret[0].(*models.PublicSlider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSliderByProvider indicates an expected call of GetSliderByProvider.
func (mr *MockSlidersRepositoryInterfaceMockRecorder) GetSliderByProvider(ctx, provider any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSliderByProvider", reflect.TypeOf((*MockSlidersRepositoryInterface)(nil).GetSliderByProvider), ctx, provider)
}

// GetSliderByTitle mocks base method.
func (m *MockSlidersRepositoryInterface) GetSliderByTitle(ctx context.Context, title string) (*models.Slider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSliderByTitle", ctx, title)
	ret0, _ := ret[0].(*models.Slider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSliderByTitle indicates an expected call of GetSliderByTitle.
func (mr *MockSlidersRepositoryInterfaceMockRecorder) GetSliderByTitle(ctx, title any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSliderByTitle", reflect.TypeOf((*MockSlidersRepositoryInterface)(nil).GetSliderByTitle), ctx, title)
}

// GetSliders mocks base method.
func (m *MockSlidersRepositoryInterface) GetSliders(ctx context.Context) ([]*models.Slider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSliders", ctx)
	ret0, _ := ret[0].([]*models.Slider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSliders indicates an expected call of GetSliders.
func (mr *MockSlidersRepositoryInterfaceMockRecorder) GetSliders(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSliders", reflect.TypeOf((*MockSlidersRepositoryInterface)(nil).GetSliders), ctx)
}

// InnerDB mocks base method.
func (m *MockSlidersRepositoryInterface) InnerDB() *sqlx.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InnerDB")
	ret0, _ := ret[0].(*sqlx.DB)
	return ret0
}

// InnerDB indicates an expected call of InnerDB.
func (mr *MockSlidersRepositoryInterfaceMockRecorder) InnerDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InnerDB", reflect.TypeOf((*MockSlidersRepositoryInterface)(nil).InnerDB))
}

// Ping mocks base method.
func (m *MockSlidersRepositoryInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockSlidersRepositoryInterfaceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockSlidersRepositoryInterface)(nil).Ping))
}

// Update mocks base method.
func (m *MockSlidersRepositoryInterface) Update(ctx context.Context, sliderMap map[string]any) ([]*models.Slider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, sliderMap)
	ret0, _ := ret[0].([]*models.Slider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockSlidersRepositoryInterfaceMockRecorder) Update(ctx, sliderMap any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSlidersRepositoryInterface)(nil).Update), ctx, sliderMap)
}
