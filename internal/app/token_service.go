//go:generate mockgen -source token_service.go -destination mock/token_service_mock.go -package mock
package app

import (
	"context"
	"net/http"

	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
)

type TokenService interface {
	GenerateAccessToken(ctx context.Context, user *model.User) (string, error)
	GenerateRefreshToken(ctx context.Context, user *model.User) (string, error)
	ParseToken(ctx context.Context, token string) (Claims, error)
	ValidateAccessTokenFromRequest(ctx context.Context, r *http.Request) (Claims, error)
	ValidateRefreshToken(ctx context.Context, token string) (string, error)
}
