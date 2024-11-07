// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/r0busta/go-shopify-graphql/v8 (interfaces: InventoryService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/r0busta/go-shopify-graphql-model/v4/graph/model"
)

// MockInventoryService is a mock of InventoryService interface.
type MockInventoryService struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryServiceMockRecorder
}

// MockInventoryServiceMockRecorder is the mock recorder for MockInventoryService.
type MockInventoryServiceMockRecorder struct {
	mock *MockInventoryService
}

// NewMockInventoryService creates a new mock instance.
func NewMockInventoryService(ctrl *gomock.Controller) *MockInventoryService {
	mock := &MockInventoryService{ctrl: ctrl}
	mock.recorder = &MockInventoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryService) EXPECT() *MockInventoryServiceMockRecorder {
	return m.recorder
}

// ActivateInventory mocks base method.
func (m *MockInventoryService) ActivateInventory(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateInventory", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActivateInventory indicates an expected call of ActivateInventory.
func (mr *MockInventoryServiceMockRecorder) ActivateInventory(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateInventory", reflect.TypeOf((*MockInventoryService)(nil).ActivateInventory), arg0, arg1, arg2)
}

// Adjust mocks base method.
func (m *MockInventoryService) Adjust(arg0 context.Context, arg1 string, arg2 []model.InventoryAdjustQuantitiesInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Adjust", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Adjust indicates an expected call of Adjust.
func (mr *MockInventoryServiceMockRecorder) Adjust(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Adjust", reflect.TypeOf((*MockInventoryService)(nil).Adjust), arg0, arg1, arg2)
}

// AdjustQuantities mocks base method.
func (m *MockInventoryService) AdjustQuantities(arg0 context.Context, arg1, arg2 string, arg3 *string, arg4 []model.InventoryChangeInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdjustQuantities", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdjustQuantities indicates an expected call of AdjustQuantities.
func (mr *MockInventoryServiceMockRecorder) AdjustQuantities(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdjustQuantities", reflect.TypeOf((*MockInventoryService)(nil).AdjustQuantities), arg0, arg1, arg2, arg3, arg4)
}

// SetOnHandQuantities mocks base method.
func (m *MockInventoryService) SetOnHandQuantities(arg0 context.Context, arg1 string, arg2 *string, arg3 []model.InventorySetQuantityInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOnHandQuantities", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOnHandQuantities indicates an expected call of SetOnHandQuantities.
func (mr *MockInventoryServiceMockRecorder) SetOnHandQuantities(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOnHandQuantities", reflect.TypeOf((*MockInventoryService)(nil).SetOnHandQuantities), arg0, arg1, arg2, arg3)
}

// Update mocks base method.
func (m *MockInventoryService) Update(arg0 context.Context, arg1 string, arg2 model.InventoryItemInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockInventoryServiceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockInventoryService)(nil).Update), arg0, arg1, arg2)
}
