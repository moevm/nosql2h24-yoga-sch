package db

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type SearchFilter bson.M

func (f SearchFilter) AddPaginationSettings(firstID, lastID string) (sortOrder int) {
	switch {
	case lastID != "":
		f["_id"] = bson.M{"$gt": objectIDFromHex(lastID)}
		return 1
	case firstID != "":
		f["_id"] = bson.M{"$lt": objectIDFromHex(firstID)}
		return -1
	default:
		return 0
	}
}

func (f SearchFilter) AddRegex(name, value string) {
	if value != "" {
		f[name] = bson.M{"$regex": value, "$options": "i"}
	}
}

func (f SearchFilter) AddTimeInterval(name string, value TimeInterval) {
	filter := bson.M{}
	if value.Begin != nil {
		filter["$gte"] = value.Begin
	}
	if value.End != nil {
		filter["$lt"] = value.End
	}
	if len(filter) > 0 {
		f[name] = filter
	}
}

func (f SearchFilter) AddSelector(name string, value []string) {
	if len(value) > 0 {
		f[name] = bson.M{"$in": value}
	}
}

func (f SearchFilter) AddIDsSelector(name string, value []bson.ObjectID) {
	if len(value) > 0 {
		f[name] = bson.M{"$in": value}
	}
}

func (f SearchFilter) AsBSON() bson.M {
	return bson.M(f)
}