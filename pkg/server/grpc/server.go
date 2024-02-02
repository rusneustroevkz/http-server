package grpc

import (
	"github.com/rusneustroevkz/http-server/pkg/logger"
)

type Server struct {
	log logger.Logger
}

func NewGRPCServer() *Server {
	return &Server{}
}
