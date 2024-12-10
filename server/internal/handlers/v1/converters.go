package v1

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

func ptr[T any](val T) *T {
	return &val
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

func convertGenGenders(genders []gen.Gender) (res []db.Gender, err error) {
	for _, g := range genders {
		r, err := convertGenGender(g)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func convertGenTimeInterval(
	begin *timestamppb.Timestamp, end *timestamppb.Timestamp,
) (res db.TimeInterval) {
	if begin != nil {
		res.Begin = ptr(bson.NewDateTimeFromTime(begin.AsTime()))
	}
	if end != nil {
		res.End = ptr(bson.NewDateTimeFromTime(end.AsTime()))
	}
	return res
}

func convertDbGender(gender db.Gender) gen.Gender {
	switch gender {
	case db.GenderMale:
		return gen.Gender_MALE
	default:
		return gen.Gender_FEMALE
	}
}

func collectClassesInfo(
	ctx context.Context, ids []bson.ObjectID, repo db.Repository,
) ([]*gen.NameIDPair, error) {
	var classesInfo []*gen.NameIDPair
	for _, id := range ids {
		class, err := repo.GetClass(ctx, id)
		if err != nil {
			return nil, err
		}
		classesInfo = append(classesInfo, &gen.NameIDPair{
			Name: class.Name,
			Id:   id.Hex(),
		})
	}
	return classesInfo, nil
}

func collectStudiosInfo(
	ctx context.Context, ids []bson.ObjectID, repo db.Repository,
) ([]*gen.NameIDPair, error) {
	var studiosInfo []*gen.NameIDPair
	for _, id := range ids {
		studio, err := repo.GetStudio(ctx, id)
		if err != nil {
			return nil, err
		}
		studiosInfo = append(studiosInfo, &gen.NameIDPair{
			Name: studio.Address,
			Id:   id.Hex(),
		})
	}
	return studiosInfo, nil
}

func collectTrainersInfo(
	ctx context.Context, ids []bson.ObjectID, repo db.Repository,
) ([]*gen.NameIDPair, error) {
	var trainersInfo []*gen.NameIDPair
	for _, id := range ids {
		trainer, err := repo.GetTrainer(ctx, id)
		if err != nil {
			return nil, err
		}
		trainersInfo = append(trainersInfo, &gen.NameIDPair{
			Name: trainer.Name,
			Id:   id.Hex(),
		})
	}
	return trainersInfo, nil
}

func collectClientsInfo(
	ctx context.Context, ids []bson.ObjectID, repo db.Repository,
) ([]*gen.NameIDPair, error) {
	var clientsInfo []*gen.NameIDPair
	for _, id := range ids {
		client, err := repo.GetClient(ctx, id)
		if err != nil {
			return nil, err
		}
		clientsInfo = append(clientsInfo, &gen.NameIDPair{
			Name: client.Name,
			Id:   id.Hex(),
		})
	}
	return clientsInfo, nil
}

func convertDbPerson(
	ctx context.Context, p db.Person, repo db.Repository,
) (*gen.Person, error) {
	classesInfo, err := collectClassesInfo(ctx, p.ClassIDs, repo)
	if err != nil {
		return nil, err
	}

	return &gen.Person{
		Id:         p.ID.Hex(),
		Phone:      p.Phone,
		Name:       p.Name,
		PictureUri: p.PictureURI,
		BirthDate:  timestamppb.New(p.BirthDate),
		Gender:     convertDbGender(p.Gender),
		CreatedAt:  timestamppb.New(p.CreatedAt),
		UpdatedAt:  timestamppb.New(p.UpdatedAt),

		ClassesInfo: classesInfo,
	}, nil
}

func convertDbPersons(
	ctx context.Context, ps []db.Person, repo db.Repository,
) (res []*gen.Person, err error) {
	for _, p := range ps {
		r, err := convertDbPerson(ctx, p, repo)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func convertDbTrainer(
	ctx context.Context, t db.Trainer, repo db.Repository,
) (*gen.Trainer, error) {
	classesInfo, err := collectClassesInfo(ctx, t.ClassIDs, repo)
	if err != nil {
		return nil, err
	}

	studio, err := repo.GetStudio(ctx, t.StudioID)
	if err != nil {
		return nil, err
	}

	return &gen.Trainer{
		Id:         t.ID.Hex(),
		Phone:      t.Phone,
		Name:       t.Name,
		PictureUri: t.PictureURI,
		BirthDate:  timestamppb.New(t.BirthDate),
		Gender:     convertDbGender(t.Gender),
		CreatedAt:  timestamppb.New(t.CreatedAt),
		UpdatedAt:  timestamppb.New(t.UpdatedAt),

		ClassesInfo: classesInfo,
		StudioInfo: &gen.NameIDPair{
			Name: studio.Address,
			Id:   t.StudioID.Hex(),
		},
	}, nil
}

func convertDbTrainers(
	ctx context.Context, ts []db.Trainer, repo db.Repository,
) (res []*gen.Trainer, err error) {
	for _, t := range ts {
		r, err := convertDbTrainer(ctx, t, repo)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func convertDbStudio(
	ctx context.Context, s db.Studio, repo db.Repository,
) (*gen.Studio, error) {
	classesInfo, err := collectClassesInfo(ctx, s.ClassIDs, repo)
	if err != nil {
		return nil, err
	}

	trainersInfo, err := collectTrainersInfo(ctx, s.TrainerIDs, repo)
	if err != nil {
		return nil, err
	}

	return &gen.Studio{
		Id:        s.ID.Hex(),
		Address:   s.Address,
		CreatedAt: timestamppb.New(s.CreatedAt),
		UpdatedAt: timestamppb.New(s.UpdatedAt),

		ClassesInfo:  classesInfo,
		TrainersInfo: trainersInfo,
	}, nil
}

func convertDbStudios(
	ctx context.Context, ss []db.Studio, repo db.Repository,
) (res []*gen.Studio, err error) {
	for _, s := range ss {
		r, err := convertDbStudio(ctx, s, repo)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func convertDbClass(
	ctx context.Context, c db.Class, repo db.Repository,
) (*gen.Class, error) {
	clientsInfo, err := collectClientsInfo(ctx, c.ClientIDs, repo)
	if err != nil {
		return nil, err
	}

	studio, err := repo.GetStudio(ctx, c.StudioID)
	if err != nil {
		return nil, err
	}

	trainer, err := repo.GetTrainer(ctx, c.TrainerID)
	if err != nil {
		return nil, err
	}

	return &gen.Class{
		Id:        c.ID.Hex(),
		Name:      c.Name,
		Time:      timestamppb.New(c.Time),
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(c.UpdatedAt),

		StudioInfo: &gen.NameIDPair{
			Name: studio.Address,
			Id:   c.StudioID.Hex(),
		},
		TrainerInfo: &gen.NameIDPair{
			Name: trainer.Name,
			Id:   c.TrainerID.Hex(),
		},
		ClientsInfo: clientsInfo,
	}, nil
}

func convertDbClasses(
	ctx context.Context, cs []db.Class, repo db.Repository,
) (res []*gen.Class, err error) {
	for _, c := range cs {
		r, err := convertDbClass(ctx, c, repo)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}