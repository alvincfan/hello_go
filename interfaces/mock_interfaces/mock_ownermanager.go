// Code generated by MockGen. DO NOT EDIT.
// Source: ownerManager.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	gomock "github.com/golang/mock/gomock"
	domain "hello_go/domain"
	reflect "reflect"
)

// MockOwnerManager is a mock of OwnerManager interface
type MockOwnerManager struct {
	ctrl     *gomock.Controller
	recorder *MockOwnerManagerMockRecorder
}

// MockOwnerManagerMockRecorder is the mock recorder for MockOwnerManager
type MockOwnerManagerMockRecorder struct {
	mock *MockOwnerManager
}

// NewMockOwnerManager creates a new mock instance
func NewMockOwnerManager(ctrl *gomock.Controller) *MockOwnerManager {
	mock := &MockOwnerManager{ctrl: ctrl}
	mock.recorder = &MockOwnerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOwnerManager) EXPECT() *MockOwnerManagerMockRecorder {
	return m.recorder
}

// GetOwners mocks base method
func (m *MockOwnerManager) GetOwners() map[string]*domain.Owner {
	ret := m.ctrl.Call(m, "GetOwners")
	ret0, _ := ret[0].(map[string]*domain.Owner)
	return ret0
}

// GetOwners indicates an expected call of GetOwners
func (mr *MockOwnerManagerMockRecorder) GetOwners() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwners", reflect.TypeOf((*MockOwnerManager)(nil).GetOwners))
}

// CreateOwner mocks base method
func (m *MockOwnerManager) CreateOwner(ownerID, name string, address *domain.Address) (*domain.Owner, error) {
	ret := m.ctrl.Call(m, "CreateOwner", ownerID, name, address)
	ret0, _ := ret[0].(*domain.Owner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOwner indicates an expected call of CreateOwner
func (mr *MockOwnerManagerMockRecorder) CreateOwner(ownerID, name, address interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOwner", reflect.TypeOf((*MockOwnerManager)(nil).CreateOwner), ownerID, name, address)
}

// GetOwner mocks base method
func (m *MockOwnerManager) GetOwner(ownerID string) *domain.Owner {
	ret := m.ctrl.Call(m, "GetOwner", ownerID)
	ret0, _ := ret[0].(*domain.Owner)
	return ret0
}

// GetOwner indicates an expected call of GetOwner
func (mr *MockOwnerManagerMockRecorder) GetOwner(ownerID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwner", reflect.TypeOf((*MockOwnerManager)(nil).GetOwner), ownerID)
}

// UpdateOwner mocks base method
func (m *MockOwnerManager) UpdateOwner(owner *domain.Owner) bool {
	ret := m.ctrl.Call(m, "UpdateOwner", owner)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateOwner indicates an expected call of UpdateOwner
func (mr *MockOwnerManagerMockRecorder) UpdateOwner(owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOwner", reflect.TypeOf((*MockOwnerManager)(nil).UpdateOwner), owner)
}

// DeleteOwner mocks base method
func (m *MockOwnerManager) DeleteOwner(owner *domain.Owner) bool {
	ret := m.ctrl.Call(m, "DeleteOwner", owner)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteOwner indicates an expected call of DeleteOwner
func (mr *MockOwnerManagerMockRecorder) DeleteOwner(owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOwner", reflect.TypeOf((*MockOwnerManager)(nil).DeleteOwner), owner)
}

// IsExisting mocks base method
func (m *MockOwnerManager) IsExisting(ownerID string) bool {
	ret := m.ctrl.Call(m, "IsExisting", ownerID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExisting indicates an expected call of IsExisting
func (mr *MockOwnerManagerMockRecorder) IsExisting(ownerID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExisting", reflect.TypeOf((*MockOwnerManager)(nil).IsExisting), ownerID)
}
