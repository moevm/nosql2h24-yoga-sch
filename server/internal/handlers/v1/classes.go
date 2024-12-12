package v1

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

func (s *FitnessAggregator) CreateClass(
	ctx context.Context, req *gen.CreateClassRequest,
) (*gen.CreateClassResponse, error) {
	if req == nil || req.Class == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	trainerID, err := bson.ObjectIDFromHex(req.Class.TrainerId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	studioID, err := bson.ObjectIDFromHex(req.Class.StudioId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bsonID, err := s.Repo.InsertClass(ctx, db.Class{
		Name: req.Class.Name,
		Time: req.Class.Time.AsTime(),

		TrainerID: trainerID,
		StudioID:  studioID,
		ClientIDs: []bson.ObjectID{},
	})
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Errorf(codes.InvalidArgument, "trainer or studio with id %s not found", req.Class.TrainerId)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.CreateClassResponse{ClassId: bsonID.Hex()}, nil
}

func (s *FitnessAggregator) GetClasses(
	ctx context.Context, _ *gen.GetClassesRequest,
) (*gen.GetClassesResponse, error) {
	classes, err := s.Repo.GetClasses(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := convertDbClasses(ctx, classes, s.Repo)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to convert db classes: %v", err))
	}
	return &gen.GetClassesResponse{Classes: result}, nil
}

func (s *FitnessAggregator) GetClass(
	ctx context.Context, req *gen.GetClassRequest,
) (*gen.GetClassResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	bsonID, err := bson.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %s", req.Id)
	}

	class, err := s.Repo.GetClass(ctx, bsonID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "class with id %s not found", req.Id)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := convertDbClass(ctx, class, s.Repo)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to convert db class: %v", err))
	}
	return &gen.GetClassResponse{Class: result}, nil
}

func (s *FitnessAggregator) DeleteClass(
	ctx context.Context, req *gen.DeleteClassRequest,
) (*gen.DeleteClassResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	bsonID, err := bson.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %v", err)
	}

	if err = s.Repo.DeleteClass(ctx, bsonID); err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "class with id %s not found", req.Id)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.DeleteClassResponse{}, nil
}