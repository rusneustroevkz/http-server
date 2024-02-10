package grpc

import (
	"context"

	petspb "github.com/rusneustroevkz/http-server/internal/pets/handlers/grpc/pb"
	"github.com/rusneustroevkz/http-server/pkg/logger"
	"google.golang.org/grpc"
)

type PetsGRPCServer struct {
	log logger.Logger
	petspb.UnimplementedPetsServer
}

func NewPetsGRPCServer(log logger.Logger) *PetsGRPCServer {
	petsServer := PetsGRPCServer{
		log:                     log,
		UnimplementedPetsServer: petspb.UnimplementedPetsServer{},
	}

	return &petsServer
}

func (h *PetsGRPCServer) Desc() *grpc.ServiceDesc {
	return &petspb.Pets_ServiceDesc
}

func (h *PetsGRPCServer) Service() any {
	return h
}

func (h *PetsGRPCServer) SayHello(ctx context.Context, request *petspb.HelloRequest) (*petspb.HelloReply, error) {
	msg := &petspb.HelloReply{
		Message: "ok",
	}

	return msg, nil
}
