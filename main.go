package main

import (
	"context"
	"github.com/rusneustroevkz/http-server/internal/config"
	petsHTTPHandlers "github.com/rusneustroevkz/http-server/internal/pets/handlers/http"
	"github.com/rusneustroevkz/http-server/pkg/logger"
	grpcServer "github.com/rusneustroevkz/http-server/pkg/server/grpc"
	httpServer "github.com/rusneustroevkz/http-server/pkg/server/http"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
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
			httpServer.NewHTTPServer,
			httpServer.AsRoute(petsHTTPHandlers.NewPetsHTTPHandler),
			fx.Annotate(
				httpServer.MountRoutes,
				fx.ParamTags(`group:"routes"`),
			),
			grpcServer.NewGRPCServer,
			//grpcServer.AsRoute(petsGRPCHandlers.NewPetsGRPCHandler),
			//fx.Annotate(
			//	grpcServer.MountRoutes,
			//	fx.ParamTags(`group:"routes"`),
			//),
			logger.NewLogger,
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
