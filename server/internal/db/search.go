package db

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TimeInterval struct {
	Begin *bson.DateTime
	End   *bson.DateTime
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

	ClassNameSubstrings []string
}

func (r MongoRepository) SearchClients(
	ctx context.Context, req ClientsFilter,
) (res []Person, err error) {
	col := r.Db().Collection(clients)

	var classIDs []bson.ObjectID
	switch filteredClasses, err := r.SearchClassesByNameRegex(
		ctx, req.ClassNameSubstrings); {
	case err != nil:
		return nil, err
	case len(filteredClasses) > 0:
		for _, class := range filteredClasses {
			classIDs = append(classIDs, class.ID)
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
	filter.AddIDsSelector("class_ids", classIDs)

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