package grpc

import (
	"context"
	"strings"

	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	. "github.com/orkungursel/hey-taxi-identity-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type GrpcUserService struct {
	logger logger.ILogger
	svc    app.UserService
	tks    app.TokenService
	UnimplementedUserServiceServer
}

func NewGrpcUserService(logger logger.ILogger, svc app.UserService, tks app.TokenService) *GrpcUserService {
	return &GrpcUserService{
		logger: logger,
		svc:    svc,
		tks:    tks,
	}
}

func (s *GrpcUserService) Register(reg grpc.ServiceRegistrar) {
	RegisterUserServiceServer(reg, s)
}

func (s *GrpcUserService) GetUserInfo(ctx context.Context, r *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	users, err := s.svc.UsersByIds(ctx, r.UserIds)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "failed to get users: %s", err)
	}

	usersData := make([]*UserInfo, 0)
	for _, user := range users {
		usersData = append(usersData, &UserInfo{
			Id:     user.Id,
			Email:  user.Email,
			Name:   strings.Join([]string{user.FirstName, user.LastName}, " "),
			Role:   user.Role,
			Avatar: user.Avatar,
		})
	}

	return &GetUserInfoResponse{
		Users: usersData,
	}, nil
}
