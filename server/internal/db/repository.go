package db

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	DropDB(ctx context.Context) error
	GetCollectionNames(ctx context.Context) ([]string, error)
	ImportDB(ctx context.Context, data map[string][]bson.M) error
	ExportDB(ctx context.Context) (map[string][]bson.M, error)

	SearchClients(ctx context.Context, req ClientsFilter) (res []Person, pageInfo PageInfo, err error)
	SearchTrainers(ctx context.Context, req TrainersFilter) (res []Trainer, pageInfo PageInfo, err error)
	SearchStudios(ctx context.Context, req StudiosFilter) (res []Studio, pageInfo PageInfo, err error)
	SearchClasses(ctx context.Context, req ClassesFilter) (res []Class, pageInfo PageInfo, err error)

	GetIDByCreds(ctx context.Context, phone, password string) (bson.ObjectID, error)

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

	MakeAppointment(ctx context.Context, classID, clientID bson.ObjectID) error
	CancelAppointment(ctx context.Context, classID, clientID bson.ObjectID) error
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

var (
	ErrNotFound = errors.New("not found")
)

func NewMongoRepository(mg *mongo.Client) Repository {
	return MongoRepository{mg: mg}
}

func (r MongoRepository) DB() *mongo.Database {
	return r.mg.Database(dbName)
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