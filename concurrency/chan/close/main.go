package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 100)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 20; i++ {
			ch <- i
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
}
