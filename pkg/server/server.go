package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/orkungursel/hey-taxi-identity-api/config"
	userServiceGrpc "github.com/orkungursel/hey-taxi-identity-api/internal/api/grpc"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/logger"
	userService "github.com/orkungursel/hey-taxi-identity-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
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

	go s.waitForCancellationToken(ctx)

	<-s.done

	s.logger.Info("Shutting down...")

	ctx, sdCancel := context.WithTimeout(ctx, time.Duration(s.config.Server.ShutdownTimeout)*time.Second)
	defer sdCancel()

	if err := s.shutdownHttpServer(ctx); err != nil {
		s.logger.Errorf("Error while shutting down server: %s", err)
	}

	<-ctx.Done()
	s.logger.Info("Server Exited Properly")

	return nil
}

func (s *Server) Logger() logger.ILogger {
	return s.logger
}

// startHttpServer starts the server
func (s *Server) startHttpServer(ctx context.Context, cancel context.CancelFunc) {
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

// shutdownHttpServer stops the server
func (s *Server) shutdownHttpServer(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

// startGRPCServer starts the gRPC server
func (s *Server) startGRPCServer(ctx context.Context, cancel context.CancelFunc) {
	s.gs = grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:                  time.Duration(s.config.Server.Grpc.Time) * time.Hour,
		Timeout:               time.Duration(s.config.Server.Grpc.Timeout) * time.Second,
		MaxConnectionIdle:     time.Duration(s.config.Server.Grpc.MaxConnectionIdle) * time.Second,
		MaxConnectionAge:      time.Duration(s.config.Server.Grpc.MaxConnectionAge) * time.Second,
		MaxConnectionAgeGrace: time.Duration(s.config.Server.Grpc.MaxConnectionAgeGrace) * time.Second,
	}))

	userService.RegisterUserServiceServer(s.gs, userServiceGrpc.NewUserServiceGrpc(s.logger))

	l, err := net.Listen("tcp", s.config.Server.Host+":"+s.config.Server.Grpc.Port)
	defer l.Close()

	if err != nil {
		s.logger.Error(err)
		cancel()
		return
	}

	s.logger.Infof("Starting gRPC server on %s", l.Addr())

	if err := s.gs.Serve(l); err != nil {
		s.logger.Error(err)
		cancel()
		return
	}
}

func (s *Server) shutdownGrpcServer(ctx context.Context) {
	s.gs.GracefulStop()
}

// waitForCancellationToken waits for the cancellation token
func (s *Server) waitForCancellationToken(ctx context.Context) {
	defer close(s.done)
	<-ctx.Done()
	s.done <- struct{}{}
}
