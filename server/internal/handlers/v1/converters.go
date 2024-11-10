package v1

import (
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

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