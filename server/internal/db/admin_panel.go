package db

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r MongoRepository) DropDB(ctx context.Context) error {
	return r.Db().Drop(ctx)
}

func (r MongoRepository) ImportDB(
	ctx context.Context,
	data map[string][]bson.M,
) error {
	if err := r.Db().Drop(ctx); err != nil {
		return err
	}

	for colName, docs := range data {
	    if len(docs) == 0 {
	        continue
	    }
		col := r.Db().Collection(colName)
		if _, err := col.InsertMany(ctx, docs); err != nil {
			return err
		}
	}
	return nil
}

func (r MongoRepository) ExportDB(
	ctx context.Context) (map[string][]bson.M, error) {
	db := r.Db()

	colNames, err := db.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	data := make(map[string][]bson.M, len(colNames))
	for _, colName := range colNames {
		col := db.Collection(colName)

		cur, err := col.Find(ctx, bson.D{})
		if err != nil {
			return nil, err
		}

		var docs []bson.M
		if err = cur.All(ctx, &docs); err != nil {
			return nil, err
		}

		if err := cur.Close(ctx); err != nil {
			return nil, err
		}

		data[colName] = docs
	}

	return data, nil
}