// Code generated by MockGen. DO NOT EDIT.
// Source: token_service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	app "github.com/orkungursel/hey-taxi-identity-api/internal/app"
	model "github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
)

// MockTokenService is a mock of TokenService interface.
type MockTokenService struct {
	ctrl     *gomock.Controller
	recorder *MockTokenServiceMockRecorder
}

// MockTokenServiceMockRecorder is the mock recorder for MockTokenService.
type MockTokenServiceMockRecorder struct {
	mock *MockTokenService
}

// NewMockTokenService creates a new mock instance.
func NewMockTokenService(ctrl *gomock.Controller) *MockTokenService {
	mock := &MockTokenService{ctrl: ctrl}
	mock.recorder = &MockTokenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenService) EXPECT() *MockTokenServiceMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockTokenService) GenerateAccessToken(ctx context.Context, user *model.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", ctx, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockTokenServiceMockRecorder) GenerateAccessToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockTokenService)(nil).GenerateAccessToken), ctx, user)
}

// GenerateRefreshToken mocks base method.
func (m *MockTokenService) GenerateRefreshToken(ctx context.Context, user *model.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", ctx, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken.
func (mr *MockTokenServiceMockRecorder) GenerateRefreshToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockTokenService)(nil).GenerateRefreshToken), ctx, user)
}

// ParseToken mocks base method.
func (m *MockTokenService) ParseToken(ctx context.Context, token string) (app.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", ctx, token)
	ret0, _ := ret[0].(app.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockTokenServiceMockRecorder) ParseToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockTokenService)(nil).ParseToken), ctx, token)
}

// ValidateAccessTokenFromRequest mocks base method.
func (m *MockTokenService) ValidateAccessTokenFromRequest(ctx context.Context, r *http.Request) (app.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAccessTokenFromRequest", ctx, r)
	ret0, _ := ret[0].(app.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateAccessTokenFromRequest indicates an expected call of ValidateAccessTokenFromRequest.
func (mr *MockTokenServiceMockRecorder) ValidateAccessTokenFromRequest(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAccessTokenFromRequest", reflect.TypeOf((*MockTokenService)(nil).ValidateAccessTokenFromRequest), ctx, r)
}

// ValidateRefreshToken mocks base method.
func (m *MockTokenService) ValidateRefreshToken(ctx context.Context, token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRefreshToken", ctx, token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateRefreshToken indicates an expected call of ValidateRefreshToken.
func (mr *MockTokenServiceMockRecorder) ValidateRefreshToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRefreshToken", reflect.TypeOf((*MockTokenService)(nil).ValidateRefreshToken), ctx, token)
}
