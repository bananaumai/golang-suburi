package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	//close(ch)

	go func() {
		select {
		case ch <- 1:
			log.Printf("could send")
		case <-time.After(1 * time.Millisecond):
			log.Printf("couldn't send")
		}
	}()

	n := <-ch
	log.Printf("received: %d", n)
}
