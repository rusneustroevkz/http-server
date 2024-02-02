package http

import (
	_ "asd/docs"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"
)

func NewServeMux(routes []Route) *chi.Mux {
	mux := chi.NewMux()

	for _, route := range routes {
		mux.Mount(route.Pattern(), route.Routes())
	}

	mux.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))

	return mux
}

type Route interface {
	Routes() *chi.Mux

	Pattern() string
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}