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

	return &gen.SearchClientsResponse{Clients: convertDbPersons(persons)}, nil
}