package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	const timeout = 10 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	g.Go(func() error {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			time.Sleep(100 * time.Millisecond)
			log.Printf("slower : %d", i)
		}
		if r.Int()%2 == 0 {
			return errors.New("slower error")
		} else {
			return nil
		}
	})

	g.Go(func() error {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			time.Sleep(50 * time.Millisecond)
			log.Printf("faster : %d", i)
		}
		if r.Int()%2 == 0 {
			return errors.New("faster error")
		} else {
			return nil
		}
	})

	if err := g.Wait(); err != nil {
		log.Printf("err: %s", err)
	} else {
		log.Printf("succeeds!")
	}
}
