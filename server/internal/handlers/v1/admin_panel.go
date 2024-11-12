package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

type AdminPanel struct {
	gen.UnimplementedAdminPanelServer

	Repo db.Repository
}

func (p *AdminPanel) DropDB(
	ctx context.Context, _ *emptypb.Empty,
) (*emptypb.Empty, error) {
	if err := p.Repo.DropDB(ctx); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}
	return &emptypb.Empty{}, nil
}