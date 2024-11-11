package db

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Class struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	Time      time.Time     `bson:"time"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`

	TrainerID bson.ObjectID   `bson:"trainer_id"`
	StudioID  bson.ObjectID   `bson:"studio_id"`
	ClientIDs []bson.ObjectID `bson:"client_ids"`
}

func (r MongoRepository) InsertClass(
	ctx context.Context, class Class,
) (bson.ObjectID, error) {
	class.CreatedAt = time.Now()
	class.UpdatedAt = time.Now()

	studiosCollection := r.Db().Collection(studios)
	trainersCollection := r.Db().Collection(trainers)
	classesCollection := r.Db().Collection(classes)

	filter := bson.M{"name": class.Name, "time": class.Time}
	update := bson.M{"$setOnInsert": class}
	opts := options.FindOneAndUpdate().
		SetUpsert(true).SetReturnDocument(options.After)

	var result Class
	if err := classesCollection.FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&result); err != nil {
		return bson.ObjectID{}, fmt.Errorf("failed to insert class: %w", err)
	}

	if err := studiosCollection.FindOneAndUpdate(ctx, bson.M{"_id": class.StudioID},
		bson.M{"$addToSet": bson.M{"class_ids": result.ID}}).Err(); err != nil {
		return bson.ObjectID{}, fmt.Errorf("failed to update studio: %w", err)
	}

	if err := trainersCollection.FindOneAndUpdate(ctx, bson.M{"_id": class.TrainerID},
		bson.M{"$addToSet": bson.M{"class_ids": result.ID}}).Err(); err != nil {
		return bson.ObjectID{}, fmt.Errorf("failed to update trainer: %w", err)
	}

	fmt.Printf("Found or inserted class with id %v\n", result.ID)
	return result.ID, nil
}

func (r MongoRepository) GetClasses(ctx context.Context) ([]Class, error) {
	collection := r.Db().Collection(classes)

	classesCursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find classes: %w", err)
	}
	defer func(classesCursor *mongo.Cursor, ctx context.Context) {
		err = classesCursor.Close(ctx)
		if err != nil {
			slog.Warn("failed to close classes: %v", err)
		}
	}(classesCursor, ctx)

	var classes []Class
	if err = classesCursor.All(ctx, &classes); err != nil {
		return nil, fmt.Errorf("failed to find classes: %w", err)
	}

	return classes, nil
}

func (r MongoRepository) GetClass(ctx context.Context, id bson.ObjectID) (Class, error) {
	collection := r.Db().Collection(classes)

	var class Class
	if err := collection.FindOne(ctx, bson.M{"_id": id}).
		Decode(&class); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Class{}, ErrNotFound
		}
		return Class{}, fmt.Errorf("failed to find class: %w", err)
	}

	return class, nil
}

func (r MongoRepository) DeleteClass(ctx context.Context, id bson.ObjectID) error {
	classesCollection := r.Db().Collection(classes)
	studiosCollection := r.Db().Collection(studios)
	trainersCollection := r.Db().Collection(trainers)
	clientsCollection := r.Db().Collection(clients)

	var class Class
	if err := classesCollection.FindOne(ctx, bson.M{"_id": id}).
		Decode(&class); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrNotFound
		}
		return fmt.Errorf("failed to find class: %w", err)
	}

	for _, clientID := range class.ClientIDs {
		if _, err := clientsCollection.UpdateOne(ctx, bson.M{"_id": clientID},
			bson.M{"$pull": bson.M{"class_ids": id}}); err != nil {
			return fmt.Errorf("failed to update client: %w", err)
		}
	}

	if _, err := studiosCollection.UpdateOne(ctx, bson.M{"_id": class.StudioID},
		bson.M{"$pull": bson.M{"class_ids": id}}); err != nil {
		return fmt.Errorf("failed to update studio: %w", err)
	}

	if _, err := trainersCollection.UpdateOne(ctx, bson.M{"_id": class.TrainerID},
		bson.M{"$pull": bson.M{"class_ids": id}}); err != nil {
		return fmt.Errorf("failed to update trainer: %w", err)
	}

	if _, err := classesCollection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return fmt.Errorf("failed to delete class: %w", err)
	}

	return nil
}