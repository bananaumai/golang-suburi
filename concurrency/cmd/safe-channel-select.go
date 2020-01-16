package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := produce(1, 5)
	c2 := produce(2, 8)
	consumeTwoChanUntilBothAreClosed(c1, c2)
}

func produce(base int, limit int) chan int {
	c := make(chan int)

	go func() {
		for i := 1; i < limit; i++ {
			time.Sleep(5 * time.Millisecond)
			c <- base * i
		}
		close(c)
	}()

	return c
}

func consumeTwoChanUntilBothAreClosed(c1, c2 chan int) {
	for {
		select {
		case v, ok := <-c1:
			if !ok {
				c1 = nil
			} else {
				fmt.Printf("c1: %d\n", v)
			}
		case v, ok := <-c2:
			if !ok {
				c2 = nil
			} else {
				fmt.Printf("c2: %d\n", v)
			}
		default:
			if c1 == nil {
				fmt.Println("c1 is nil")
			}
			if c2 == nil {
				fmt.Println("c2 is nil")
			}
			if c1 == nil && c2 == nil {
				return
			}
		}
	}
}
