// see https://golang.org/ref/spec#Variables

package main

import (
	"fmt"
	"time"
)

func reAssignSlice(ns []int) {
	ns = []int{}
}

func useSlice(ns []int) {
	go func() {
		time.Sleep(1*time.Second)
		fmt.Printf("useSlice - %v\n", ns)
	}()
}

type T struct {
	n int
}

func reAssignPointerOfT(tp *T) {
	tp = &T{}
}

func reAssignDereferencePointerOfT(tp *T) {
	*tp = T{}
}

func usePointerOfT(tp *T) {
	go func() {
		time.Sleep(1*time.Second)
		fmt.Printf("usePointerOfT - %v\n", tp)
	}()
}

func main() {
	s := []int{1,2,3}
	reAssignSlice(s)
	fmt.Printf("main s - %v\n", s)
	useSlice(s)
	s = []int{}
	time.Sleep(2*time.Second)

	t := T{1}
	reAssignPointerOfT(&t)
	fmt.Printf("main t - %v\n", t)
	reAssignDereferencePointerOfT(&t)
	fmt.Printf("main t - %v\n", t)
	usePointerOfT(&t)
	t = T{2}
	time.Sleep(2*time.Second)

	tp := &T{1}
	reAssignPointerOfT(tp)
	fmt.Printf("main tp - %v\n", tp)
	reAssignDereferencePointerOfT(tp)
	fmt.Printf("main tp - %v\n", tp)
	usePointerOfT(tp)
	tp = &T{2}
	time.Sleep(2*time.Second)
}
