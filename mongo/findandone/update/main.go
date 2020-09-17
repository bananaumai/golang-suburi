package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
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

type Bnn struct {
	Sweet int `bson:"sweet"`
}

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
	if _, err := coll.InsertOne(ctx, Bnn{Sweet: 1}); err != nil {
		log.Fatalf("failed to insert: %s", err)
	}

	test(ctx, coll, concurrencyNum)
}

func test(ctx context.Context, coll *mongo.Collection, concurrency int) {
	sweet := 1
	for i := int64(0); i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		func() {
			wg := &sync.WaitGroup{}
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			bnns := make(chan Bnn, concurrency)
			for j := 0; j < concurrency; j++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					bnn, err := findOneAndUpdate(ctx, coll, sweet, rand.Int())
					if err != nil {
						if errors.Is(err, mongo.ErrNoDocuments) {
							return
						}
						log.Fatalf("error: %s", err)
					}
					log.Printf("#%d : %d: bnn: %+v", i, sweet, bnn)
					bnns <- bnn
				}()
			}

			var updated bool
			var nextSweet int
			go func() {
				for bnn := range bnns {
					if !updated {
						updated = true
						nextSweet = bnn.Sweet
					} else {
						log.Fatalf("unexpected bnn: %+v", bnn)
					}
				}
			}()

			wg.Wait()
			close(bnns)
			if !updated {
				log.Fatalf("not updated: %d", i)
			}
			sweet = nextSweet
		}()
	}
}

func findOneAndUpdate(ctx context.Context, coll *mongo.Collection, sweet int, newSweet int) (Bnn, error) {
	filter := bson.M{"sweet": sweet}
	update := bson.M{"$set": bson.M{"sweet": newSweet}}

	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)

	var bnn Bnn
	if err := coll.FindOneAndUpdate(ctx, filter, update, opts).Decode(&bnn); err != nil {
		return bnn, err
	}
	return bnn, nil
}
