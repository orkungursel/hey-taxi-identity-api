//go:generate mockgen -source password_service.go -destination mock/password_service_mock.go -package mock
package app

import "context"

type PasswordService interface {
	Hash(ctx context.Context, password string) (string, error)
	Compare(ctx context.Context, hashedPassword string, password string) error
}
