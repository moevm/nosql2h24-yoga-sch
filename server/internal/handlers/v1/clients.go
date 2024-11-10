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

func (s *FitnessAggregator) CreateClient(
	ctx context.Context, req *gen.CreateClientRequest,
) (*gen.CreateClientResponse, error) {
	if req == nil || req.Client == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	gender, err := convertGenGender(req.Client.Gender)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bsonID, err := s.Repo.InsertClient(ctx, db.Client{
		Person: db.Person{
			Name:       req.Client.Name,
			Phone:      req.Client.Phone,
			PictureURI: "cdn.example.com",
			BirthDate:  req.Client.BirthDate.AsTime(),
			Gender:     gender,

			ClassIDs: []bson.ObjectID{},
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

	tokenClientID := GetUserIDFromContext(ctx)
	if tokenClientID != bsonID.Hex() {
		return nil, status.Error(codes.PermissionDenied, "client ID does not match token")
	}

	client, err := s.Repo.GetClient(ctx, bsonID)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "client with id %s not found", req.Id)
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

	if err = s.Repo.DeleteClient(ctx, bsonID); err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("client with id %s not found", req.Id))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.DeleteClientResponse{}, nil
}