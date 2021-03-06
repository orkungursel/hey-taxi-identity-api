//go:build dev

package swagger

import (
	"github.com/labstack/echo/v4"
	_ "github.com/orkungursel/hey-taxi-identity-api/docs" // docs is generated by Swag CLI
	"github.com/orkungursel/hey-taxi-identity-api/internal/server"
	es "github.com/swaggo/echo-swagger"
)

type SwaggerApi struct {
	*server.Server
}

func (s *SwaggerApi) RegisterRoutes(group *echo.Group) {
	group.GET("", es.WrapHandler)
}
