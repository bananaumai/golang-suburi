package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	src := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	r := bytes.NewReader(src)
	buf := make([]byte, 3)
	for {
		size, err := r.Read(buf)
		if err != nil && err == io.EOF {
			break
		}
		for _, b := range buf[:size] {
			fmt.Printf("%X ", b)
		}
	}
	fmt.Println("done")
}
