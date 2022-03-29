package server

import "google.golang.org/grpc"

type GrpcService interface {
	Register(registrar grpc.ServiceRegistrar)
}

func (s *Server) RegisterGrpcService(gs GrpcService) error {
	s.grpcServices = append(s.grpcServices, gs)

	return nil
}

func (s *Server) mapServices() {
	for _, gs := range s.grpcServices {
		gs.Register(s.gs)
	}
}
