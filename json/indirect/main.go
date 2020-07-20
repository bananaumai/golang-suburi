package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type (
	person struct {
		Name string
		Age  uint
	}
)

func main() {
	jsn := `{"name": "banana", "age": 17}`
	var p person
	if err := load(strings.NewReader(jsn), &p); err != nil {
		panic(err)
	}
	fmt.Printf("p: %v\n", p)
}

func load(source io.Reader, obj interface{}) error {
	dec := json.NewDecoder(source)
	return dec.Decode(obj)
}
