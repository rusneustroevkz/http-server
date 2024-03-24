package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/rusneustroevkz/http-server/internal/config"
	"github.com/rusneustroevkz/http-server/pkg/logger"

	"github.com/go-chi/chi/v5"
)

type PublicServer struct {
	cfg *config.Config
	log logger.Logger
	srv *http.Server
}

func NewPublicServer(
	cfg *config.Config,
	log logger.Logger,
) *PublicServer {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.PublicServer.Port),
	}

	return &PublicServer{
		log: log,
		srv: srv,
	}
}

func (s *PublicServer) Start(_ context.Context) error {
	listener, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}
	s.log.Info("starting public server", logger.String("addr", s.srv.Addr))
	go func() {
		if err := s.srv.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Fatal("cannot start public server", logger.Error(err), logger.String("port", s.srv.Addr))
		}
	}()
	return nil
}

func (s *PublicServer) SetRoutes(mux *chi.Mux) {
	s.srv.Handler = mux
}

func (s *PublicServer) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
