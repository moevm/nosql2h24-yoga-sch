package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

func (c Client) IsPasswordValid(password string) bool {
	return c.Password == password
}

func (r MongoRepository) GetIDByCreds(ctx context.Context, phone, password string) (bson.ObjectID, error) {
	collection := r.DB().Collection(clients)

	var client Client
	if err := collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&client); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return bson.ObjectID{}, ErrNotFound
		}
		return bson.ObjectID{}, fmt.Errorf("error finding client by phone %s: %v", phone, err)
	}

	if !client.IsPasswordValid(password) {
		return bson.ObjectID{}, ErrInvalidCredentials
	}

	return client.ID, nil
}