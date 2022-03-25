package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/internal/api/http/middleware"
	"github.com/orkungursel/hey-taxi-identity-api/internal/app"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
)

type Controller struct {
	config *config.Config
	logger logger.ILogger
	svc    app.Service
	ts     app.TokenService
}

func NewController(config *config.Config, logger logger.ILogger, s app.Service, ts app.TokenService) *Controller {
	return &Controller{
		svc:    s,
		ts:     ts,
		logger: logger,
		config: config,
	}
}

// RegisterRoutes registers the routes to the echo server
func (a *Controller) RegisterRoutes(e *echo.Group) {
	e.POST("/login/", a.login())
	e.POST("/register/", a.register())
	e.GET("/me/", a.me(), middleware.Auth(a.ts))
}

// @Summary      Login
// @Description  User Login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        payload  body      app.LoginRequest  true  "Payload"
// @Success      200      {array}   app.SuccessAuthResponse
// @Failure      400  {object}  echo.HTTPError
// @Failure      500  {object}  echo.HTTPError
// @Router       /auth/login [post]
func (a *Controller) login() echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := &app.LoginRequest{}
		if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := app.Validate(payload); err != nil {
			a.logger.Debugf("invalid login request: %s", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		res, err := a.svc.Login(c.Request().Context(), payload)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}

// @Summary      Register
// @Description  User Registration
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        payload  body      app.RegisterRequest  true  "Payload"
// @Success      200      {array}   app.SuccessAuthResponse
// @Failure      400      {object}  echo.HTTPError
// @Failure      500      {object}  echo.HTTPError
// @Router       /auth/register [post]
func (a *Controller) register() echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := &app.RegisterRequest{}
		if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := app.Validate(payload); err != nil {
			a.logger.Debugf("invalid login request: %s", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		res, err := a.svc.Register(c.Request().Context(), payload)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}

// @Summary      User Details
// @Description  Fetch the details of logged-in user by access token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200  {array}   app.UserResponse
// @Failure      400      {object}  echo.HTTPError
// @Failure      500      {object}  echo.HTTPError
// @Router       /auth/me [get]
// @Security     BearerAuth
func (a *Controller) me() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := GetUserId(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		res, err := a.svc.Me(c.Request().Context(), userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, res)
	}
}
