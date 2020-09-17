package main

import (
	"encoding/json"
	"log"
)

type (
	T1 struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	T2 struct {
		T1
		C int `json:"c"`
	}
)

func main() {
	t2 := T2{
		T1: T1{A: 1, B: 2},
		C:  3,
	}
	bs, _ := json.Marshal(t2)
	log.Printf("%s", bs)
}
