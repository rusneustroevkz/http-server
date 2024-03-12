package http

import (
	"github.com/rusneustroevkz/http-server/graph/generated"
	"github.com/rusneustroevkz/http-server/graph/resolvers"
	"github.com/rusneustroevkz/http-server/internal/config"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

type Graphql struct {
	cfg *config.Config
}

func NewGraphql(cfg *config.Config) *Graphql {
	return &Graphql{cfg: cfg}
}

func (*Graphql) Pattern() string {
	return "/query"
}

func (g *Graphql) Routes() *chi.Mux {
	router := chi.NewRouter()

	if g.cfg.HTTPServer.Test {
		router.Get("/", playground.Handler("GraphQL playground", "/query"))
	}

	schemaConfig := generated.Config{Resolvers: &resolvers.Resolver{}}
	schema := generated.NewExecutableSchema(schemaConfig)
	srv := handler.NewDefaultServer(schema)
	router.Get("/query", srv.ServeHTTP)

	return router
}
