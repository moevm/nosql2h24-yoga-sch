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
	if req == nil || req.Trainer == nil || req.Trainer.Person == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	gender, err := convertGenGender(req.Trainer.Person.Gender)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	studioID, err := bson.ObjectIDFromHex(req.Trainer.StudioId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bsonID, err := s.Repo.InsertTrainer(ctx, db.Trainer{
		Person: db.Person{
			Name:       req.Trainer.Person.Name,
			Phone:      req.Trainer.Person.Phone,
			PictureURI: req.Trainer.Person.PictureUri,
			BirthDate:  req.Trainer.Person.BirthDate.AsTime(),
			Gender:     gender,
			ClassIDs:   []bson.ObjectID{},
		},
		StudioID: studioID,
	})
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Errorf(codes.InvalidArgument, "studio with id %s not found", req.Trainer.StudioId)
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

	var trainersResp []*gen.Trainer
	for _, t := range trainers {
		trainersResp = append(trainersResp, &gen.Trainer{
			Person:   convertDbPerson(t.Person),
			StudioId: t.StudioID.Hex(),
		})
	}

	return &gen.GetTrainersResponse{Trainers: trainersResp}, nil
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

	return &gen.GetTrainerResponse{
		Trainer: &gen.Trainer{
			Person:   convertDbPerson(trainer.Person),
			StudioId: trainer.StudioID.Hex(),
		},
	}, nil
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