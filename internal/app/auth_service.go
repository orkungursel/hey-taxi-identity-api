//go:generate mockgen -source auth_service.go -destination mock/auth_service_mock.go -package mock

package app

import (
	"context"
)

type AuthService interface {
	Login(ctx context.Context, r *LoginRequest) (*SuccessAuthResponse, error)
	Register(ctx context.Context, r *RegisterRequest) (*SuccessAuthResponse, error)
	RefreshToken(ctx context.Context, r *RefreshTokenRequest) (*SuccessAuthResponse, error)
	Me(ctx context.Context, uid string) (*UserResponse, error)
}
