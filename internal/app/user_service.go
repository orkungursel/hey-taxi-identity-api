//go:generate mockgen -source user_service.go -destination mock/user_service_mock.go -package mock
package app

import "context"

type UserService interface {
	UsersByIds(ctx context.Context, ids []string) ([]*UserResponse, error)
}
