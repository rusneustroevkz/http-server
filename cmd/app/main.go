package main

import (
	"context"

	"github.com/rusneustroevkz/http-server/internal/config"
	"github.com/rusneustroevkz/http-server/internal/graph/resolvers"
	kafkaClient "github.com/rusneustroevkz/http-server/internal/kafka"
	grpcServer "github.com/rusneustroevkz/http-server/internal/server/grpc"
	httpServer "github.com/rusneustroevkz/http-server/internal/server/http"
	"github.com/rusneustroevkz/http-server/pkg/logger"
	categoriesGRPCHandlers "github.com/rusneustroevkz/http-server/src/categories/handlers/grpc"
	productGraph "github.com/rusneustroevkz/http-server/src/product/handlers/graph"
	productGRPCHandlers "github.com/rusneustroevkz/http-server/src/product/handlers/grpc"
	productKafka "github.com/rusneustroevkz/http-server/src/product/handlers/kafka"
	"github.com/rusneustroevkz/http-server/src/product/handlers/kafka/observers"
	productsRest "github.com/rusneustroevkz/http-server/src/product/handlers/rest"

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
			httpServer.NewPublicServer,
			productsRest.NewProductsRest,
			productGRPCHandlers.NewProductsGRPCServer,
			categoriesGRPCHandlers.NewCategoriesGRPCServer,
			resolvers.NewResolver,
			productGraph.NewProductGraph,
			kafkaClient.NewClient,
			httpServer.NewRouter,
			grpcServer.NewGRPCServer,
			observers.NewCollectProduct,
			productKafka.NewPublisher,
			httpServer.NewMetricsServer,
			httpServer.NewPrivateServer,
		),
		fx.Invoke(
			func(
				lc fx.Lifecycle,
				srv *httpServer.PublicServer,
				router *httpServer.Router,
				productRest *productsRest.ProductsRest,
			) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						routers := router.Mount(productRest)
						srv.SetRoutes(routers)

						return srv.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return srv.Stop(ctx)
					},
				})
			},
			func(
				lc fx.Lifecycle,
				srv *httpServer.MetricsServer,
			) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						srv.SetRoutes()

						return srv.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return srv.Stop(ctx)
					},
				})
			},
			func(
				lc fx.Lifecycle,
				srv *httpServer.PrivateServer,
			) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						srv.SetRoutes()

						return srv.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return srv.Stop(ctx)
					},
				})
			},
			func(
				lc fx.Lifecycle,
				srv *grpcServer.Server,
				productsGRPCServer *productGRPCHandlers.ProductsGRPCServer,
				categoriesGRPCServer *categoriesGRPCHandlers.CategoriesGRPCServer,
			) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						srv.MountServices(productsGRPCServer, categoriesGRPCServer)

						return srv.Start(ctx)
					},
					OnStop: func(ctx context.Context) error {
						return srv.Stop(ctx)
					},
				})
			},
			func(
				lc fx.Lifecycle,
				client *kafkaClient.Client,
				productObserver *observers.CollectProduct,
			) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						return client.Run(ctx, productObserver)
					},
					OnStop: func(ctx context.Context) error {
						return client.Stop(ctx)
					},
				})
			},
		),
	).Run()
}
