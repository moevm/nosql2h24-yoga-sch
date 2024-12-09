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

func (s *FitnessAggregator) CreateTrainer(
	ctx context.Context, req *gen.CreateTrainerRequest,
) (*gen.CreateTrainerResponse, error) {
	if req == nil || req.Trainer == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	gender, err := convertGenGender(req.Trainer.Gender)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	studioID, err := bson.ObjectIDFromHex(req.Trainer.StudioInfo.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bsonID, err := s.Repo.InsertTrainer(ctx, db.Trainer{
		Person: db.Person{
			Name:       req.Trainer.Name,
			Phone:      req.Trainer.Phone,
			PictureURI: "cdn.example.com",
			BirthDate:  req.Trainer.BirthDate.AsTime(),
			Gender:     gender,

			ClassIDs: []bson.ObjectID{},
		},
		StudioID: studioID,
	})
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Errorf(codes.InvalidArgument, "studio with id %s not found", req.Trainer.StudioInfo.Id)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.CreateTrainerResponse{TrainerId: bsonID.Hex()}, nil
}

func (s *FitnessAggregator) GetTrainers(
	ctx context.Context, _ *gen.GetTrainersRequest,
) (*gen.GetTrainersResponse, error) {
	trainers, err := s.Repo.GetTrainers(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := convertDbTrainers(ctx, trainers, s.Repo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not convert db trainers: %s", err.Error())
	}
	return &gen.GetTrainersResponse{Trainers: result}, nil
}

func (s *FitnessAggregator) GetTrainer(
	ctx context.Context, req *gen.GetTrainerRequest,
) (*gen.GetTrainerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	bsonID, err := bson.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	trainer, err := s.Repo.GetTrainer(ctx, bsonID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	result, err := convertDbTrainer(ctx, trainer, s.Repo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not convert db trainer: %s", err.Error())
	}
	return &gen.GetTrainerResponse{Trainer: result}, nil
}

func (s *FitnessAggregator) DeleteTrainer(
	ctx context.Context, req *gen.DeleteTrainerRequest,
) (*gen.DeleteTrainerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	bsonID, err := bson.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = s.Repo.DeleteTrainer(ctx, bsonID); err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "trainer with id %s not found", req.Id)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.DeleteTrainerResponse{}, nil
}