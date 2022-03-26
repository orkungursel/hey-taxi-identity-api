package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/orkungursel/hey-taxi-identity-api/config"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	"google.golang.org/grpc"
)

type Server struct {
	echo   *echo.Echo
	gs     *grpc.Server
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

	go s.startHttpServer(ctx, cancel)
	go s.startGRPCServer(ctx, cancel)
	go s.waitSignal(ctx)

	<-s.done

	s.logger.Info("shutting down...")

	s.shutdownHttpServer(ctx)
	s.shutdownGrpcServer(ctx)

	return nil
}

func (s *Server) Config() *config.Config {
	return s.config
}

func (s *Server) Logger() logger.ILogger {
	return s.logger
}

// waitSignal waits for the cancellation token
func (s *Server) waitSignal(ctx context.Context) {
	defer close(s.done)
	<-ctx.Done()
	s.done <- struct{}{}
}
