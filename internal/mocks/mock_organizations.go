// Code generated by MockGen. DO NOT EDIT.
// Source: internal/store/organizations.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrganizationLister is a mock of OrganizationLister interface
type MockOrganizationLister struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationListerMockRecorder
}

// MockOrganizationListerMockRecorder is the mock recorder for MockOrganizationLister
type MockOrganizationListerMockRecorder struct {
	mock *MockOrganizationLister
}

// NewMockOrganizationLister creates a new mock instance
func NewMockOrganizationLister(ctrl *gomock.Controller) *MockOrganizationLister {
	mock := &MockOrganizationLister{ctrl: ctrl}
	mock.recorder = &MockOrganizationListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationLister) EXPECT() *MockOrganizationListerMockRecorder {
	return m.recorder
}

// GetAllOrganizations mocks base method
func (m *MockOrganizationLister) GetAllOrganizations() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrganizations")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrganizations indicates an expected call of GetAllOrganizations
func (mr *MockOrganizationListerMockRecorder) GetAllOrganizations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrganizations", reflect.TypeOf((*MockOrganizationLister)(nil).GetAllOrganizations))
}

// MockOrganizationCreator is a mock of OrganizationCreator interface
type MockOrganizationCreator struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationCreatorMockRecorder
}

// MockOrganizationCreatorMockRecorder is the mock recorder for MockOrganizationCreator
type MockOrganizationCreatorMockRecorder struct {
	mock *MockOrganizationCreator
}

// NewMockOrganizationCreator creates a new mock instance
func NewMockOrganizationCreator(ctrl *gomock.Controller) *MockOrganizationCreator {
	mock := &MockOrganizationCreator{ctrl: ctrl}
	mock.recorder = &MockOrganizationCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationCreator) EXPECT() *MockOrganizationCreatorMockRecorder {
	return m.recorder
}

// CreateOrganization mocks base method
func (m *MockOrganizationCreator) CreateOrganization(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganization", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganization indicates an expected call of CreateOrganization
func (mr *MockOrganizationCreatorMockRecorder) CreateOrganization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganization", reflect.TypeOf((*MockOrganizationCreator)(nil).CreateOrganization), arg0)
}

// MockOrganizationDeleter is a mock of OrganizationDeleter interface
type MockOrganizationDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationDeleterMockRecorder
}

// MockOrganizationDeleterMockRecorder is the mock recorder for MockOrganizationDeleter
type MockOrganizationDeleterMockRecorder struct {
	mock *MockOrganizationDeleter
}

// NewMockOrganizationDeleter creates a new mock instance
func NewMockOrganizationDeleter(ctrl *gomock.Controller) *MockOrganizationDeleter {
	mock := &MockOrganizationDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationDeleter) EXPECT() *MockOrganizationDeleterMockRecorder {
	return m.recorder
}

// DeleteOrganization mocks base method
func (m *MockOrganizationDeleter) DeleteOrganization(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrganization", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrganization indicates an expected call of DeleteOrganization
func (mr *MockOrganizationDeleterMockRecorder) DeleteOrganization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganization", reflect.TypeOf((*MockOrganizationDeleter)(nil).DeleteOrganization), arg0)
}

// MockOrganizationStore is a mock of OrganizationStore interface
type MockOrganizationStore struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationStoreMockRecorder
}

// MockOrganizationStoreMockRecorder is the mock recorder for MockOrganizationStore
type MockOrganizationStoreMockRecorder struct {
	mock *MockOrganizationStore
}

// NewMockOrganizationStore creates a new mock instance
func NewMockOrganizationStore(ctrl *gomock.Controller) *MockOrganizationStore {
	mock := &MockOrganizationStore{ctrl: ctrl}
	mock.recorder = &MockOrganizationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationStore) EXPECT() *MockOrganizationStoreMockRecorder {
	return m.recorder
}

// GetAllOrganizations mocks base method
func (m *MockOrganizationStore) GetAllOrganizations() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrganizations")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrganizations indicates an expected call of GetAllOrganizations
func (mr *MockOrganizationStoreMockRecorder) GetAllOrganizations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrganizations", reflect.TypeOf((*MockOrganizationStore)(nil).GetAllOrganizations))
}

// CreateOrganization mocks base method
func (m *MockOrganizationStore) CreateOrganization(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganization", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganization indicates an expected call of CreateOrganization
func (mr *MockOrganizationStoreMockRecorder) CreateOrganization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganization", reflect.TypeOf((*MockOrganizationStore)(nil).CreateOrganization), arg0)
}

// DeleteOrganization mocks base method
func (m *MockOrganizationStore) DeleteOrganization(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrganization", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrganization indicates an expected call of DeleteOrganization
func (mr *MockOrganizationStoreMockRecorder) DeleteOrganization(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganization", reflect.TypeOf((*MockOrganizationStore)(nil).DeleteOrganization), arg0)
}
