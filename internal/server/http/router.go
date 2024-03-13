package http

import (
	_ "github.com/rusneustroevkz/http-server/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(routes ...Route) *chi.Mux {
	router := chi.NewRouter()

	for _, route := range routes {
		router.Mount(route.Pattern(), route.Routes())
	}

	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))

	return router
}

type Route interface {
	Routes() *chi.Mux

	Pattern() string
}
