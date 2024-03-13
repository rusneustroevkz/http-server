package graph

import (
	"context"

	models_graph "github.com/rusneustroevkz/http-server/internal/graph/models"
	"github.com/rusneustroevkz/http-server/utils/pointer"
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
