package grpc

import (
	"context"

	"github.com/rusneustroevkz/http-server/pkg/logger"
	productspb "github.com/rusneustroevkz/http-server/src/product/handlers/grpc/pb"

	"google.golang.org/grpc"
)

type ProductsGRPCServer struct {
	log logger.Logger
	productspb.UnimplementedProductsServer
}

func NewProductsGRPCServer(log logger.Logger) *ProductsGRPCServer {
	productsServer := ProductsGRPCServer{
		log:                         log,
		UnimplementedProductsServer: productspb.UnimplementedProductsServer{},
	}

	return &productsServer
}

func (h *ProductsGRPCServer) Desc() *grpc.ServiceDesc {
	return &productspb.Products_ServiceDesc
}

func (h *ProductsGRPCServer) Service() any {
	return h
}

func (h *ProductsGRPCServer) Product(ctx context.Context, request *productspb.ProductRequest) (*productspb.ProductResponse, error) {
	msg := &productspb.ProductResponse{
		Message: "asd",
	}

	return msg, nil
}
