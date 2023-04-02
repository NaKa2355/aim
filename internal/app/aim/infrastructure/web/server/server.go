package server

import (
	"net"

	v1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	s *grpc.Server
}

func New(handler v1.AimServiceServer) *Server {
	s := grpc.NewServer()
	reflection.Register(s)
	v1.RegisterAimServiceServer(s, handler)
	return &Server{
		s: s,
	}
}

func (s *Server) Start(listener net.Listener) {
	go s.s.Serve(listener)
}

func (s *Server) Stop() {
	s.s.GracefulStop()
}