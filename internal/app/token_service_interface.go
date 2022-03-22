//go:generate mockgen -source token_service_interface.go -destination mock/token_service_mock.go -package mock
package app

import (
	"context"
	"net/http"

	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
)

type TokenService interface {
	GenerateAccessToken(ctx context.Context, user *model.User) (string, error)
	GenerateRefreshToken(ctx context.Context, user *model.User) (string, error)
	ValidateAccessToken(ctx context.Context, token string) error
	ValidateRefreshToken(ctx context.Context, token string) error
	ExtractFromRequest(ctx context.Context, r *http.Request) (map[string]interface{}, error)
}
