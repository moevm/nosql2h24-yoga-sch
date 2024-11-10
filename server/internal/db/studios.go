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

type Studio struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	Address   string        `bson:"address"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`

	ClassIDs   []bson.ObjectID `bson:"class_ids"`
	TrainerIDs []bson.ObjectID `bson:"trainer_ids"`
}

func (r MongoRepository) InsertStudio(
	ctx context.Context, studio Studio,
) (bson.ObjectID, error) {
	studio.CreatedAt = time.Now()
	studio.UpdatedAt = time.Now()

	collection := r.Db().Collection(studios)

	filter := bson.M{"name": studio.Name, "address": studio.Address}
	update := bson.M{"$setOnInsert": studio}
	opts := options.FindOneAndUpdate().
		SetUpsert(true).SetReturnDocument(options.After)

	var result Studio
	if err := collection.FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&result); err != nil {
		return bson.ObjectID{}, fmt.Errorf("failed to insert studio: %w", err)
	}

	fmt.Printf("Found or inserted studio with id %v\n", result.ID)
	return result.ID, nil
}

func (r MongoRepository) GetStudios(ctx context.Context) ([]Studio, error) {
	collection := r.Db().Collection(studios)

	studiosCursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find studios: %w", err)
	}
	defer func(studiosCursor *mongo.Cursor, ctx context.Context) {
		err = studiosCursor.Close(ctx)
		if err != nil {
			slog.Warn("failed to close studios: %v", err)
		}
	}(studiosCursor, ctx)

	var studios []Studio
	if err = studiosCursor.All(ctx, &studios); err != nil {
		return nil, fmt.Errorf("failed to find studios: %w", err)
	}

	return studios, nil
}

func (r MongoRepository) GetStudio(ctx context.Context, id bson.ObjectID) (Studio, error) {
	collection := r.Db().Collection(studios)

	var studio Studio
	if err := collection.FindOne(ctx, bson.M{"_id": id}).
		Decode(&studio); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Studio{}, ErrNotFound
		}
		return Studio{}, fmt.Errorf("failed to find studio: %w", err)
	}

	return studio, nil
}

func (r MongoRepository) DeleteStudio(ctx context.Context, id bson.ObjectID) error {
	studiosCollection := r.Db().Collection(studios)
	trainersCollection := r.Db().Collection(trainers)
	classesCollection := r.Db().Collection(classes)

	var studio Studio
	if err := studiosCollection.FindOne(ctx, bson.M{"_id": id}).
		Decode(&studio); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrNotFound
		}
		return fmt.Errorf("failed to find studio: %w", err)
	}

	for _, trainerID := range studio.TrainerIDs {
		if _, err := trainersCollection.DeleteOne(ctx, bson.M{"_id": trainerID}); err != nil {
			return fmt.Errorf("failed to delete trainer: %w", err)
		}
	}

	for _, classID := range studio.ClassIDs {
		if _, err := classesCollection.DeleteOne(ctx, bson.M{"_id": classID}); err != nil {
			return fmt.Errorf("failed to delete class: %w", err)
		}
	}

	if _, err := studiosCollection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return fmt.Errorf("failed to delete studio: %w", err)
	}

	return nil
}