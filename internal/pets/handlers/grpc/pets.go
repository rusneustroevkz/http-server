package grpc

import (
	petspb "github.com/rusneustroevkz/http-server/internal/pets/handlers/grpc/pb"
	"github.com/rusneustroevkz/http-server/pkg/logger"
	grpcServer "google.golang.org/grpc"
)

type PetsGRPCHandler struct {
	log logger.Logger
}

func NewPetsGRPCHandler(log logger.Logger) *PetsGRPCHandler {
	return &PetsGRPCHandler{
		log: log,
	}
}

func (h *PetsGRPCHandler) ServiceDesc() *grpcServer.ServiceDesc {
	return &petspb.Pets_ServiceDesc
}

func (h *PetsGRPCHandler) SS() any {
	return ""
}
