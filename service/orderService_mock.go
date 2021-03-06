// Code generated by MockGen. DO NOT EDIT.
// Source: orderService.go

// Package service is a generated GoMock package.
package service

import (
	model "demo01/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderService is a mock of OrderService interface.
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService.
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance.
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderService) CreateOrder(orderNo, userName string, amount float64, status, fileUrl string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", orderNo, userName, amount, status, fileUrl)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderServiceMockRecorder) CreateOrder(orderNo, userName, amount, status, fileUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderService)(nil).CreateOrder), orderNo, userName, amount, status, fileUrl)
}

// DeleteOrderByOrderNo mocks base method.
func (m *MockOrderService) DeleteOrderByOrderNo(orderNo string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderByOrderNo", orderNo)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderByOrderNo indicates an expected call of DeleteOrderByOrderNo.
func (mr *MockOrderServiceMockRecorder) DeleteOrderByOrderNo(orderNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderByOrderNo", reflect.TypeOf((*MockOrderService)(nil).DeleteOrderByOrderNo), orderNo)
}

// GetAllOrder mocks base method.
func (m *MockOrderService) GetAllOrder() ([]*model.DemoOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrder")
	ret0, _ := ret[0].([]*model.DemoOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrder indicates an expected call of GetAllOrder.
func (mr *MockOrderServiceMockRecorder) GetAllOrder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrder", reflect.TypeOf((*MockOrderService)(nil).GetAllOrder))
}

// GetOrderByOrderNo mocks base method.
func (m *MockOrderService) GetOrderByOrderNo(orderNo string) (*model.DemoOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByOrderNo", orderNo)
	ret0, _ := ret[0].(*model.DemoOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByOrderNo indicates an expected call of GetOrderByOrderNo.
func (mr *MockOrderServiceMockRecorder) GetOrderByOrderNo(orderNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByOrderNo", reflect.TypeOf((*MockOrderService)(nil).GetOrderByOrderNo), orderNo)
}

// UpdatesOrder mocks base method.
func (m *MockOrderService) UpdatesOrder(orderNo string, amount float64, status, fileUrl string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatesOrder", orderNo, amount, status, fileUrl)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatesOrder indicates an expected call of UpdatesOrder.
func (mr *MockOrderServiceMockRecorder) UpdatesOrder(orderNo, amount, status, fileUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatesOrder", reflect.TypeOf((*MockOrderService)(nil).UpdatesOrder), orderNo, amount, status, fileUrl)
}
