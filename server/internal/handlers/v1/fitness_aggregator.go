package v1

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

type FitnessAggregator struct {
	gen.UnimplementedFitnessAggregatorServer

	Repo db.Repository
}

func (s *FitnessAggregator) CreateClient(
	ctx context.Context, req *gen.CreateClientRequest,
) (*gen.CreateClientResponse, error) {
	if req == nil || req.Client == nil || req.Client.Person == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	gender, err := convertGenGender(req.Client.Person.Gender)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bsonID, err := s.Repo.InsertClient(ctx, db.Client{
		Person: db.Person{
			Name:       req.Client.Person.Name,
			Phone:      req.Client.Person.Phone,
			PictureURI: req.Client.Person.PictureUri,
			BirthDate:  req.Client.Person.BirthDate.AsTime(),
			Gender:     gender,
			Classes:    []bson.ObjectID{},
		},
		Password: req.Client.Password,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.CreateClientResponse{ClientId: bsonID.Hex()}, nil
}

func (s *FitnessAggregator) GetClients(
	ctx context.Context, _ *gen.GetClientsRequest,
) (*gen.GetClientsResponse, error) {
	clients, err := s.Repo.GetClients(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var persons []*gen.Person
	for _, c := range clients {
		persons = append(persons, convertDbPerson(c.Person))
	}

	return &gen.GetClientsResponse{Clients: persons}, nil
}

func (s *FitnessAggregator) GetClient(
	ctx context.Context, req *gen.GetClientRequest,
) (*gen.GetClientResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	bsonID, err := bson.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %v", err)
	}

	client, err := s.Repo.GetClient(ctx, bsonID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("client with id %s not found", req.Id))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.GetClientResponse{Client: convertDbPerson(client.Person)}, nil
}

func (s *FitnessAggregator) DeleteClient(
	ctx context.Context, req *gen.DeleteClientRequest,
) (*gen.DeleteClientResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	bsonID, err := bson.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id: %v", err)
	}

	err = s.Repo.DeleteClient(ctx, bsonID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("client with id %s not found", req.Id))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.DeleteClientResponse{}, nil
}

func convertGenGender(gender gen.Gender) (db.Gender, error) {
	switch gender {
	case gen.Gender_FEMALE:
		return db.GenderFemale, nil
	case gen.Gender_MALE:
		return db.GenderMale, nil
	default:
		return db.GenderFemale, fmt.Errorf("unknown gender: %s", gender)
	}
}

func convertDbGender(gender db.Gender) gen.Gender {
	switch gender {
	case db.GenderMale:
		return gen.Gender_MALE
	default:
		return gen.Gender_FEMALE
	}
}

func convertDbPerson(p db.Person) *gen.Person {
	return &gen.Person{
		Id:         p.ID.Hex(),
		Phone:      p.Phone,
		Name:       p.Name,
		PictureUri: p.PictureURI,
		BirthDate:  timestamppb.New(p.BirthDate),
		Gender:     convertDbGender(p.Gender),
		CreatedAt:  timestamppb.New(p.CreatedAt),
		UpdatedAt:  timestamppb.New(p.UpdatedAt),
	}
}