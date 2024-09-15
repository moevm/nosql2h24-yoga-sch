package v1

import (
	"context"
	"fmt"

	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

type PlaceRepository struct {
	gen.UnimplementedPlaceRepositoryServer
}

func (r *PlaceRepository) Echo(_ context.Context, req *gen.EchoRequest) (*gen.EchoResponse, error) {
	fmt.Printf("Echo: %s\n", req.Value)
	return &gen.EchoResponse{Value: fmt.Sprintf("Echo: %s", req.Value)}, nil
}
