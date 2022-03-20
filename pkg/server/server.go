package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
	logger logger.ILogger
	ctx    context.Context
	apis   []ApiHandlerItem
	done   chan struct{}
}

func New(ctx context.Context, config *config.Config, logger logger.ILogger) *Server {
	return &Server{
		ctx:    ctx,
		echo:   echo.New(),
		config: config,
		logger: logger,
		done:   make(chan struct{}),
	}
}

// Run starts the server
func (s *Server) Run() error {
	ctx, cancel := signal.NotifyContext(s.ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.configure()
	s.mapHandlers()

	go s.start(ctx, cancel)
	go s.waitForCancellationToken(ctx)

	<-s.done

	s.logger.Info("Shutting down...")

	ctx, sdCancel := context.WithTimeout(ctx, time.Duration(s.config.Server.ShutdownTimeout)*time.Second)
	defer sdCancel()

	if err := s.shutdown(ctx); err != nil {
		s.logger.Errorf("Error while shutting down server: %s", err)
		return err
	}

	<-ctx.Done()

	s.logger.Info("Server Exited Properly")

	return nil
}

func (s *Server) Logger() logger.ILogger {
	return s.logger
}

// start starts the server
func (s *Server) start(ctx context.Context, cancel context.CancelFunc) {
	// create http server
	httpServer := &http.Server{
		Addr: s.config.Server.Host + ":" + s.config.Server.Port,
	}

	s.logger.Infof("Starting server on %s", httpServer.Addr)

	// start the server
	if err := s.echo.StartServer(httpServer); err != nil && err != http.ErrServerClosed {
		s.logger.Error(err)
		cancel()
		return
	}
}

// shutdown stops the server
func (s *Server) shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

// waitForCancellationToken waits for the cancellation token
func (s *Server) waitForCancellationToken(ctx context.Context) {
	defer close(s.done)
	<-ctx.Done()
	s.done <- struct{}{}
}
