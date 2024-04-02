package graph

import (
	"context"
	"github.com/rusneustroevkz/http-server/pkg/pointer"

	models_graph "github.com/rusneustroevkz/http-server/internal/graph/models"
)

type ProductGraph struct {
}

func NewProductGraph() *ProductGraph {
	return &ProductGraph{}
}

func (p ProductGraph) Update(ctx context.Context) (*models_graph.ProductResponse, error) {
	return &models_graph.ProductResponse{
		ID: pointer.Of(1),
	}, nil
}

func (p ProductGraph) Get(ctx context.Context) (*models_graph.ProductResponse, error) {
	return &models_graph.ProductResponse{
		ID: pointer.Of(1),
	}, nil
}
