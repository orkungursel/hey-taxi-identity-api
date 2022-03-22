//go:generate mockgen -source service_interface.go -destination mock/service_mock.go -package mock

package app

import (
	"context"
)

type Service interface {
	Login(ctx context.Context, r *LoginRequest) (*LoginResponse, error)
	Register(ctx context.Context, r *RegisterRequest) (*LoginResponse, error)
}
