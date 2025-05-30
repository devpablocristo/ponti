// Code generated by MockGen. DO NOT EDIT.
// Source: internal/field/ports.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/alphacodinggroup/ponti-backend/projects/ponti-api/internal/field/usecases/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCases is a mock of UseCases interface.
type MockUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockUseCasesMockRecorder
}

// MockUseCasesMockRecorder is the mock recorder for MockUseCases.
type MockUseCasesMockRecorder struct {
	mock *MockUseCases
}

// NewMockUseCases creates a new mock instance.
func NewMockUseCases(ctrl *gomock.Controller) *MockUseCases {
	mock := &MockUseCases{ctrl: ctrl}
	mock.recorder = &MockUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCases) EXPECT() *MockUseCasesMockRecorder {
	return m.recorder
}

// CreateField mocks base method.
func (m *MockUseCases) CreateField(ctx context.Context, f *domain.Field) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateField", ctx, f)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateField indicates an expected call of CreateField.
func (mr *MockUseCasesMockRecorder) CreateField(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateField", reflect.TypeOf((*MockUseCases)(nil).CreateField), ctx, f)
}

// DeleteField mocks base method.
func (m *MockUseCases) DeleteField(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteField", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteField indicates an expected call of DeleteField.
func (mr *MockUseCasesMockRecorder) DeleteField(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteField", reflect.TypeOf((*MockUseCases)(nil).DeleteField), ctx, id)
}

// GetField mocks base method.
func (m *MockUseCases) GetField(ctx context.Context, id int64) (*domain.Field, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetField", ctx, id)
	ret0, _ := ret[0].(*domain.Field)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetField indicates an expected call of GetField.
func (mr *MockUseCasesMockRecorder) GetField(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetField", reflect.TypeOf((*MockUseCases)(nil).GetField), ctx, id)
}

// ListFields mocks base method.
func (m *MockUseCases) ListFields(ctx context.Context) ([]domain.Field, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFields", ctx)
	ret0, _ := ret[0].([]domain.Field)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFields indicates an expected call of ListFields.
func (mr *MockUseCasesMockRecorder) ListFields(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFields", reflect.TypeOf((*MockUseCases)(nil).ListFields), ctx)
}

// UpdateField mocks base method.
func (m *MockUseCases) UpdateField(ctx context.Context, f *domain.Field) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateField", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateField indicates an expected call of UpdateField.
func (mr *MockUseCasesMockRecorder) UpdateField(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateField", reflect.TypeOf((*MockUseCases)(nil).UpdateField), ctx, f)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateField mocks base method.
func (m *MockRepository) CreateField(ctx context.Context, f *domain.Field) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateField", ctx, f)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateField indicates an expected call of CreateField.
func (mr *MockRepositoryMockRecorder) CreateField(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateField", reflect.TypeOf((*MockRepository)(nil).CreateField), ctx, f)
}

// DeleteField mocks base method.
func (m *MockRepository) DeleteField(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteField", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteField indicates an expected call of DeleteField.
func (mr *MockRepositoryMockRecorder) DeleteField(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteField", reflect.TypeOf((*MockRepository)(nil).DeleteField), ctx, id)
}

// GetField mocks base method.
func (m *MockRepository) GetField(ctx context.Context, id int64) (*domain.Field, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetField", ctx, id)
	ret0, _ := ret[0].(*domain.Field)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetField indicates an expected call of GetField.
func (mr *MockRepositoryMockRecorder) GetField(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetField", reflect.TypeOf((*MockRepository)(nil).GetField), ctx, id)
}

// ListFields mocks base method.
func (m *MockRepository) ListFields(ctx context.Context) ([]domain.Field, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFields", ctx)
	ret0, _ := ret[0].([]domain.Field)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFields indicates an expected call of ListFields.
func (mr *MockRepositoryMockRecorder) ListFields(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFields", reflect.TypeOf((*MockRepository)(nil).ListFields), ctx)
}

// UpdateField mocks base method.
func (m *MockRepository) UpdateField(ctx context.Context, f *domain.Field) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateField", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateField indicates an expected call of UpdateField.
func (mr *MockRepositoryMockRecorder) UpdateField(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateField", reflect.TypeOf((*MockRepository)(nil).UpdateField), ctx, f)
}
