package main

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collName       = "bnn"
	dbName         = "suburi"
	url            = "mongodb://localhost:27017/"
	concurrencyNum = 100
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalf("Failed to connect to %s: %v", url, err)
	}
	db := client.Database(dbName)
	coll := db.Collection(collName)

	if err := coll.Drop(ctx); err != nil {
		log.Fatalf("failed to drop: %s", err)
	}
	//if _, err := coll.Indexes().CreateOne(ctx, mongo.IndexModel{
	//	Keys:    bson.M{"uid": 1},
	//	Options: options.Index().SetUnique(true),
	//}); err != nil {
	//	log.Fatalf("failed to create index: %s", err)
	//}

	test(ctx, coll, concurrencyNum)
}

func test(ctx context.Context, coll *mongo.Collection, concurrency int) {
	for i := int64(0); i < 10; i++ {
		i := i
		time.Sleep(100 * time.Millisecond)
		func() {
			wg := &sync.WaitGroup{}
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			for j := 0; j < concurrency; j++ {
				j := j
				wg.Add(1)
				go func() {
					defer wg.Done()
					err := update(ctx, coll, i, j)
					if err != nil {
						log.Fatalf("error: %s", err)
					}
				}()
			}
			wg.Wait()
		}()
	}
}

func update(ctx context.Context, coll *mongo.Collection, uid int64, sweet int) error {
	filter := bson.M{"uid": uid}
	update := bson.M{"$set": bson.M{"sweet": sweet}}

	opts := options.Update()
	opts.SetUpsert(true)

	_, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
