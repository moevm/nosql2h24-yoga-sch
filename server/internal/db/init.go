package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func New() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)

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