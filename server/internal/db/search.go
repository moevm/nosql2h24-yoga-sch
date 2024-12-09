package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TimeInterval struct {
	Begin *bson.DateTime
	End   *bson.DateTime
}

func SearchEntitiesByRegexName[T any](
	ctx context.Context, col *mongo.Collection, field string, regexes []string,
) (res []T, err error) {
	res = []T{}

	if len(regexes) == 0 {
		return nil, nil
	}

	var nameRegexes []bson.M
	for _, r := range regexes {
		if r == "" {
			continue
		}
		nameRegexes = append(nameRegexes, bson.M{
			field: bson.M{"$regex": r, "$options": "i"},
		})
	}
	filter := bson.M{"$or": nameRegexes}

	cur, err := col.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to search entities by name regex: %w", err)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		if err = cur.Close(ctx); err != nil {
			err = fmt.Errorf("failed to close cursor: %w", err)
		}
	}(cur, ctx)

	if err = cur.All(ctx, &res); err != nil {
		return nil, fmt.Errorf("failed to find entities by name regex: %w", err)
	}

	return res, nil
}

func IDsRegexFilter(field string, regexes []string) bson.D {
	var filter []bson.D
	for _, r := range regexes {
		filter = append(filter, bson.D{{
			"$regexMatch", bson.M{
				"input": bson.D{{"$toString", "$$elem_id"}},
				"regex": r,
			},
		}})
	}
	return bson.D{{
		"$match", bson.D{{
			"$expr", bson.D{{
				"$gt", bson.A{
					bson.D{{
						"$size", bson.D{{
							"$filter", bson.M{
								"input": fmt.Sprintf("$%s", field),
								"as":    "elem_id",
								"cond":  bson.D{{"$or", regexes}},
							},
						}},
					}},
				},
			}},
		}},
	}}
}

func IDRegexFilter(field string, regexes []string) bson.D {
	var filter []bson.D
	for _, r := range regexes {
		filter = append(filter, bson.D{{
			"$regexMatch", bson.M{
				"input": bson.D{{"$toString", fmt.Sprintf("$%s", field)}},
				"regex": r,
			},
		}})
	}
	return bson.D{{
		"$match", bson.D{{
			"$expr", bson.D{{
				"$or", regexes,
			}},
		}},
	}}
}

type ClientsFilter struct {
	IDSubstring         string
	PhoneSubstring      string
	NameSubstring       string
	PictureURISubstring string
	BirthDateInterval   TimeInterval
	Genders             []Gender
	CreatedAtInterval   TimeInterval
	UpdatedAtInterval   TimeInterval

	ClassIDSubstrings []string
}

