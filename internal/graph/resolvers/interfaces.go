package resolvers

import (
	"context"

	models_graph "github.com/rusneustroevkz/http-server/internal/graph/models"
)

type product interface {
	Update(ctx context.Context) (*models_graph.ProductResponse, error)
	Get(ctx context.Context) (*models_graph.ProductResponse, error)
}
