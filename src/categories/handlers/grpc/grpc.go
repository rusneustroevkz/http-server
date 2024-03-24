package grpc

import (
	"context"

	categoriesspb "github.com/rusneustroevkz/http-server/src/categories/handlers/grpc/pb"
	"google.golang.org/grpc"
)

type CategoriesGRPCServer struct {
	*categoriesspb.UnimplementedCategoriesServer
}

func NewCategoriesGRPCServer() *CategoriesGRPCServer {
	return &CategoriesGRPCServer{
		&categoriesspb.UnimplementedCategoriesServer{},
	}
}

func (g *CategoriesGRPCServer) Desc() *grpc.ServiceDesc {
	return &categoriesspb.Categories_ServiceDesc
}

func (g *CategoriesGRPCServer) Service() any {
	return g
}

func (g *CategoriesGRPCServer) Category(context.Context, *categoriesspb.CategoryRequest) (*categoriesspb.CategoryResponse, error) {
	return &categoriesspb.CategoryResponse{
		Message: "ok",
	}, nil
}
