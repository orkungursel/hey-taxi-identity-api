package infrastructure

import (
	"context"

	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
)

type UserService struct {
	app.UserService
	config *config.Config
	logger logger.ILogger
	repo   app.Repository
}

func NewUserService(config *config.Config, logger logger.ILogger, repo app.Repository) *UserService {
	return &UserService{
		config: config,
		logger: logger,
		repo:   repo,
	}
}

func (s *UserService) UsersByIds(ctx context.Context, ids []string) ([]*app.UserResponse, error) {
	result, err := s.repo.GetUsersByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	var users []*app.UserResponse
	for _, user := range result {
		users = append(users, app.UserResponseFromUser(user))
	}

	return users, nil
}
