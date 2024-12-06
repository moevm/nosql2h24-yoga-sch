package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
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
		return nil, status.Errorf(codes.Internal, "Internal error: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (p *AdminPanel) ImportDB(
	ctx context.Context, req *gen.DBData,
) (*emptypb.Empty, error) {
	slog.Info(req.Data)

	var data map[string][]bson.M
	if err := json.Unmarshal([]byte(req.Data), &data); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid JSON data: %s", err))
	}

	if err := p.Repo.ImportDB(ctx, data); err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func (p *AdminPanel) ExportDB(
	ctx context.Context, _ *emptypb.Empty,
) (*gen.DBData, error) {
	data, err := p.Repo.ExportDB(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %w", err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %w", err)
	}

	return &gen.DBData{Data: string(jsonData)}, nil
}