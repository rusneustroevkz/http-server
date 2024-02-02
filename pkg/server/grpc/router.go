package grpc

import (
	"go.uber.org/fx"
	grpcServer "google.golang.org/grpc"
)

func MountRoutes(server *grpcServer.Server, routes []Route) *grpcServer.Server {
	for _, route := range routes {
		server.RegisterService(route.ServiceDesc(), route.SS())
	}

	return server
}

type Route interface {
	ServiceDesc() *grpcServer.ServiceDesc
	SS() any
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
