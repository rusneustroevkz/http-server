package http

import (
	"github.com/rusneustroevkz/http-server/graph"
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

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Get("/query", srv.ServeHTTP)

	return router
}
