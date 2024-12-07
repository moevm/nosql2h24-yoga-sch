package v1

import (
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

func convertDbPerson(p db.Person) *gen.Person {
	var classIDs []string
	for _, id := range p.ClassIDs {
		classIDs = append(classIDs, id.Hex())
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

		ClassIds: classIDs,
	}
}

func convertDbPersons(ps []db.Person) (res []*gen.Person) {
	for _, p := range ps {
		res = append(res, convertDbPerson(p))
	}
	return res
}

func convertDbTrainer(t db.Trainer) *gen.Trainer {
	var classIDs []string
	for _, id := range t.ClassIDs {
		classIDs = append(classIDs, id.Hex())
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

		ClassIds: classIDs,

		StudioId: t.StudioID.Hex(),
	}
}

func convertDbStudio(s db.Studio) *gen.Studio {
	var classIDs []string
	for _, id := range s.ClassIDs {
		classIDs = append(classIDs, id.Hex())
	}

	var trainerIDs []string
	for _, id := range s.TrainerIDs {
		trainerIDs = append(trainerIDs, id.Hex())
	}

	return &gen.Studio{
		Id:        s.ID.Hex(),
		Address:   s.Address,
		CreatedAt: timestamppb.New(s.CreatedAt),
		UpdatedAt: timestamppb.New(s.UpdatedAt),

		ClassIds:   classIDs,
		TrainerIds: trainerIDs,
	}
}

func convertDbClass(c db.Class) *gen.Class {
	var clientIDs []string
	for _, id := range c.ClientIDs {
		clientIDs = append(clientIDs, id.Hex())
	}

	return &gen.Class{
		Id:        c.ID.Hex(),
		Name:      c.Name,
		Time:      timestamppb.New(c.Time),
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(c.UpdatedAt),

		StudioId:  c.StudioID.Hex(),
		TrainerId: c.TrainerID.Hex(),
		ClientIds: clientIDs,
	}
}