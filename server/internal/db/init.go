package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func New() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://db:27017").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	var res bson.M
	if err := client.Database("admin").
		RunCommand(context.Background(), bson.D{{"ping", 1}}).
		Decode(&res); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

func objectIDFromHex(val string) bson.ObjectID {
	id, err := bson.ObjectIDFromHex(val)
	if err != nil {
		panic(fmt.Sprintf("ObjectIDFromHex: %v", err))
	}
	return id
}

func timeFromRFC3339(val string) time.Time {
	tm, err := time.Parse(time.RFC3339, val)
	if err != nil {
		panic(fmt.Sprintf("TimeFromRFC3339: %v", err))
	}
	return tm
}

func ImportDB(ctx context.Context, c *mongo.Client) {
	repo := NewMongoRepository(c)
	if err := repo.ImportDB(ctx, map[string][]bson.M{
		clients: {
			{
				"_id":  objectIDFromHex("client_id_1"),
				"name": "Mike",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("class_id_1"),
					objectIDFromHex("class_id_2"),
				},
			},
			{
				"_id":  objectIDFromHex("client_id_2"),
				"name": "Mike",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("class_id_1"),
					objectIDFromHex("class_id_2"),
				},
			},
		},
		classes: {
			{
				"_id":        objectIDFromHex("class_id_1"),
				"time":       timeFromRFC3339("2024-10-10T13:00:00Z"),
				"created_at": time.Now(),
			},
		},
	}); err != nil {
		panic(fmt.Sprintf("Failed to import DB: %v", err))
	}
}