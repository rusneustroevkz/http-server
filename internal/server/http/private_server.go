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
	"github.com/go-chi/chi/v5/middleware"
)

type PrivateServer struct {
	cfg *config.Config
	log logger.Logger
	srv *http.Server
}

func NewPrivateServer(
	cfg *config.Config,
	log logger.Logger,
) *PrivateServer {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.PrivateServer.Port),
	}

	return &PrivateServer{
		log: log,
		srv: srv,
	}
}

func (s *PrivateServer) Start(_ context.Context) error {
	listener, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}
	s.log.Info("starting private server", logger.String("addr", s.srv.Addr))
	go func() {
		if err := s.srv.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Fatal("cannot start private server", logger.Error(err), logger.String("port", s.srv.Addr))
		}
	}()
	return nil
}

func (s *PrivateServer) SetRoutes() {
	r := chi.NewRouter()
	r.Mount("/debug", middleware.Profiler())
	s.srv.Handler = r
}

func (s *PrivateServer) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
