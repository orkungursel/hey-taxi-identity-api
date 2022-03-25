package infrastructure

import (
	"context"
	"time"

	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/internal/domain/model"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	app.Service
	config *config.Config
	logger logger.ILogger
	repo   app.Repository
	ts     app.TokenService
	pws    IPasswordService
}

func NewService(config *config.Config, logger logger.ILogger, repo app.Repository, ts app.TokenService, pws IPasswordService) *Service {
	return &Service{
		config: config,
		logger: logger,
		repo:   repo,
		ts:     ts,
		pws:    pws,
	}
}

// Login is used to authenticate user
func (s *Service) Login(ctx context.Context, r *app.LoginRequest) (*app.SuccessAuthResponse, error) {
	if err := app.Validate(r); err != nil {
		s.logger.Debugf("invalid login request: %s", err)
		return nil, err
	}

	user, err := s.repo.GetUserByEmail(ctx, r.Email)
	if err != nil {
		return nil, err
	}

	if err := s.pws.Compare(ctx, user.Password, r.Password); err != nil {
		s.logger.Debugf("invalid password: %s", err)
		return nil, errors.New("invalid password")
	}

	accessToken, err := s.ts.GenerateAccessToken(ctx, user)
	if err != nil {
		s.logger.Warnf("failed to generate access token: %s", err)
		return nil, err
	}

	refreshToken, err := s.ts.GenerateRefreshToken(ctx, user)
	if err != nil {
		s.logger.Warnf("failed to generate refresh token: %s", err)
		return nil, err
	}

	return &app.SuccessAuthResponse{
		UserDto:               *app.UserResponseFromUser(user),
		AccessToken:           accessToken,
		AccessTokenExpiresIn:  s.config.Jwt.AccessTokenExp,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresIn: s.config.Jwt.RefreshTokenExp,
	}, nil
}

// Register is used to register new user
func (s *Service) Register(ctx context.Context, r *app.RegisterRequest) (*app.SuccessAuthResponse, error) {
	if err := app.Validate(r); err != nil {
		s.logger.Warnf("invalid register request: %s", err)
		return nil, err
	}

	if hasUser, _ := s.repo.GetUserByEmail(ctx, r.Email); hasUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := s.pws.Hash(ctx, r.Password)
	if err != nil {
		s.logger.Warnf("failed to hash password: %s", err)
		return nil, err
	}

	user := &model.User{
		Email:     r.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      model.RoleUser,
	}

	uid, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return nil, err
	}

	user.Id = objectId

	accessToken, err := s.ts.GenerateAccessToken(ctx, user)
	if err != nil {
		s.logger.Warnf("failed to generate access token: %s", err)
		return nil, err
	}

	refreshToken, err := s.ts.GenerateRefreshToken(ctx, user)
	if err != nil {
		s.logger.Warnf("failed to generate refresh token: %s", err)
		return nil, err
	}

	return &app.SuccessAuthResponse{
		UserDto:               *app.UserResponseFromUser(user),
		AccessToken:           accessToken,
		AccessTokenExpiresIn:  s.config.Jwt.AccessTokenExp,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresIn: s.config.Jwt.RefreshTokenExp,
	}, nil
}

func (s *Service) Me(ctx context.Context, uid string) (*app.UserResponse, error) {
	user, err := s.repo.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return app.UserResponseFromUser(user), nil
}
