package v1

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

func (s *FitnessAggregator) CreateAppointment(
	ctx context.Context, req *gen.CreateAppointmentRequest,
) (*gen.CreateAppointmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	clientBsonID, err := bson.ObjectIDFromHex(req.ClientId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid client ID")
	}

	tokenClientID := GetUserIDFromContext(ctx)
	if tokenClientID != clientBsonID.Hex() {
		return nil, status.Error(codes.PermissionDenied, "client ID does not match token")
	}

	classBsonID, err := bson.ObjectIDFromHex(req.ClassId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid class ID")
	}

	if err := s.Repo.MakeAppointment(ctx, classBsonID, clientBsonID); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.CreateAppointmentResponse{}, nil
}

func (s *FitnessAggregator) CancelAppointment(
	ctx context.Context, req *gen.CancelAppointmentRequest,
) (*gen.CancelAppointmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "no request provided")
	}

	clientBsonID, err := bson.ObjectIDFromHex(req.ClientId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid client ID")
	}

	tokenClientID := GetUserIDFromContext(ctx)
	if tokenClientID != clientBsonID.Hex() {
		return nil, status.Error(codes.PermissionDenied, "client ID does not match token")
	}

	classBsonID, err := bson.ObjectIDFromHex(req.ClassId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid class ID")
	}

	if err := s.Repo.CancelAppointment(ctx, classBsonID, clientBsonID); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.CancelAppointmentResponse{}, nil
}