//go:generate mockgen -source repository.go -destination mock/repository_mock.go -package mock
package app

import (
	"context"

	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	GetUsersByIds(ctx context.Context, ids []string) ([]*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (string, error)
	UpdateUser(ctx context.Context, id string, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
}
