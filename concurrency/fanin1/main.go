package main

import (
	"fmt"
	"sync"
	"time"
)

func ns(done chan interface{}, n int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case ch <- n:
			}
		}
	}()
	return ch
}

func fanIN(done chan interface{}, chs ...chan int) chan int {
	wg := &sync.WaitGroup{}
	ch := make(chan int, len(chs))
	for _, c := range chs {
		wg.Add(1)
		c := c
		go func() {
			defer wg.Done()
			for n := range c {
				select {
				case <-done:
					return
				case ch <- n:
				}
			}
		}()
	}
	go func() {
		defer close(ch)
		wg.Wait()
	}()
	return ch
}

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(1 * time.Millisecond)
		close(done)
	}()
	ones := ns(done, 1)
	twos := ns(done, 2)
	var cnt1, cnt2 int
	for n := range fanIN(done, ones, twos) {
		if n == 1 {
			cnt1++
		} else {
			cnt2++
		}
	}
	fmt.Printf("1: %d, 2: %d", cnt1, cnt2)
}
