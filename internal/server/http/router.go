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

type Router struct {
	cfg      *config.Config
	resolver *resolvers.Resolver
}

func NewRouter(cfg *config.Config, resolver *resolvers.Resolver) *Router {
	return &Router{
		cfg:      cfg,
		resolver: resolver,
	}
}

func (r *Router) GetRouters(routes ...Route) *chi.Mux {
	mux := chi.NewRouter()

	for _, route := range routes {
		mux.Mount(route.Pattern(), route.Routes())
	}

	if !r.cfg.App.Production {
		mux.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("doc.json")))
		mux.Handle("/graph/playground", playground.Handler("GraphQL playground", "/graph/query"))
	}

	schemaConfig := generated.Config{Resolvers: r.resolver}
	schema := generated.NewExecutableSchema(schemaConfig)
	srv := handler.New(schema)
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})
	mux.Handle("/graph/query", srv)

	return mux
}

type Route interface {
	Routes() *chi.Mux

	Pattern() string
}
