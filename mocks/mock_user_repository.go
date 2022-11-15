// Code generated by MockGen. DO NOT EDIT.
// Source: app/repositories/user_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "go_base_project/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepositoriesInterface is a mock of UserRepositoriesInterface interface.
type MockUserRepositoriesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoriesInterfaceMockRecorder
}

// MockUserRepositoriesInterfaceMockRecorder is the mock recorder for MockUserRepositoriesInterface.
type MockUserRepositoriesInterfaceMockRecorder struct {
	mock *MockUserRepositoriesInterface
}

// NewMockUserRepositoriesInterface creates a new mock instance.
func NewMockUserRepositoriesInterface(ctrl *gomock.Controller) *MockUserRepositoriesInterface {
	mock := &MockUserRepositoriesInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoriesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoriesInterface) EXPECT() *MockUserRepositoriesInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepositoriesInterface) Create(name any) models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name)
	ret0, _ := ret[0].(models.User)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoriesInterfaceMockRecorder) Create(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepositoriesInterface)(nil).Create), name)
}

// Find mocks base method.
func (m *MockUserRepositoriesInterface) Find(Id any) models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", Id)
	ret0, _ := ret[0].(models.User)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockUserRepositoriesInterfaceMockRecorder) Find(Id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserRepositoriesInterface)(nil).Find), Id)
}
