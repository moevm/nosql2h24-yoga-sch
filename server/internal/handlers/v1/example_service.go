package v1

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

type ExampleService struct {
	gen.UnimplementedExampleServiceServer

	DbC *mongo.Client
}

func (s *ExampleService) Echo(_ context.Context, req *gen.EchoRequest) (*gen.EchoResponse, error) {
	fmt.Printf("Echo: %s\n", req.Value)
	return &gen.EchoResponse{Value: req.Value}, nil
}

type ExampleRow struct {
	ID    bson.ObjectID `bson:"_id,omitempty"`
	Key   string        `bson:"key"`
	Value string        `bson:"value"`
}

func (s *ExampleService) Store(ctx context.Context, req *gen.StoreRequest) (*gen.StoreResponse, error) {
	fmt.Printf("Store: %s — %s\n", req.Key, req.Value)

	res, err := s.DbC.Database("example").
		Collection("example").
		InsertOne(ctx, ExampleRow{Key: req.Key, Value: req.Value})
	if err != nil {
		return nil, err
	}

	insertedID := res.InsertedID.(bson.ObjectID)
	return &gen.StoreResponse{Id: insertedID.Hex()}, nil
}

func (s *ExampleService) Load(ctx context.Context, req *gen.LoadRequest) (*gen.LoadResponse, error) {
	fmt.Printf("Load: %s\n", req.Key)

	var res ExampleRow
	if err := s.DbC.Database("example").
		Collection("example").
		FindOne(ctx, bson.D{{"key", req.Key}}).
		Decode(&res); err != nil {
		return nil, err
	}

	return &gen.LoadResponse{
		Key:   res.Key,
		Value: res.Value,
	}, nil
}