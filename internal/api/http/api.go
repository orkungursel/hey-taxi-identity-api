package http

import (
	"errors"

	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/infrastructure"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/server"
	"go.mongodb.org/mongo-driver/mongo"
)

func Api(c *config.Config, s *server.Server, mng *mongo.Client) error {
	if c == nil {
		return errors.New("config is nil")
	}

	if s == nil {
		return errors.New("server is nil")
	}

	logger := s.Logger()

	repo := infrastructure.NewRepository(c, logger, mng)
	ts := infrastructure.NewTokenService(c, logger)
	psw := infrastructure.NewPasswordService(logger)

	svc := infrastructure.NewService(c, logger, repo, ts, psw)
	controller := NewController(c, logger, svc, ts)

	if err := s.RegisterHttpApi("/auth", controller); err != nil {
		return err
	}

	return nil
}
