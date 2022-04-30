//go:generate mockgen -source auth_service.go -destination mock/auth_service_mock.go -package mock

package app

import (
	"context"
)

type AuthService interface {
	Login(ctx context.Context, r *LoginRequest) (*LoginResponse, error)
	Register(ctx context.Context, r *RegisterRequest) (*LoginResponse, error)
	RefreshToken(ctx context.Context, r *RefreshTokenRequest) (*RefreshTokenResponse, error)
	Me(ctx context.Context, uid string) (*UserResponse, error)
}
