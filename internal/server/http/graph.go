package http

import (
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/rusneustroevkz/http-server/internal/config"
	"github.com/rusneustroevkz/http-server/internal/graph/generated"
	"github.com/rusneustroevkz/http-server/internal/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

type Graphql struct {
	cfg      *config.Config
	resolver *resolvers.Resolver
}

func NewGraphql(cfg *config.Config, resolver *resolvers.Resolver) *Graphql {
	return &Graphql{cfg: cfg, resolver: resolver}
}

func (*Graphql) Pattern() string {
	return "/"
}

func (g *Graphql) Routes() *chi.Mux {
	router := chi.NewRouter()

	if g.cfg.HTTPServer.Test {
		router.Handle("/graph/playground", playground.Handler("GraphQL playground", "/graph/query"))
	}

	schemaConfig := generated.Config{Resolvers: g.resolver}
	schema := generated.NewExecutableSchema(schemaConfig)
	srv := handler.New(schema)
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})
	router.Handle("/graph/query", srv)

	return router
}
