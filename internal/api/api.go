package api

import (
	"errors"

	"github.com/orkungursel/hey-taxi-identity-api/internal/api/grpc"
	"github.com/orkungursel/hey-taxi-identity-api/internal/api/http"
	"github.com/orkungursel/hey-taxi-identity-api/internal/infrastructure"
	"github.com/orkungursel/hey-taxi-identity-api/internal/server"
	"go.mongodb.org/mongo-driver/mongo"
)

func Api(s *server.Server, mng *mongo.Client) error {
	if s == nil {
		return errors.New("server is nil")
	}

	c := s.Config()
	if c == nil {
		return errors.New("config is nil")
	}

	logger := s.Logger()

	repo := infrastructure.NewRepository(c, logger, mng)
	tks := infrastructure.NewTokenService(c, logger)
	psw := infrastructure.NewPasswordService(logger)
	svc := infrastructure.NewAuthService(c, logger, repo, tks, psw)
	usvc := infrastructure.NewUserService(c, logger, repo)

	ctrl := http.NewController(c, logger, svc, tks)
	if err := s.RegisterHttpApi("/auth", ctrl); err != nil {
		return err
	}

	g := grpc.NewGrpcUserService(logger, usvc, tks)
	if err := s.RegisterGrpcService(g); err != nil {
		return err
	}

	return nil
}
