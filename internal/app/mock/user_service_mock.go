// Code generated by MockGen. DO NOT EDIT.
// Source: user_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	app "github.com/orkungursel/hey-taxi-identity-api/internal/app"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// UsersByIds mocks base method.
func (m *MockUserService) UsersByIds(ctx context.Context, ids []string) ([]*app.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersByIds", ctx, ids)
	ret0, _ := ret[0].([]*app.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersByIds indicates an expected call of UsersByIds.
func (mr *MockUserServiceMockRecorder) UsersByIds(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersByIds", reflect.TypeOf((*MockUserService)(nil).UsersByIds), ctx, ids)
}