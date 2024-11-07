// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/r0busta/go-shopify-graphql/v8 (interfaces: BulkOperationService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	model "github.com/r0busta/go-shopify-graphql-model/v4/graph/model"
)

// MockBulkOperationService is a mock of BulkOperationService interface.
type MockBulkOperationService struct {
	ctrl     *gomock.Controller
	recorder *MockBulkOperationServiceMockRecorder
}

// MockBulkOperationServiceMockRecorder is the mock recorder for MockBulkOperationService.
type MockBulkOperationServiceMockRecorder struct {
	mock *MockBulkOperationService
}

// NewMockBulkOperationService creates a new mock instance.
func NewMockBulkOperationService(ctrl *gomock.Controller) *MockBulkOperationService {
	mock := &MockBulkOperationService{ctrl: ctrl}
	mock.recorder = &MockBulkOperationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBulkOperationService) EXPECT() *MockBulkOperationServiceMockRecorder {
	return m.recorder
}

// BulkQuery mocks base method.
func (m *MockBulkOperationService) BulkQuery(arg0 context.Context, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkQuery", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkQuery indicates an expected call of BulkQuery.
func (mr *MockBulkOperationServiceMockRecorder) BulkQuery(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkQuery", reflect.TypeOf((*MockBulkOperationService)(nil).BulkQuery), arg0, arg1, arg2)
}

// CancelRunningBulkQuery mocks base method.
func (m *MockBulkOperationService) CancelRunningBulkQuery(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelRunningBulkQuery", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelRunningBulkQuery indicates an expected call of CancelRunningBulkQuery.
func (mr *MockBulkOperationServiceMockRecorder) CancelRunningBulkQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRunningBulkQuery", reflect.TypeOf((*MockBulkOperationService)(nil).CancelRunningBulkQuery), arg0)
}

// GetCurrentBulkQuery mocks base method.
func (m *MockBulkOperationService) GetCurrentBulkQuery(arg0 context.Context) (*model.BulkOperation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentBulkQuery", arg0)
	ret0, _ := ret[0].(*model.BulkOperation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentBulkQuery indicates an expected call of GetCurrentBulkQuery.
func (mr *MockBulkOperationServiceMockRecorder) GetCurrentBulkQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentBulkQuery", reflect.TypeOf((*MockBulkOperationService)(nil).GetCurrentBulkQuery), arg0)
}

// GetCurrentBulkQueryResultURL mocks base method.
func (m *MockBulkOperationService) GetCurrentBulkQueryResultURL(arg0 context.Context) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentBulkQueryResultURL", arg0)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentBulkQueryResultURL indicates an expected call of GetCurrentBulkQueryResultURL.
func (mr *MockBulkOperationServiceMockRecorder) GetCurrentBulkQueryResultURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentBulkQueryResultURL", reflect.TypeOf((*MockBulkOperationService)(nil).GetCurrentBulkQueryResultURL), arg0)
}

// PostBulkQuery mocks base method.
func (m *MockBulkOperationService) PostBulkQuery(arg0 context.Context, arg1 string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostBulkQuery", arg0, arg1)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostBulkQuery indicates an expected call of PostBulkQuery.
func (mr *MockBulkOperationServiceMockRecorder) PostBulkQuery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBulkQuery", reflect.TypeOf((*MockBulkOperationService)(nil).PostBulkQuery), arg0, arg1)
}

// ShouldGetBulkQueryResultURL mocks base method.
func (m *MockBulkOperationService) ShouldGetBulkQueryResultURL(arg0 context.Context, arg1 *string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldGetBulkQueryResultURL", arg0, arg1)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldGetBulkQueryResultURL indicates an expected call of ShouldGetBulkQueryResultURL.
func (mr *MockBulkOperationServiceMockRecorder) ShouldGetBulkQueryResultURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldGetBulkQueryResultURL", reflect.TypeOf((*MockBulkOperationService)(nil).ShouldGetBulkQueryResultURL), arg0, arg1)
}

// WaitForCurrentBulkQuery mocks base method.
func (m *MockBulkOperationService) WaitForCurrentBulkQuery(arg0 context.Context, arg1 time.Duration) (*model.BulkOperation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForCurrentBulkQuery", arg0, arg1)
	ret0, _ := ret[0].(*model.BulkOperation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForCurrentBulkQuery indicates an expected call of WaitForCurrentBulkQuery.
func (mr *MockBulkOperationServiceMockRecorder) WaitForCurrentBulkQuery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForCurrentBulkQuery", reflect.TypeOf((*MockBulkOperationService)(nil).WaitForCurrentBulkQuery), arg0, arg1)
}
