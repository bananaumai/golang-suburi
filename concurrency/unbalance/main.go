package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	senders := 10
	processors := 30

	incomingCh := make(chan string)
	dispatcherCh := make(chan string)

	for i := 0; i < senders; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				var msg string

				select {
				case <-ctx.Done():
					log.Printf("sender finished")
					return
				case msg = <-incomingCh:
					log.Printf("received %s", msg)
				}

				for i := 0; i < 2; i++ {
					msg := fmt.Sprintf("%d: %s", i, msg)
					select {
					case dispatcherCh <- msg:
						log.Printf("dispatched %s", msg)
					default:
						log.Printf("[ERROR] couldn't dispatch since processors are busy")
					}
				}
			}
		}()
	}

	for i := 0; i < processors; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var msg string
			select {
			case <-ctx.Done():
				log.Printf("processor finished")
				return
			case msg = <-dispatcherCh:
				log.Printf("processing %s", msg)
			}

			time.Sleep(1 * time.Second) // emulate heavy process

			log.Printf("processed %s", msg)
		}()
	}

	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("message %d", i)
		incomingCh <- msg
	}

	cancel()

	wg.Wait()

	log.Printf("done")
}
