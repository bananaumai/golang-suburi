package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	runUntilCanceled(ctx)
}

func runUntilCanceled(ctx context.Context) {
	defer finalize(ctx)

	for {
		select {
		case <- ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Println("bnn")
			time.Sleep(time.Second)
		}
	}
}

func finalize(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	done := make(chan bool)
	go func() {
		heavy()
		done <- true
	}()

	select {
	case <- ctx.Done():
		fmt.Println("finalization failed")
	case <-done:
		fmt.Println("finalization succeeded")
	}
}

func heavy() {
	time.Sleep(2*time.Second)
}
