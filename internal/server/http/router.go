package http

import (
	_ "github.com/rusneustroevkz/http-server/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(routes ...Route) *chi.Mux {
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
