package adder

import (
	"context"
	"turbo/pkg/api"
)

type GRPCServer struct{ api.UnimplementedAdderServer }

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
