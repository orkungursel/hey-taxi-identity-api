// Code generated by MockGen. DO NOT EDIT.
// Source: auth_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	app "github.com/orkungursel/hey-taxi-identity-api/internal/app"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthService) Login(ctx context.Context, r *app.LoginRequest) (*app.SuccessAuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, r)
	ret0, _ := ret[0].(*app.SuccessAuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceMockRecorder) Login(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), ctx, r)
}

// Me mocks base method.
func (m *MockAuthService) Me(ctx context.Context, uid string) (*app.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Me", ctx, uid)
	ret0, _ := ret[0].(*app.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Me indicates an expected call of Me.
func (mr *MockAuthServiceMockRecorder) Me(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Me", reflect.TypeOf((*MockAuthService)(nil).Me), ctx, uid)
}

// RefreshToken mocks base method.
func (m *MockAuthService) RefreshToken(ctx context.Context, r *app.RefreshTokenRequest) (*app.SuccessAuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", ctx, r)
	ret0, _ := ret[0].(*app.SuccessAuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockAuthServiceMockRecorder) RefreshToken(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockAuthService)(nil).RefreshToken), ctx, r)
}

// Register mocks base method.
func (m *MockAuthService) Register(ctx context.Context, r *app.RegisterRequest) (*app.SuccessAuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, r)
	ret0, _ := ret[0].(*app.SuccessAuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthServiceMockRecorder) Register(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthService)(nil).Register), ctx, r)
}
