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
		studios: {
			{
				"_id":        objectIDFromHex("507f1f77bcf86cd799439011"),
				"address":    "ul. Popova 1",
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"class_ids": []bson.ObjectID{
					objectIDFromHex("009f1f77bcf86cd799439011"),
					objectIDFromHex("019f1f77bcf86cd799439011"),
					objectIDFromHex("029f1f77bcf86cd799439011"),
					objectIDFromHex("039f1f77bcf86cd799439011"),
				},
				"trainer_ids": []bson.ObjectID{
					objectIDFromHex("407f1f77bcf86cd799439011"),
					objectIDFromHex("417f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("507f1f77bcf86cd799439012"),
				"address":    "ul. Lesnaya 2",
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"class_ids": []bson.ObjectID{
					objectIDFromHex("049f1f77bcf86cd799439011"),
					objectIDFromHex("059f1f77bcf86cd799439011"),
					objectIDFromHex("069f1f77bcf86cd799439011"),
					objectIDFromHex("079f1f77bcf86cd799439011"),
				},
				"trainer_ids": []bson.ObjectID{
					objectIDFromHex("427f1f77bcf86cd799439011"),
					objectIDFromHex("437f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("507f1f77bcf86cd799439013"),
				"address":    "ul. Kosmonavtov 3",
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"class_ids": []bson.ObjectID{
					objectIDFromHex("089f1f77bcf86cd799439011"),
					objectIDFromHex("099f1f77bcf86cd799439011"),
					objectIDFromHex("109f1f77bcf86cd799439011"),
					objectIDFromHex("119f1f77bcf86cd799439011"),
				},
				"trainer_ids": []bson.ObjectID{
					objectIDFromHex("447f1f77bcf86cd799439011"),
					objectIDFromHex("457f1f77bcf86cd799439011"),
				},
			},
		},
		trainers: {
			{
				"_id":         objectIDFromHex("407f1f77bcf86cd799439011"),
				"name":        "Boris Va",
				"phone":       "+7(999)444-4444",
				"gender":      "MALE",
				"birth_date":  timeFromRFC3339("2001-10-28T23:58:18Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("009f1f77bcf86cd799439011"),
					objectIDFromHex("019f1f77bcf86cd799439011"),
				},
				"studio_id": objectIDFromHex("507f1f77bcf86cd799439011"),
			},
			{
				"_id":         objectIDFromHex("417f1f77bcf86cd799439011"),
				"name":        "Egor Shmatcko",
				"phone":       "+7(999)333-3333",
				"gender":      "MALE",
				"birth_date":  timeFromRFC3339("2002-09-11T23:52:14Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("029f1f77bcf86cd799439011"),
					objectIDFromHex("039f1f77bcf86cd799439011"),
				},
				"studio_id": objectIDFromHex("507f1f77bcf86cd799439011"),
			},
			{
				"_id":         objectIDFromHex("427f1f77bcf86cd799439011"),
				"name":        "Oleg Gi",
				"phone":       "+7(999)222-2222",
				"gender":      "MALE",
				"birth_date":  timeFromRFC3339("2012-04-22T23:52:14Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("049f1f77bcf86cd799439011"),
					objectIDFromHex("059f1f77bcf86cd799439011"),
				},
				"studio_id": objectIDFromHex("507f1f77bcf86cd799439012"),
			},
			{
				"_id":         objectIDFromHex("437f1f77bcf86cd799439011"),
				"name":        "Vera Hans",
				"phone":       "+7(999)111-1111",
				"gender":      "FEMALE",
				"birth_date":  timeFromRFC3339("2008-03-03T23:52:14Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("069f1f77bcf86cd799439011"),
					objectIDFromHex("079f1f77bcf86cd799439011"),
				},
				"studio_id": objectIDFromHex("507f1f77bcf86cd799439012"),
			},
			{
				"_id":         objectIDFromHex("447f1f77bcf86cd799439011"),
				"name":        "Lilya Kio",
				"phone":       "+7(999)000-0000",
				"gender":      "FEMALE",
				"birth_date":  timeFromRFC3339("1998-11-12T23:52:14Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("089f1f77bcf86cd799439011"),
					objectIDFromHex("099f1f77bcf86cd799439011"),
				},
				"studio_id": objectIDFromHex("507f1f77bcf86cd799439013"),
			},
			{
				"_id":         objectIDFromHex("457f1f77bcf86cd799439011"),
				"name":        "Nastya Vecher",
				"phone":       "+7(998)999-9999",
				"gender":      "MALE",
				"birth_date":  timeFromRFC3339("1998-11-12T23:52:14Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("109f1f77bcf86cd799439011"),
					objectIDFromHex("119f1f77bcf86cd799439011"),
				},
				"studio_id": objectIDFromHex("507f1f77bcf86cd799439013"),
			},
		},
		clients: {
			{
				"_id":         objectIDFromHex("307f1f77bcf86cd799439011"),
				"name":        "Elizaveta Andreeva",
				"phone":       "+7(999)999-9999",
				"gender":      "FEMALE",
				"birth_date":  timeFromRFC3339("2001-10-28T23:58:18Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"password":    "1",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("009f1f77bcf86cd799439011"),
					objectIDFromHex("019f1f77bcf86cd799439011"),
					objectIDFromHex("109f1f77bcf86cd799439011"),
					objectIDFromHex("119f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":         objectIDFromHex("317f1f77bcf86cd799439011"),
				"name":        "Egor Butylo",
				"phone":       "+7(999)888-8888",
				"gender":      "MALE",
				"birth_date":  timeFromRFC3339("2002-09-11T23:52:14Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"password":    "1",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("009f1f77bcf86cd799439011"),
					objectIDFromHex("019f1f77bcf86cd799439011"),
					objectIDFromHex("029f1f77bcf86cd799439011"),
					objectIDFromHex("039f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":         objectIDFromHex("327f1f77bcf86cd799439011"),
				"name":        "Oleg Mongol",
				"phone":       "+7(999)777-7777",
				"gender":      "MALE",
				"birth_date":  timeFromRFC3339("2012-04-22T23:52:14Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"password":    "1",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("029f1f77bcf86cd799439011"),
					objectIDFromHex("039f1f77bcf86cd799439011"),
					objectIDFromHex("049f1f77bcf86cd799439011"),
					objectIDFromHex("059f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":         objectIDFromHex("337f1f77bcf86cd799439011"),
				"name":        "Irina Chikipiki",
				"phone":       "+7(999)666-6666",
				"gender":      "FEMALE",
				"birth_date":  timeFromRFC3339("2008-03-03T23:52:14Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"password":    "1",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("049f1f77bcf86cd799439011"),
					objectIDFromHex("059f1f77bcf86cd799439011"),
					objectIDFromHex("069f1f77bcf86cd799439011"),
					objectIDFromHex("079f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":         objectIDFromHex("347f1f77bcf86cd799439011"),
				"name":        "Vladislav Frolov",
				"phone":       "+7(999)555-5555",
				"gender":      "MALE",
				"birth_date":  timeFromRFC3339("1998-11-12T23:52:14Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"password":    "1",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("069f1f77bcf86cd799439011"),
					objectIDFromHex("079f1f77bcf86cd799439011"),
					objectIDFromHex("089f1f77bcf86cd799439011"),
					objectIDFromHex("099f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":         objectIDFromHex("357f1f77bcf86cd799439011"),
				"name":        "Vova Shustiy",
				"phone":       "+7(998)888-8888",
				"gender":      "MALE",
				"birth_date":  timeFromRFC3339("2011-01-11T23:52:14Z"),
				"created_at":  time.Now(),
				"updated_at":  time.Now(),
				"picture_uri": "cdn.example.com",
				"password":    "1",
				"class_ids": []bson.ObjectID{
					objectIDFromHex("089f1f77bcf86cd799439011"),
					objectIDFromHex("099f1f77bcf86cd799439011"),
					objectIDFromHex("109f1f77bcf86cd799439011"),
					objectIDFromHex("119f1f77bcf86cd799439011"),
				},
			},
		},
		classes: {
			{
				"_id":        objectIDFromHex("009f1f77bcf86cd799439011"),
				"name":       "Vodnaya yoga",
				"time":       timeFromRFC3339("2025-01-09T13:00:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("407f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439011"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("307f1f77bcf86cd799439011"),
					objectIDFromHex("317f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("019f1f77bcf86cd799439011"),
				"name":       "Goryachaya yoga",
				"time":       timeFromRFC3339("2025-01-09T13:30:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("407f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439011"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("307f1f77bcf86cd799439011"),
					objectIDFromHex("317f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("029f1f77bcf86cd799439011"),
				"name":       "Holodnaya yoga",
				"time":       timeFromRFC3339("2025-01-09T14:00:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("417f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439011"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("317f1f77bcf86cd799439011"),
					objectIDFromHex("327f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("039f1f77bcf86cd799439011"),
				"name":       "Coal yoga",
				"time":       timeFromRFC3339("2025-01-09T14:30:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("417f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439011"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("317f1f77bcf86cd799439011"),
					objectIDFromHex("327f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("049f1f77bcf86cd799439011"),
				"name":       "Vodnaya yoga",
				"time":       timeFromRFC3339("2025-01-09T13:00:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("427f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439012"),
				"class_ids": []bson.ObjectID{
					objectIDFromHex("327f1f77bcf86cd799439011"),
					objectIDFromHex("337f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("059f1f77bcf86cd799439011"),
				"name":       "Goryachaya yoga",
				"time":       timeFromRFC3339("2025-01-09T13:30:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("427f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439012"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("327f1f77bcf86cd799439011"),
					objectIDFromHex("337f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("069f1f77bcf86cd799439011"),
				"name":       "Holodnaya yoga",
				"time":       timeFromRFC3339("2025-01-09T14:00:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("437f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439012"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("337f1f77bcf86cd799439011"),
					objectIDFromHex("347f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("079f1f77bcf86cd799439011"),
				"name":       "Coal yoga",
				"time":       timeFromRFC3339("2025-01-09T14:30:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("437f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439012"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("337f1f77bcf86cd799439011"),
					objectIDFromHex("347f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("089f1f77bcf86cd799439011"),
				"name":       "Vodnaya yoga",
				"time":       timeFromRFC3339("2025-01-09T13:00:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("447f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439013"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("347f1f77bcf86cd799439011"),
					objectIDFromHex("357f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("099f1f77bcf86cd799439011"),
				"name":       "Goryachaya yoga",
				"time":       timeFromRFC3339("2025-01-09T13:30:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("447f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439013"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("347f1f77bcf86cd799439011"),
					objectIDFromHex("357f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("109f1f77bcf86cd799439011"),
				"name":       "Holodnaya yoga",
				"time":       timeFromRFC3339("2025-01-09T14:00:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("457f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439013"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("357f1f77bcf86cd799439011"),
					objectIDFromHex("307f1f77bcf86cd799439011"),
				},
			},
			{
				"_id":        objectIDFromHex("119f1f77bcf86cd799439011"),
				"name":       "Coal yoga",
				"time":       timeFromRFC3339("2025-01-09T14:30:00Z"),
				"created_at": time.Now(),
				"updated_at": time.Now(),
				"trainer_id": objectIDFromHex("457f1f77bcf86cd799439011"),
				"studio_id":  objectIDFromHex("507f1f77bcf86cd799439013"),
				"client_ids": []bson.ObjectID{
					objectIDFromHex("357f1f77bcf86cd799439011"),
					objectIDFromHex("307f1f77bcf86cd799439011"),
				},
			},
		},
	}); err != nil {
		panic(fmt.Sprintf("Failed to import DB: %v", err))
	}
}
