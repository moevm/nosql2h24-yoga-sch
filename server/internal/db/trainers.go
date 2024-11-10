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

type Trainer struct {
	Person   `bson:",inline"`
	StudioID bson.ObjectID `bson:"studio_id"`
}

func (r MongoRepository) InsertTrainer(
	ctx context.Context, trainer Trainer,
) (bson.ObjectID, error) {
	trainer.CreatedAt = time.Now()
	trainer.UpdatedAt = time.Now()

	collection := r.mg.Database("yoga").Collection(trainers)

	filter := bson.M{"phone": trainer.Person.Phone}
	update := bson.M{"$setOnInsert": trainer}
	opts := options.FindOneAndUpdate().
		SetUpsert(true).SetReturnDocument(options.After)

	var result Person
	if err := collection.FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&result); err != nil {
		return bson.ObjectID{}, fmt.Errorf("failed to insert trainer: %w", err)
	}

	fmt.Printf("Found or inserted trainer with id %v\n", result.ID)
	return result.ID, nil
}

func (r MongoRepository) GetTrainers(ctx context.Context) ([]Trainer, error) {
	collection := r.mg.Database("yoga").Collection(trainers)

	trainersCursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find trainers: %w", err)
	}
	defer func(trainersCursor *mongo.Cursor, ctx context.Context) {
		err = trainersCursor.Close(ctx)
		if err != nil {
			slog.Warn("failed to close trainers: %v", err)
		}
	}(trainersCursor, ctx)

	var trainers []Trainer
	if err = trainersCursor.All(ctx, &trainers); err != nil {
		return nil, fmt.Errorf("failed to find trainers: %w", err)
	}

	return trainers, nil
}

func (r MongoRepository) GetTrainer(ctx context.Context, id bson.ObjectID) (Trainer, error) {
	collection := r.mg.Database("yoga").Collection(trainers)

	var trainer Trainer
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&trainer); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Trainer{}, ErrNotFound
		}
		return Trainer{}, fmt.Errorf("failed to find trainer: %w", err)
	}

	return trainer, nil
}

func (r MongoRepository) DeleteTrainer(ctx context.Context, id bson.ObjectID) error {
	collection := r.mg.Database("yoga").Collection(trainers)

	switch res, err := collection.DeleteOne(ctx, bson.M{"_id": id}); {
	case errors.Is(err, mongo.ErrNoDocuments) || res.DeletedCount == 0:
		return ErrNotFound
	case err != nil:
		return fmt.Errorf("failed to delete trainer: %w", err)
	}

	return nil
}