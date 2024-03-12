package main

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/rusneustroevkz/http-server/internal/config"
	productGRPCHandlers "github.com/rusneustroevkz/http-server/internal/product/handlers/grpc"
	productHTTPHandlers "github.com/rusneustroevkz/http-server/internal/product/handlers/http"
	grpcServer "github.com/rusneustroevkz/http-server/internal/server/grpc"
	httpServer "github.com/rusneustroevkz/http-server/internal/server/http"
	"github.com/rusneustroevkz/http-server/pkg/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Store server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	fx.New(
		fx.WithLogger(func(log logger.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log.Logger()}
		}),
		fx.Provide(
			config.NewConfig,
			logger.NewLogger,
			httpServer.NewHTTPServer,
			productHTTPHandlers.NewProductsHTTPHandler,
			func(productHTTPHandler *productHTTPHandlers.ProductHTTPHandler) *chi.Mux {
				return httpServer.MountRoutes(productHTTPHandler)
			},
			productGRPCHandlers.NewProductsGRPCServer,
			func(
				cfg *config.Config,
				log logger.Logger,
				productsGRPCServer *productGRPCHandlers.ProductsGRPCServer,
			) *grpcServer.Server {
				return grpcServer.NewGRPCServer(
					cfg,
					log,
					productsGRPCServer,
				)
			},
		),
		fx.Invoke(
			func(lc fx.Lifecycle, srv *httpServer.Server) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return srv.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return srv.Stop(ctx)
					},
				})
			},
			func(lc fx.Lifecycle, srv *grpcServer.Server) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						srv.MountServices()

						return srv.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return srv.Stop(ctx)
					},
				})
			},
		),
	).Run()
}
