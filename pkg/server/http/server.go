package http

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rusneustroevkz/http-server/internal/config"
	"github.com/rusneustroevkz/http-server/pkg/logger"
)

type Server struct {
	cfg *config.Config
	log logger.Logger
	srv *http.Server
}

func NewHTTPServer(cfg *config.Config, mux *chi.Mux, log logger.Logger) *Server {
	srv := &http.Server{Addr: fmt.Sprintf(":%d", cfg.HTTPServer.Port), Handler: mux}

	return &Server{
		log: log,
		srv: srv,
	}
}

func (s *Server) Start(_ context.Context) error {
	listener, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}
	s.log.Info("starting HTTP server", logger.String("addr", s.srv.Addr))
	go func() {
		if err := s.srv.Serve(listener); err != nil {
			s.log.Fatal("cannot HTTP start server", logger.Error(err), logger.String("port", s.srv.Addr))
		}
	}()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}