//go:generate mockgen -source password_service.go -destination mock/password_service_mock.go -package mock
package infrastructure

import (
	"context"
	"errors"
	"strings"

	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type IPasswordService interface {
	Hash(ctx context.Context, password string) (string, error)
	Compare(ctx context.Context, hashedPassword string, password string) error
}

type PasswordService struct {
	logger logger.ILogger
}

func NewPasswordService(logger logger.ILogger) *PasswordService {
	return &PasswordService{
		logger: logger,
	}
}

func (s *PasswordService) Hash(ctx context.Context, password string) (string, error) {
	sanitizedPassword, err := sanitize(password)
	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(sanitizedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *PasswordService) Compare(ctx context.Context, hashedPassword string, password string) error {
	sanitizedPassword, err := sanitize(password)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(sanitizedPassword)); err != nil {
		return err
	}
	return nil
}

func sanitize(password string) (string, error) {
	sanitizedPassword := strings.TrimSpace(password)
	if sanitizedPassword == "" {
		return "", errors.New("password is empty")
	}

	return sanitizedPassword, nil
}
