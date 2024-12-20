package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r MongoRepository) MakeAppointment(
	ctx context.Context, classID, clientID bson.ObjectID,
) error {
	classesCollection := r.DB().Collection(classes)
	clientsCollection := r.DB().Collection(clients)

	if err := classesCollection.FindOneAndUpdate(ctx,
		bson.M{
			"_id": classID,
			"$expr": bson.M{
				"$lt": bson.A{bson.M{"$size": bson.D{{
					"$ifNull", bson.A{"$client_ids", bson.A{}},
				}}}, MaxClientCount},
			},
		},
		bson.M{"$addToSet": bson.M{"client_ids": clientID}},
	).Err(); err != nil {
		if errors.As(err, &mongo.ErrNoDocuments) {
			return ErrNotFound
		}
		return fmt.Errorf("error updating class %s: %v", classID, err)
	}

	if err := clientsCollection.FindOneAndUpdate(ctx, bson.M{"_id": clientID},
		bson.M{"$addToSet": bson.M{"class_ids": classID}}).Err(); err != nil {
		return fmt.Errorf("error updating client %s: %v", clientID, err)
	}

	return nil
}

func (r MongoRepository) CancelAppointment(
	ctx context.Context, classID, clientID bson.ObjectID,
) error {
	classesCollection := r.DB().Collection(classes)
	clientsCollection := r.DB().Collection(clients)

	if _, err := classesCollection.UpdateOne(ctx, bson.M{"_id": classID},
		bson.M{"$pull": bson.M{"client_ids": clientID}}); err != nil {
		return fmt.Errorf("error updating class %s: %v", classID, err)
	}

	if _, err := clientsCollection.UpdateOne(ctx, bson.M{"_id": clientID},
		bson.M{"$pull": bson.M{"class_ids": classID}}); err != nil {
		return fmt.Errorf("error updating client %s: %v", clientID, err)
	}

	return nil
}