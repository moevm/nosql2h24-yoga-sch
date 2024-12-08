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

type Client struct {
	Person   `bson:",inline"`
	Password string `bson:"password"`
}

func (r MongoRepository) InsertClient(
	ctx context.Context, client Client,
) (bson.ObjectID, error) {
	client.CreatedAt = time.Now()
	client.UpdatedAt = time.Now()

	collection := r.DB().Collection(clients)

	filter := bson.M{"phone": client.Person.Phone}
	update := bson.M{"$setOnInsert": client}
	opts := options.FindOneAndUpdate().
		SetUpsert(true).SetReturnDocument(options.After)

	var result Person
	if err := collection.FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&result); err != nil {
		return bson.ObjectID{}, fmt.Errorf("failed to insert client: %w", err)
	}

	fmt.Printf("Found or inserted client with id %v\n", result.ID)
	return result.ID, nil
}

func (r MongoRepository) GetClients(ctx context.Context) ([]Client, error) {
	collection := r.DB().Collection(clients)

	clientsCursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find clients: %w", err)
	}
	defer func(clientsCursor *mongo.Cursor, ctx context.Context) {
		err = clientsCursor.Close(ctx)
		if err != nil {
			slog.Warn("failed to close clients: %v", err)
		}
	}(clientsCursor, ctx)

	var clients []Client
	if err = clientsCursor.All(ctx, &clients); err != nil {
		return nil, fmt.Errorf("failed to find clients: %w", err)
	}

	if err = clientsCursor.Err(); err != nil {
		return nil, fmt.Errorf("failed to find clients: %w", err)
	}

	return clients, nil
}

func (r MongoRepository) GetClient(ctx context.Context, id bson.ObjectID) (Client, error) {
	collection := r.DB().Collection(clients)

	var client Client
	if err := collection.FindOne(ctx, bson.M{"_id": id}).
		Decode(&client); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Client{}, ErrNotFound
		}
		return Client{}, fmt.Errorf("failed to find client: %w", err)
	}

	return client, nil
}

func (r MongoRepository) DeleteClient(ctx context.Context, id bson.ObjectID) error {
	collection := r.DB().Collection(clients)

	switch res, err := collection.DeleteOne(ctx, bson.M{"_id": id}); {
	case err != nil:
		return fmt.Errorf("failed to delete client: %w", err)
	case res.DeletedCount == 0:
		return ErrNotFound
	}

	return nil
}