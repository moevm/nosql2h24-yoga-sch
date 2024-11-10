package v1

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

func (s *FitnessAggregator) CreateStudio(
	ctx context.Context, req *gen.CreateStudioRequest,
) (*gen.CreateStudioResponse, error) {
	if req == nil || req.Studio == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	bsonID, err := s.Repo.InsertStudio(ctx, db.Studio{
		Name:    req.Studio.Name,
		Address: req.Studio.Address,

		ClassIDs:   []bson.ObjectID{},
		TrainerIDs: []bson.ObjectID{},
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.CreateStudioResponse{StudioId: bsonID.Hex()}, nil
}

func (s *FitnessAggregator) GetStudios(
	ctx context.Context, _ *gen.GetStudiosRequest,
) (*gen.GetStudiosResponse, error) {
	studios, err := s.Repo.GetStudios(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var studiosResp []*gen.Studio
	for _, s := range studios {
		studiosResp = append(studiosResp, convertDbStudio(s))
	}

	return &gen.GetStudiosResponse{Studios: studiosResp}, nil
}

func (s *FitnessAggregator) GetStudio(
	ctx context.Context, req *gen.GetStudioRequest,
) (*gen.GetStudioResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	bsonID, err := bson.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %s", req.Id)
	}

	studio, err := s.Repo.GetStudio(ctx, bsonID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "studio with id %s not found", req.Id)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.GetStudioResponse{Studio: convertDbStudio(studio)}, nil
}

func (s *FitnessAggregator) DeleteStudio(
	ctx context.Context, req *gen.DeleteStudioRequest,
) (*gen.DeleteStudioResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	bsonID, err := bson.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %v", err)
	}

	if err = s.Repo.DeleteStudio(ctx, bsonID); err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "studio with id %s not found", req.Id)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.DeleteStudioResponse{}, nil
}