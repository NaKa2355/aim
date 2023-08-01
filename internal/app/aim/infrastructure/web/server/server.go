package server

import (
	"net"
	"os"
	"os/signal"

	v1 "github.com/NaKa2355/irdeck-proto/gen/go/aim/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	s *grpc.Server
}

func New(handler v1.AimServiceServer, useReflection bool) *Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryErrInterceptor),
	)
	if useReflection {
		reflection.Register(s)
	}
	v1.RegisterAimServiceServer(s, handler)
	return &Server{
		s: s,
	}
}

func (s *Server) Start(listener net.Listener) {
	go func() {
		defer listener.Close()
		s.s.Serve(listener)
	}()
}

func (s *Server) WaitSigAndStop(sig ...os.Signal) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, sig...)
	<-sigCh
	s.s.GracefulStop()
}

func (s *Server) Stop() {
	s.s.GracefulStop()
}
