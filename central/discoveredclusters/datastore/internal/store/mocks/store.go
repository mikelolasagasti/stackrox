// Code generated by MockGen. DO NOT EDIT.
// Source: store.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/store.go -source store.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteByQuery mocks base method.
func (m *MockStore) DeleteByQuery(ctx context.Context, query *v1.Query) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByQuery", ctx, query)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByQuery indicates an expected call of DeleteByQuery.
func (mr *MockStoreMockRecorder) DeleteByQuery(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByQuery", reflect.TypeOf((*MockStore)(nil).DeleteByQuery), ctx, query)
}

// Get mocks base method.
func (m *MockStore) Get(ctx context.Context, id string) (*storage.DiscoveredCluster, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.DiscoveredCluster)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), ctx, id)
}

// GetByQuery mocks base method.
func (m *MockStore) GetByQuery(ctx context.Context, query *v1.Query) ([]*storage.DiscoveredCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByQuery", ctx, query)
	ret0, _ := ret[0].([]*storage.DiscoveredCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByQuery indicates an expected call of GetByQuery.
func (mr *MockStoreMockRecorder) GetByQuery(ctx, query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByQuery", reflect.TypeOf((*MockStore)(nil).GetByQuery), ctx, query)
}

// UpsertMany mocks base method.
func (m *MockStore) UpsertMany(ctx context.Context, objs []*storage.DiscoveredCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertMany", ctx, objs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertMany indicates an expected call of UpsertMany.
func (mr *MockStoreMockRecorder) UpsertMany(ctx, objs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertMany", reflect.TypeOf((*MockStore)(nil).UpsertMany), ctx, objs)
}
