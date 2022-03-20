//go:build !swagger
// +build !swagger

package swagger

import (
	"github.com/labstack/echo/v4"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/server"
)

type SwaggerApi struct {
	*server.Server
}

func Api(s *server.Server) *SwaggerApi {
	api := &SwaggerApi{
		Server: s,
	}

	return api
}

func (s *SwaggerApi) RegisterRoutes(group *echo.Group) {}
