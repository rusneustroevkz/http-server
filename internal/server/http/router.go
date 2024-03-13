package http

import (
	_ "github.com/rusneustroevkz/http-server/docs"
	"github.com/rusneustroevkz/http-server/internal/config"
	"github.com/rusneustroevkz/http-server/internal/graph/generated"
	"github.com/rusneustroevkz/http-server/internal/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(cfg *config.Config, resolver *resolvers.Resolver, routes ...Route) *chi.Mux {
	router := chi.NewRouter()

	for _, route := range routes {
		router.Mount(route.Pattern(), route.Routes())
	}

	if !cfg.App.Production {
		router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))
		router.Handle("/graph/playground", playground.Handler("GraphQL playground", "/graph/query"))
	}

	schemaConfig := generated.Config{Resolvers: resolver}
	schema := generated.NewExecutableSchema(schemaConfig)
	srv := handler.New(schema)
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})
	router.Handle("/graph/query", srv)

	return router
}

type Route interface {
	Routes() *chi.Mux

	Pattern() string
}
