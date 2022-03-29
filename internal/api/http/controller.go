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
	config       *config.Config
	logger       logger.ILogger
	authService  app.AuthService
	tokenService app.TokenService
}

func NewController(config *config.Config, logger logger.ILogger, s app.AuthService, ts app.TokenService) *Controller {
	return &Controller{
		authService:  s,
		tokenService: ts,
		logger:       logger,
		config:       config,
	}
}

// RegisterRoutes registers the routes to the echo server
func (a *Controller) RegisterRoutes(e *echo.Group) {
	e.Use(middleware.ErrorHandler())

	e.POST("/login/", a.login())
	e.POST("/register/", a.register())
	e.POST("/refresh-token/", a.refreshToken())
	e.GET("/me/", a.me(), middleware.Auth(a.tokenService))
}

// @Summary      Login
// @Description  User Login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        payload  body      app.LoginRequest  true  "Payload"
// @Success      200      {array}   app.SuccessAuthResponse
// @Failure      400  {object}  app.HTTPError
// @Failure      500  {object}  app.HTTPError
// @Router       /auth/login [post]
func (a *Controller) login() echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := &app.LoginRequest{}
		if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
			return err
		}

		if err := app.Validate(payload); err != nil {
			return err
		}

		println(c.Request().Header.Get(echo.HeaderXRequestID))

		res, err := a.authService.Login(c.Request().Context(), payload)
		if err != nil {
			return err
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
// @Failure      400      {object}  app.HTTPError
// @Failure      500      {object}  app.HTTPError
// @Router       /auth/register [post]
func (a *Controller) register() echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := &app.RegisterRequest{}
		if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
			return err
		}

		if err := app.Validate(payload); err != nil {
			return err
		}

		res, err := a.authService.Register(c.Request().Context(), payload)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}

// @Summary      Refreshes all tokens
// @Description  Fetch the details of logged-in user by access token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        payload  body      app.RefreshTokenRequest  true  "Payload"
// @Success      200  {array}   app.UserResponse
// @Failure      400      {object}  app.HTTPError
// @Failure      500      {object}  app.HTTPError
// @Router       /auth/refresh-token [post]
// @Security     BearerAuth
func (a *Controller) refreshToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := &app.RefreshTokenRequest{}
		if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
			return err
		}

		if err := app.Validate(payload); err != nil {
			return err
		}

		res, err := a.authService.RefreshToken(c.Request().Context(), payload)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}

// @Summary      User Details
// @Description  Fetch the details of logged-in user by access token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200      {array}   app.UserResponse
// @Failure      400      {object}  app.HTTPError
// @Failure      401      {object}  app.HTTPError
// @Failure      500      {object}  app.HTTPError
// @Router       /auth/me [get]
// @Security     BearerAuth
func (a *Controller) me() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := GetUserId(c)
		if err != nil {
			return err
		}

		res, err := a.authService.Me(c.Request().Context(), userId)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	}
}
