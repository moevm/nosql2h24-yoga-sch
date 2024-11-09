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

type Repository interface {
	InsertClient(ctx context.Context, client Client) (bson.ObjectID, error)
	GetClients(ctx context.Context) ([]Client, error)
	GetClient(ctx context.Context, id bson.ObjectID) (Client, error)
	DeleteClient(ctx context.Context, id bson.ObjectID) error
}

type MongoRepository struct {
	mg *mongo.Client
}

const (
	dbName = "fitness_aggregator"

	studios  = "studios"
	clients  = "clients"
	trainers = "trainers"
	classes  = "classes"
)

var ErrNotFound = errors.New("not found")

func (r MongoRepository) Db() *mongo.Database {
	return r.mg.Database(dbName)
}

func NewMongoRepository(mg *mongo.Client) Repository {
	return MongoRepository{mg: mg}
}

type Gender string

const (
	GenderFemale Gender = "F"
	GenderMale   Gender = "M"
)

type Person struct {
	ID         bson.ObjectID   `bson:"_id,omitempty"`
	Name       string          `bson:"name"`
	Phone      string          `bson:"phone"`
	PictureURI string          `bson:"picture_uri"`
	BirthDate  time.Time       `bson:"birth_date"`
	Gender     Gender          `bson:"gender"`
	Classes    []bson.ObjectID `bson:"classes"`
	CreatedAt  time.Time       `bson:"created_at"`
	UpdatedAt  time.Time       `bson:"updated_at"`
}

type Client struct {
	Person   `bson:",inline"`
	Password string `bson:"password"`
}

func (r MongoRepository) InsertClient(
	ctx context.Context, client Client,
) (bson.ObjectID, error) {
	client.CreatedAt = time.Now()
	client.UpdatedAt = time.Now()

	collection := r.mg.Database("yoga").Collection(clients)

	filter := bson.M{"phone": client.Person.Phone, "password": client.Password}
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
	collection := r.mg.Database("yoga").Collection(clients)

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
	collection := r.mg.Database("yoga").Collection(clients)

	var client Client
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&client); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Client{}, ErrNotFound
		}
		return Client{}, fmt.Errorf("failed to find client: %w", err)
	}

	return client, nil
}

func (r MongoRepository) DeleteClient(ctx context.Context, id bson.ObjectID) error {
	collection := r.mg.Database("yoga").Collection(clients)

	if _, err := collection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrNotFound
		}
		return fmt.Errorf("failed to delete client: %w", err)
	}

	return nil
}