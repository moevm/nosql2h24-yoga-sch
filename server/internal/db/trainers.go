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

	studiosCollection := r.Db().Collection(studios)
	trainersCollection := r.Db().Collection(trainers)

	filter := bson.M{"phone": trainer.Person.Phone}
	update := bson.M{"$setOnInsert": trainer}
	opts := options.FindOneAndUpdate().
		SetUpsert(true).SetReturnDocument(options.After)

	var result Person
	if err := trainersCollection.FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&result); err != nil {
		return bson.ObjectID{}, fmt.Errorf("failed to insert trainer: %w", err)
	}

	if err := studiosCollection.FindOneAndUpdate(ctx, bson.M{"_id": trainer.StudioID},
		bson.M{"$push": bson.M{"trainer_ids": result.ID}}).Err(); err != nil {
		return bson.ObjectID{}, fmt.Errorf("failed to update studio: %w", err)
	}

	fmt.Printf("Found or inserted trainer with id %v\n", result.ID)
	return result.ID, nil
}

func (r MongoRepository) GetTrainers(ctx context.Context) ([]Trainer, error) {
	collection := r.Db().Collection(trainers)

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
	collection := r.Db().Collection(trainers)

	var trainer Trainer
	if err := collection.FindOne(ctx, bson.M{"_id": id}).
		Decode(&trainer); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Trainer{}, ErrNotFound
		}
		return Trainer{}, fmt.Errorf("failed to find trainer: %w", err)
	}

	return trainer, nil
}

func (r MongoRepository) DeleteTrainer(ctx context.Context, id bson.ObjectID) error {
	trainersCollection := r.Db().Collection(trainers)
	studiosCollection := r.Db().Collection(studios)

	var trainer Trainer
	if err := trainersCollection.FindOne(ctx, bson.M{"_id": id}).
		Decode(&trainer); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrNotFound
		}
		return fmt.Errorf("failed to find trainer: %w", err)
	}

	if _, err := studiosCollection.UpdateOne(ctx, bson.M{"_id": trainer.StudioID},
		bson.M{"$pull": bson.M{"trainer_ids": id}}); err != nil {
		return fmt.Errorf("failed to update studio: %w", err)
	}

	if _, err := trainersCollection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return fmt.Errorf("failed to delete trainer: %w", err)
	}

	return nil
}