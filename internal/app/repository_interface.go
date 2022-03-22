//go:generate mockgen -source repository_interface.go -destination mock/repository_mock.go -package mock
package app

import (
	"context"

	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, id string, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
}