func (r MongoRepository) SearchClients(
	ctx context.Context, req ClientsFilter,
) (res []Person, err error) {
	col := r.DB().Collection(clients)

	var pipeline mongo.Pipeline
	if req.IDSubstring != "" {
		pipeline = append(pipeline, bson.D{{
			"$addFields", bson.D{{
				"_id_string", bson.D{{
					"$toString", "$_id"}}}}},
		})
		pipeline = append(pipeline, bson.D{{
			"$match", bson.D{{
				"_id_string", bson.M{
					"$regex":   req.IDSubstring,
					"$options": "i",
				}}},
		}})
	}

	if len(req.ClassIDSubstrings) > 0 {
		pipeline = append(pipeline, IDsRegexFilter("class_ids", req.ClassIDSubstrings))
	}

	filter := SearchFilter{}
	filter.AddRegex("phone", req.PhoneSubstring)
	filter.AddRegex("name", req.NameSubstring)
	filter.AddRegex("picture_uri", req.PictureURISubstring)
	filter.AddTimeInterval("birth_date", req.BirthDateInterval)

	var genders []string
	if len(req.Genders) > 0 {
		for _, g := range req.Genders {
			genders = append(genders, string(g))
		}
	}
	filter.AddSelector("gender", genders)

	filter.AddTimeInterval("created_at", req.CreatedAtInterval)
	filter.AddTimeInterval("updated_at", req.UpdatedAtInterval)

	pipeline = append(pipeline, bson.D{{"$match", filter}})

	cur, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err = cur.Close(ctx)
	}(cur, ctx)

	if err = cur.All(ctx, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type TrainersFilter struct {
	IDSubstring         string
	PhoneSubstring      string
	NameSubstring       string
	PictureURISubstring string
	BirthDateInterval   TimeInterval
	Genders             []Gender
	CreatedAtInterval   TimeInterval
	UpdatedAtInterval   TimeInterval

	ClassIDSubstrings  []string
	StudioIDSubstrings []string
}

func (r MongoRepository) SearchTrainers(
	ctx context.Context, req TrainersFilter,
) (res []Trainer, err error) {
	col := r.DB().Collection(trainers)

	var pipeline mongo.Pipeline
	if req.IDSubstring != "" {
		pipeline = append(pipeline, bson.D{{
			"$addFields", bson.D{{
				"_id_string", bson.D{{
					"$toString", "$_id"}}}}},
		})
		pipeline = append(pipeline, bson.D{{
			"$match", bson.D{{
				"_id_string", bson.M{
					"$regex":   req.IDSubstring,
					"$options": "i",
				}}},
		}})
	}

	if len(req.ClassIDSubstrings) > 0 {
		pipeline = append(pipeline, IDsRegexFilter("class_ids", req.ClassIDSubstrings))
	}

	if len(req.StudioIDSubstrings) > 0 {
		pipeline = append(pipeline, IDRegexFilter("studio_id", req.StudioIDSubstrings))
	}

	filter := SearchFilter{}
	filter.AddRegex("phone", req.PhoneSubstring)
	filter.AddRegex("name", req.NameSubstring)
	filter.AddRegex("picture_uri", req.PictureURISubstring)
	filter.AddTimeInterval("birth_date", req.BirthDateInterval)

	var genders []string
	if len(req.Genders) > 0 {
		for _, g := range req.Genders {
			genders = append(genders, string(g))
		}
	}
	filter.AddSelector("gender", genders)

	filter.AddTimeInterval("created_at", req.CreatedAtInterval)
	filter.AddTimeInterval("updated_at", req.UpdatedAtInterval)

	pipeline = append(pipeline, bson.D{{"$match", filter}})

	cur, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err = cur.Close(ctx)
	}(cur, ctx)

	if err = cur.All(ctx, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type StudiosFilter struct {
	IDSubstring       string
	AddressSubstring  string
	CreatedAtInterval TimeInterval
	UpdatedAtInterval TimeInterval

	ClassIDSubstrings   []string
	TrainerIDSubstrings []string
}

func (r MongoRepository) SearchStudios(
	ctx context.Context, req StudiosFilter,
) (res []Studio, err error) {
	col := r.DB().Collection(studios)

	var classIDs []bson.ObjectID
	switch filteredClasses, err := SearchEntitiesByRegexName[Class](
		ctx, r.DB().Collection(classes), "name", req.ClassIDSubstrings); {
	case err != nil:
		return nil, err
	case filteredClasses == nil:
		break
	case len(filteredClasses) == 0:
		return res, nil
	default:
		for _, e := range filteredClasses {
			classIDs = append(classIDs, e.ID)
		}
	}

	var trainerIDs []bson.ObjectID
	switch filteredTrainers, err := SearchEntitiesByRegexName[Trainer](
		ctx, r.DB().Collection(trainers), "name", req.TrainerIDSubstrings); {
	case err != nil:
		return nil, err
	case filteredTrainers == nil:
		break
	case len(filteredTrainers) == 0:
		return res, nil
	default:
		for _, e := range filteredTrainers {
			trainerIDs = append(trainerIDs, e.ID)
		}
	}

	var pipeline mongo.Pipeline
	if req.IDSubstring != "" {
		pipeline = append(pipeline, bson.D{{
			"$addFields", bson.D{{
				"_id_string", bson.D{{
					"$toString", "$_id"}}}}},
		})
		pipeline = append(pipeline, bson.D{{
			"$match", bson.D{{
				"_id_string", bson.M{
					"$regex":   req.IDSubstring,
					"$options": "i",
				}}},
		}})
	}

	filter := SearchFilter{}
	filter.AddRegex("address", req.AddressSubstring)
	filter.AddTimeInterval("created_at", req.CreatedAtInterval)
	filter.AddTimeInterval("updated_at", req.UpdatedAtInterval)
	filter.AddIDsSelector("class_ids", classIDs)
	filter.AddIDsSelector("trainer_ids", trainerIDs)

	pipeline = append(pipeline, bson.D{{"$match", filter}})

	cur, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err = cur.Close(ctx)
	}(cur, ctx)

	if err = cur.All(ctx, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type ClassesFilter struct {
	IDSubstring       string
	NameSubstring     string
	TimeInterval      TimeInterval
	CreatedAtInterval TimeInterval
	UpdatedAtInterval TimeInterval

	StudioIDSubstrings  []string
	TrainerIDSubstrings []string
	ClientIDSubstrings  []string
}

func (r MongoRepository) SearchClasses(
	ctx context.Context, req ClassesFilter,
) (res []Class, err error) {
	col := r.DB().Collection(classes)

	var studioIDs []bson.ObjectID
	switch filteredStudios, err := SearchEntitiesByRegexName[Studio](
		ctx, r.DB().Collection(studios), "address", req.StudioIDSubstrings); {
	case err != nil:
		return nil, err
	case filteredStudios == nil:
		break
	case len(filteredStudios) == 0:
		return res, nil
	default:
		for _, e := range filteredStudios {
			studioIDs = append(studioIDs, e.ID)
		}
	}

	var trainerIDs []bson.ObjectID
	switch filteredTrainers, err := SearchEntitiesByRegexName[Trainer](
		ctx, r.DB().Collection(trainers), "name", req.TrainerIDSubstrings); {
	case err != nil:
		return nil, err
	case filteredTrainers == nil:
		break
	case len(filteredTrainers) == 0:
		return res, nil
	default:
		for _, e := range filteredTrainers {
			trainerIDs = append(trainerIDs, e.ID)
		}
	}

	var clientIDs []bson.ObjectID
	switch filteredClients, err := SearchEntitiesByRegexName[Client](
		ctx, r.DB().Collection(clients), "name", req.ClientIDSubstrings); {
	case err != nil:
		return nil, err
	case filteredClients == nil:
		break
	case len(filteredClients) == 0:
		return res, nil
	default:
		for _, e := range filteredClients {
			clientIDs = append(clientIDs, e.ID)
		}
	}

	var pipeline mongo.Pipeline
	if req.IDSubstring != "" {
		pipeline = append(pipeline, bson.D{{
			"$addFields", bson.D{{
				"_id_string", bson.D{{
					"$toString", "$_id"}}}}},
		})
		pipeline = append(pipeline, bson.D{{
			"$match", bson.D{{
				"_id_string", bson.M{
					"$regex":   req.IDSubstring,
					"$options": "i",
				}}},
		}})
	}

	filter := SearchFilter{}
	filter.AddRegex("name", req.NameSubstring)
	filter.AddTimeInterval("time", req.TimeInterval)
	filter.AddTimeInterval("created_at", req.CreatedAtInterval)
	filter.AddTimeInterval("updated_at", req.UpdatedAtInterval)
	filter.AddIDsSelector("studio_id", studioIDs)
	filter.AddIDsSelector("trainer_id", trainerIDs)
	filter.AddIDsSelector("client_ids", clientIDs)

	pipeline = append(pipeline, bson.D{{"$match", filter}})

	cur, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err = cur.Close(ctx)
	}(cur, ctx)

	if err = cur.All(ctx, &res); err != nil {
		return nil, err
	}

	return res, nil
}