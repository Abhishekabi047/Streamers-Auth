package api

import (
	"fmt"
	"net"
	"service/pkg/config"
	"service/pkg/pb/auth"

	"google.golang.org/grpc"
)

type Server struct {
	gs   *grpc.Server
	Lis  net.Listener
	Port string
}

func NewGrpcServe(c *config.Config, service auth.AuthServiceServer) (*Server, error) {
	grpcserver := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcserver, service)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		return nil, err
	}
	return &Server{
		gs:   grpcserver,
		Lis:  lis,
		Port: c.Port,
	}, nil
}

func (s *Server) Start() error {
	fmt.Println("Service start on", s.Port)
	return s.gs.Serve(s.Lis)
}
