package resolvers

import (
	productGrap "github.com/rusneustroevkz/http-server/src/product/handlers/graph"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	product product
}

func NewResolver(product *productGrap.ProductGraph) *Resolver {
	return &Resolver{
		product: product,
	}
}
