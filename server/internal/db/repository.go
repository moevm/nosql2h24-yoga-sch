package db

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	InsertClient(ctx context.Context, client Client) (bson.ObjectID, error)
	GetClients(ctx context.Context) ([]Client, error)
	GetClient(ctx context.Context, id bson.ObjectID) (Client, error)
	DeleteClient(ctx context.Context, id bson.ObjectID) error

	InsertTrainer(ctx context.Context, trainer Trainer) (bson.ObjectID, error)
	GetTrainers(ctx context.Context) ([]Trainer, error)
	GetTrainer(ctx context.Context, id bson.ObjectID) (Trainer, error)
	DeleteTrainer(ctx context.Context, id bson.ObjectID) error

	InsertStudio(ctx context.Context, studio Studio) (bson.ObjectID, error)
	GetStudios(ctx context.Context) ([]Studio, error)
	GetStudio(ctx context.Context, id bson.ObjectID) (Studio, error)
	DeleteStudio(ctx context.Context, id bson.ObjectID) error

	InsertClass(ctx context.Context, class Class) (bson.ObjectID, error)
	GetClasses(ctx context.Context) ([]Class, error)
	GetClass(ctx context.Context, id bson.ObjectID) (Class, error)
	DeleteClass(ctx context.Context, id bson.ObjectID) error
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
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Name       string        `bson:"name"`
	Phone      string        `bson:"phone"`
	PictureURI string        `bson:"picture_uri"`
	BirthDate  time.Time     `bson:"birth_date"`
	Gender     Gender        `bson:"gender"`
	CreatedAt  time.Time     `bson:"created_at"`
	UpdatedAt  time.Time     `bson:"updated_at"`

	ClassIDs []bson.ObjectID `bson:"class_ids"`
}