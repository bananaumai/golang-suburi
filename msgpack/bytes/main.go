package main

import (
	"bytes"
	"fmt"

	"github.com/vmihailenco/msgpack/v4"
)

func main() {
	bs := []byte{1, 2, 3}
	buf := bytes.Buffer{}
	enc := msgpack.NewEncoder(&buf)
	_ = enc.Encode(bs)
	for _, b := range buf.Bytes() {
		fmt.Printf("%X ", b)
	}
}
