package main

import "fmt"

type T struct {
	f string
}

func (t T) sprint() {
	fmt.Println(t.f)
}

func (t *T) pprint() {
	if t == nil {
		fmt.Println("nil...")
		return
	}
	fmt.Println(t.f)
}

func main() {
	var t *T
	t.pprint()
	t.sprint()
}
