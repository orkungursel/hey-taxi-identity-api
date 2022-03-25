package grpc

import (
	"context"

	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	userService "github.com/orkungursel/hey-taxi-identity-api/proto"
)

type UserServiceGrpc struct {
	logger logger.ILogger
	userService.UnimplementedUserServiceServer
}

func NewUserServiceGrpc(logger logger.ILogger) *UserServiceGrpc {
	return &UserServiceGrpc{
		logger: logger,
	}
}

func (s *UserServiceGrpc) GetUserInfo(ctx context.Context, r *userService.GetUserInfoRequest) (*userService.GetUserInfoResponse, error) {
	s.logger.Debugf("Token: %s", r.Token)

	out := &userService.GetUserInfoResponse{
		UserId:    "123",
		Email:     "orquncom@gmail.com",
		FirstName: "Orkun",
		Type:      "driver",
		Role:      "admin",
		Avatar:    "https://avatars0.githubusercontent.com/u/17098981?s=460&v=4",
	}
	return out, nil
}
