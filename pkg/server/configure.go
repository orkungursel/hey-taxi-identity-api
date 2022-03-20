package server

import (
	emw "github.com/labstack/echo/v4/middleware"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/server/middleware"
)

// configure the echo server
func (s *Server) configure() {
	s.echo.HidePort = true
	s.echo.HideBanner = true

	// add pre middlewares
	s.echo.Pre(middleware.AddTrailingSlash())
	s.echo.Pre(middleware.Logger(s.logger))

	// add middlewares
	s.echo.Use(emw.Recover())
	s.echo.Use(middleware.CORS(s.config))
	s.echo.Use(emw.Secure())
}
