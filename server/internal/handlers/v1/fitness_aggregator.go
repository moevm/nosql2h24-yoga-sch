package v1

import (
	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

type FitnessAggregator struct {
	gen.UnimplementedFitnessAggregatorServer

	Repo db.Repository
}