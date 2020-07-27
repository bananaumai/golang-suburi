package main

import (
	"errors"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	e1 := errors.New("test")
	e2 := fmt.Errorf("test")

	if cmp.Equal(e1, e2) {
		fmt.Printf("same")
	} else {
		fmt.Printf("not same")
	}
}
