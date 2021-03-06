package middleware

import (
	"strings"

	emw "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func AddTrailingSlash() echo.MiddlewareFunc {
	return emw.AddTrailingSlashWithConfig(
		emw.TrailingSlashConfig{
			Skipper: func(c echo.Context) bool {
				if strings.HasPrefix(c.Request().URL.Path, "/swagger/") {
					return true
				}
				return false
			},
		},
	)
}
