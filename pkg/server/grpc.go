package server

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	userServiceGrpc "github.com/orkungursel/hey-taxi-identity-api/internal/api/grpc"
	userServiceProto "github.com/orkungursel/hey-taxi-identity-api/proto"
)

// startGRPCServer starts the gRPC server
func (s *Server) startGRPCServer(ctx context.Context, cancel context.CancelFunc) {
	s.gs = grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:                  time.Duration(s.config.Server.Grpc.Time) * time.Hour,
		Timeout:               time.Duration(s.config.Server.Grpc.Timeout) * time.Second,
		MaxConnectionIdle:     time.Duration(s.config.Server.Grpc.MaxConnectionIdle) * time.Second,
		MaxConnectionAge:      time.Duration(s.config.Server.Grpc.MaxConnectionAge) * time.Second,
		MaxConnectionAgeGrace: time.Duration(s.config.Server.Grpc.MaxConnectionAgeGrace) * time.Second,
	}))

	userServiceProto.RegisterUserServiceServer(s.gs, userServiceGrpc.NewUserServiceGrpc(s.logger))

	l, err := net.Listen("tcp", s.config.Server.Host+":"+s.config.Server.Grpc.Port)
	defer l.Close()

	if err != nil {
		s.logger.Error(err)
		cancel()
		return
	}

	s.logger.Infof("starting gRPC server on %s", l.Addr())

	if err := s.gs.Serve(l); err != nil {
		s.logger.Error(err)
		cancel()
		return
	}
}

// shutdownGrpcServer stops the gRPC server
func (s *Server) shutdownGrpcServer(ctx context.Context) {
	s.logger.Info("stopping gRPC server")
	s.gs.GracefulStop()
	s.logger.Info("stopped gRPC server")
}
