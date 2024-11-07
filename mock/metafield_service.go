// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/r0busta/go-shopify-graphql/v8 (interfaces: MetafieldService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/r0busta/go-shopify-graphql-model/v4/graph/model"
)

// MockMetafieldService is a mock of MetafieldService interface.
type MockMetafieldService struct {
	ctrl     *gomock.Controller
	recorder *MockMetafieldServiceMockRecorder
}

// MockMetafieldServiceMockRecorder is the mock recorder for MockMetafieldService.
type MockMetafieldServiceMockRecorder struct {
	mock *MockMetafieldService
}

// NewMockMetafieldService creates a new mock instance.
func NewMockMetafieldService(ctrl *gomock.Controller) *MockMetafieldService {
	mock := &MockMetafieldService{ctrl: ctrl}
	mock.recorder = &MockMetafieldServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetafieldService) EXPECT() *MockMetafieldServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMetafieldService) Delete(arg0 context.Context, arg1 model.MetafieldIdentifierInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMetafieldServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMetafieldService)(nil).Delete), arg0, arg1)
}

// DeleteBulk mocks base method.
func (m *MockMetafieldService) DeleteBulk(arg0 context.Context, arg1 []model.MetafieldIdentifierInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBulk", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBulk indicates an expected call of DeleteBulk.
func (mr *MockMetafieldServiceMockRecorder) DeleteBulk(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBulk", reflect.TypeOf((*MockMetafieldService)(nil).DeleteBulk), arg0, arg1)
}

// GetShopMetafieldByKey mocks base method.
func (m *MockMetafieldService) GetShopMetafieldByKey(arg0 context.Context, arg1, arg2 string) (*model.Metafield, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShopMetafieldByKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.Metafield)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShopMetafieldByKey indicates an expected call of GetShopMetafieldByKey.
func (mr *MockMetafieldServiceMockRecorder) GetShopMetafieldByKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShopMetafieldByKey", reflect.TypeOf((*MockMetafieldService)(nil).GetShopMetafieldByKey), arg0, arg1, arg2)
}

// ListAllShopMetafields mocks base method.
func (m *MockMetafieldService) ListAllShopMetafields(arg0 context.Context) ([]model.Metafield, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllShopMetafields", arg0)
	ret0, _ := ret[0].([]model.Metafield)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllShopMetafields indicates an expected call of ListAllShopMetafields.
func (mr *MockMetafieldServiceMockRecorder) ListAllShopMetafields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllShopMetafields", reflect.TypeOf((*MockMetafieldService)(nil).ListAllShopMetafields), arg0)
}

// ListShopMetafieldsByNamespace mocks base method.
func (m *MockMetafieldService) ListShopMetafieldsByNamespace(arg0 context.Context, arg1 string) ([]model.Metafield, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListShopMetafieldsByNamespace", arg0, arg1)
	ret0, _ := ret[0].([]model.Metafield)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListShopMetafieldsByNamespace indicates an expected call of ListShopMetafieldsByNamespace.
func (mr *MockMetafieldServiceMockRecorder) ListShopMetafieldsByNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShopMetafieldsByNamespace", reflect.TypeOf((*MockMetafieldService)(nil).ListShopMetafieldsByNamespace), arg0, arg1)
}
