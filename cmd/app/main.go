package main

import (
	"context"

	categoriesGRPCHandlers "github.com/rusneustroevkz/http-server/internal/categories/handlers/grpc"
	"github.com/rusneustroevkz/http-server/internal/config"
	"github.com/rusneustroevkz/http-server/internal/graph/resolvers"
	productGraph "github.com/rusneustroevkz/http-server/internal/product/handlers/graph"
	productGRPCHandlers "github.com/rusneustroevkz/http-server/internal/product/handlers/grpc"
	productsRest "github.com/rusneustroevkz/http-server/internal/product/handlers/rest"
	grpcServer "github.com/rusneustroevkz/http-server/internal/server/grpc"
	httpServer "github.com/rusneustroevkz/http-server/internal/server/http"
	"github.com/rusneustroevkz/http-server/pkg/logger"

	"github.com/go-chi/chi/v5"
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
			productsRest.NewProductsRest,
			httpServer.NewGraphql,
			productGRPCHandlers.NewProductsGRPCServer,
			categoriesGRPCHandlers.NewCategoriesGRPCServer,
			resolvers.NewResolver,
			productGraph.NewProductGraph,
			func(productRest *productsRest.ProductsRest, graphRoutes *httpServer.Graphql) *chi.Mux {
				return httpServer.Routes(productRest, graphRoutes)
			},
			func(
				cfg *config.Config,
				log logger.Logger,
				productsGRPCServer *productGRPCHandlers.ProductsGRPCServer,
				categoriesGRPCServer *categoriesGRPCHandlers.CategoriesGRPCServer,
			) *grpcServer.Server {
				return grpcServer.NewGRPCServer(
					cfg,
					log,
					productsGRPCServer,
					categoriesGRPCServer,
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
