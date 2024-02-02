package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/rusneustroevkz/http-server/internal/config"
	"github.com/rusneustroevkz/http-server/pkg/logger"
	"google.golang.org/grpc"
)

type Server struct {
	log    logger.Logger
	cfg    *config.Config
	server *grpc.Server
}

func NewGRPCServer(cfg *config.Config, log logger.Logger) *Server {
	server := grpc.NewServer()

	return &Server{
		cfg:    cfg,
		log:    log,
		server: server,
	}
}

func (s *Server) Start(_ context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GRPCServer.Port))
	if err != nil {
		s.log.Fatal("failed to listen GRPC", logger.Error(err))
	}
	s.log.Info("starting GRPC server", logger.Int64("addr", s.cfg.GRPCServer.Port))
	go func() {
		if err := s.server.Serve(lis); err != nil {
			s.log.Fatal("failed to serve GRPC", logger.Error(err), logger.Int64("port", s.cfg.GRPCServer.Port))
		}
	}()
	return nil
}

func (s *Server) Stop(_ context.Context) error {
	s.server.GracefulStop()
	return nil
}
