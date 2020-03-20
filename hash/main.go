package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	input1 := "foo"
	hash := sha256.New()
	hash.Write([]byte(input1))
	for _, b := range hash.Sum(nil) {
		fmt.Printf("%2x ", b)
	}
	fmt.Printf("(%d)\n", hash.Size())
}
