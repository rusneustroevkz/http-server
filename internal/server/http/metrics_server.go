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
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsServer struct {
	cfg *config.Config
	log logger.Logger
	srv *http.Server
}

func NewMetricsServer(
	cfg *config.Config,
	log logger.Logger,
) *MetricsServer {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.MetricsServer.Port),
	}

	return &MetricsServer{
		log: log,
		srv: srv,
	}
}

func (s *MetricsServer) Start(_ context.Context) error {
	listener, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}
	s.log.Info("starting metrics server", logger.String("addr", s.srv.Addr))
	go func() {
		if err := s.srv.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Fatal("cannot start metrics server", logger.Error(err), logger.String("port", s.srv.Addr))
		}
	}()
	return nil
}

func (s *MetricsServer) SetRoutes() {
	mr := chi.NewRouter()
	mr.Handle("/metrics", promhttp.Handler())
	s.srv.Handler = mr
}

func (s *MetricsServer) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
