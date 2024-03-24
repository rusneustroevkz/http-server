package main

import (
	"context"

	"github.com/rusneustroevkz/http-server/internal/config"
	petsGRPCHandlers "github.com/rusneustroevkz/http-server/internal/pets/handlers/grpc"
	petsHTTPHandlers "github.com/rusneustroevkz/http-server/internal/pets/handlers/http"
	grpcServer "github.com/rusneustroevkz/http-server/internal/server/grpc"
	httpServer "github.com/rusneustroevkz/http-server/internal/server/http"
	"github.com/rusneustroevkz/http-server/pkg/logger"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server PetStore server.
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
			petsHTTPHandlers.NewPetsHTTPHandler,
			httpServer.MountRoutes,
			petsGRPCHandlers.NewPetsGRPCServer,
			grpcServer.NewGRPCServer,
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
			func(lc fx.Lifecycle, srv *grpcServer.Server, petsGRPCServer *petsGRPCHandlers.PetsGRPCServer) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						srv.MountServices(petsGRPCServer)

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
