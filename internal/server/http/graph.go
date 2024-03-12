package http

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Graphql struct {
	schema graphql.Schema
}

func NewGraphql() (*Graphql, error) {
	cfg := graphql.SchemaConfig{
		Query: &graphql.Object{
			PrivateName:        "",
			PrivateDescription: "",
			IsTypeOf:           nil,
		},
		Mutation: &graphql.Object{
			PrivateName:        "",
			PrivateDescription: "",
			IsTypeOf:           nil,
		},
		Subscription: &graphql.Object{
			PrivateName:        "",
			PrivateDescription: "",
			IsTypeOf:           nil,
		},
		Types: []graphql.Type{},
		Directives: []*graphql.Directive{
			{
				Name:        "",
				Description: "",
				Locations:   nil,
				Args:        nil,
			},
		},
		Extensions: []graphql.Extension{},
	}

	schema, err := graphql.NewSchema(cfg)
	if err != nil {
		return nil, err
	}

	return &Graphql{schema: schema}, nil
}

func (g *Graphql) Handler() *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   &g.schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

func (g *Graphql) Playground() *handler.Handler {
	return handler.New(&handler.Config{
		Schema:     &g.schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
}
