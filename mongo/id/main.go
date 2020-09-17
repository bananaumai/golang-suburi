package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	BookWrite struct {
		Name string `bson:"name"`
	}

	BookRead struct {
		ID        string `bson:"_id"`
		BookWrite `bson:",inline"`
	}
)

func main() {
	ctx := context.Background()

	db, err := PrepareMongoDB(ctx, "suburi", "mongodb://127.0.0.1:27017", 30*time.Second)
	if err != nil {
		log.Fatalf("failed to prepare db: %s", err)
	}

	b := BookWrite{Name: "banana"}

	_, err = db.Collection("books").InsertOne(ctx, b)
	if err != nil {
		log.Fatalf("failed to insert book: %s", err)
	}

	oid, _ := primitive.ObjectIDFromHex("5f48a3e77067e7f414836e4f")
	q := bson.M{
		"_id": bson.M{
			"$gt": oid,
		},
		"name": bson.M{
			"$regex": "^banana",
		},
	}
	cur, err := db.Collection("books").Find(ctx, q)
	if err != nil {
		log.Fatalf("failed to find: %s", err)
	}

	var bs []BookRead
	if err := cur.All(ctx, &bs); err != nil {
		log.Fatalf("failed to iterate cursor: %s", err)
	}

	log.Printf("books: %+v", bs)
}

func PrepareMongoDB(ctx context.Context, dbname, uri string, connectTimeout time.Duration) (*mongo.Database, error) {
	uriOpt := options.Client().ApplyURI(uri)
	timeoutOpt := options.Client().SetConnectTimeout(connectTimeout)

	if err := uriOpt.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate uri option: %w", err)
	}
	if err := timeoutOpt.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate timeout option: %w", err)
	}

	client, err := mongo.Connect(ctx, uriOpt, timeoutOpt)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("unable to connect to mongo daemon at specified address with specified credentials: %s", err)
	}

	return client.Database(dbname), nil
}
