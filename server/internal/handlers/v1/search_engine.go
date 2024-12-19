package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

type SearchEngine struct {
	gen.UnimplementedSearchEngineServer

	Repo db.Repository
}

func (e *SearchEngine) SearchClients(
	ctx context.Context, req *gen.ClientsFilter,
) (*gen.SearchClientsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}

	genders, err := convertGenGenders(req.Genders)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid genders: %w", err)
	}

	birthDateInterval := convertGenTimeInterval(req.BirthDateIntervalBegin, req.BirthDateIntervalEnd)
	createdAtInterval := convertGenTimeInterval(req.CreatedAtIntervalBegin, req.CreatedAtIntervalEnd)
	updatedAtInterval := convertGenTimeInterval(req.UpdatedAtIntervalBegin, req.UpdatedAtIntervalEnd)

	persons, err := e.Repo.SearchClients(ctx, db.ClientsFilter{
		IDSubstring:         req.IdSubstring,
		NameSubstring:       req.NameSubstring,
		PhoneSubstring:      req.PhoneSubstring,
		PictureURISubstring: req.PictureUriSubstring,
		BirthDateInterval:   birthDateInterval,
		Genders:             genders,
		CreatedAtInterval:   createdAtInterval,
		UpdatedAtInterval:   updatedAtInterval,
		ClassNameSubstrings: req.ClassNameSubstrings,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "searching clients error: %w", err)
	}

	result, err := convertDbPersons(ctx, persons, e.Repo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "converting db persons error: %w", err)
	}
	return &gen.SearchClientsResponse{Clients: result}, nil
}

func (e *SearchEngine) SearchTrainers(
	ctx context.Context, req *gen.TrainersFilter,
) (*gen.SearchTrainersResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}

	genders, err := convertGenGenders(req.Genders)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid genders: %w", err)
	}

	birthDateInterval := convertGenTimeInterval(req.BirthDateIntervalBegin, req.BirthDateIntervalEnd)
	createdAtInterval := convertGenTimeInterval(req.CreatedAtIntervalBegin, req.CreatedAtIntervalEnd)
	updatedAtInterval := convertGenTimeInterval(req.UpdatedAtIntervalBegin, req.UpdatedAtIntervalEnd)

	trainers, err := e.Repo.SearchTrainers(ctx, db.TrainersFilter{
		IDSubstring:             req.IdSubstring,
		NameSubstring:           req.NameSubstring,
		PhoneSubstring:          req.PhoneSubstring,
		PictureURISubstring:     req.PictureUriSubstring,
		BirthDateInterval:       birthDateInterval,
		Genders:                 genders,
		CreatedAtInterval:       createdAtInterval,
		UpdatedAtInterval:       updatedAtInterval,
		ClassNameSubstrings:     req.ClassNameSubstrings,
		StudioAddressSubstrings: req.StudioAddressSubstrings,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "searching trainers error: %w", err)
	}

	result, err := convertDbTrainers(ctx, trainers, e.Repo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "converting db trainers error: %w", err)
	}
	return &gen.SearchTrainersResponse{Trainers: result}, nil
}

func (e *SearchEngine) SearchStudios(
	ctx context.Context, req *gen.StudiosFilter,
) (*gen.SearchStudiosResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}

	createdAtInterval := convertGenTimeInterval(req.CreatedAtIntervalBegin, req.CreatedAtIntervalEnd)
	updatedAtInterval := convertGenTimeInterval(req.UpdatedAtIntervalBegin, req.UpdatedAtIntervalEnd)

	studios, err := e.Repo.SearchStudios(ctx, db.StudiosFilter{
		IDSubstring:           req.IdSubstring,
		AddressSubstring:      req.AddressSubstring,
		CreatedAtInterval:     createdAtInterval,
		UpdatedAtInterval:     updatedAtInterval,
		ClassNameSubstrings:   req.ClassNameSubstrings,
		TrainerNameSubstrings: req.TrainerNameSubstrings,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "searching studios error: %w", err)
	}

	result, err := convertDbStudios(ctx, studios, e.Repo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "converting db studios error: %w", err)
	}
	return &gen.SearchStudiosResponse{Studios: result}, nil
}

func (e *SearchEngine) SearchClasses(
	ctx context.Context, req *gen.ClassesFilter,
) (*gen.SearchClassesResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nil request")
	}

	timeInterval := convertGenTimeInterval(req.TimeIntervalBegin, req.TimeIntervalEnd)
	createdAtInterval := convertGenTimeInterval(req.CreatedAtIntervalBegin, req.CreatedAtIntervalEnd)
	updatedAtInterval := convertGenTimeInterval(req.UpdatedAtIntervalBegin, req.UpdatedAtIntervalEnd)

	classes, err := e.Repo.SearchClasses(ctx, db.ClassesFilter{
		IDSubstring:             req.IdSubstring,
		NameSubstring:           req.NameSubstring,
		TimeInterval:            timeInterval,
		CreatedAtInterval:       createdAtInterval,
		UpdatedAtInterval:       updatedAtInterval,
		OnlyAvailable:           req.OnlyAvailable,
		StudioAddressSubstrings: req.StudioAddressSubstrings,
		TrainerNameSubstrings:   req.TrainerNameSubstrings,
		ClientNameSubstrings:    req.ClientNameSubstrings,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "searching classes error: %w", err)
	}

	result, err := convertDbClasses(ctx, classes, e.Repo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "converting db classes error: %w", err)
	}
	return &gen.SearchClassesResponse{Classes: result}, nil
}