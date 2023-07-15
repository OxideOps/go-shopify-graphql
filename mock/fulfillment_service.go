// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/r0busta/go-shopify-graphql/v8 (interfaces: FulfillmentService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/r0busta/go-shopify-graphql-model/v3/graph/model"
)

// MockFulfillmentService is a mock of FulfillmentService interface.
type MockFulfillmentService struct {
	ctrl     *gomock.Controller
	recorder *MockFulfillmentServiceMockRecorder
}

// MockFulfillmentServiceMockRecorder is the mock recorder for MockFulfillmentService.
type MockFulfillmentServiceMockRecorder struct {
	mock *MockFulfillmentService
}

// NewMockFulfillmentService creates a new mock instance.
func NewMockFulfillmentService(ctrl *gomock.Controller) *MockFulfillmentService {
	mock := &MockFulfillmentService{ctrl: ctrl}
	mock.recorder = &MockFulfillmentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFulfillmentService) EXPECT() *MockFulfillmentServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFulfillmentService) Create(arg0 context.Context, arg1 model.FulfillmentV2Input) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockFulfillmentServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFulfillmentService)(nil).Create), arg0, arg1)
}
