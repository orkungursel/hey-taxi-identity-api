package infrastructure

import (
	"context"
	"time"

	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	"github.com/pkg/errors"
)

type Service struct {
	app.Service
	config *config.Config
	logger logger.ILogger
	repo   app.Repository
}

func NewService(config *config.Config, logger logger.ILogger, repo app.Repository) *Service {
	return &Service{
		config: config,
		logger: logger,
		repo:   repo,
	}
}

func (s *Service) Login(ctx context.Context, r *app.LoginRequest) (*app.LoginResponse, error) {
	if err := app.Validate(r); err != nil {
		s.logger.Warnf("invalid login request: %s", err)
		return nil, err
	}

	user, err := s.repo.GetUserByEmail(ctx, r.Email)
	if err != nil {
		return nil, err
	}

	if user.Password != r.Password {
		return nil, errors.New("invalid password")
	}

	return &app.LoginResponse{
		Token:        "token",
		RefreshToken: "refresh_token",
		ExpiresIn:    s.config.Auth.AccessTokenExp,
	}, nil
}

func (s *Service) Register(ctx context.Context, r *app.RegisterRequest) (*app.LoginResponse, error) {
	if err := app.Validate(r); err != nil {
		s.logger.Warnf("invalid register request: %s", err)
		return nil, err
	}

	user := &model.User{
		Email:     r.Email,
		Password:  r.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      &model.RoleUser,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return &app.LoginResponse{
		Token:        "token",
		RefreshToken: "refresh_token",
		ExpiresIn:    s.config.Auth.AccessTokenExp,
	}, nil
}
