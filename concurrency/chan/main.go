package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	close(ch)

	select {
	case ch <- 1:
		fmt.Printf("could send")
	case <-time.After(1 * time.Millisecond):
		fmt.Printf("couldn't send")
	}
}
