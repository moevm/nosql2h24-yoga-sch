package v1

import (
	"context"

	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

type LibraryServiceServer struct {
	gen.UnimplementedLibraryServiceServer
}

func (s *LibraryServiceServer) Echo(context.Context, *gen.EchoRequest) (*gen.EchoResponse, error) {
	return &gen.EchoResponse{Value: "Echo success"}, nil
}
