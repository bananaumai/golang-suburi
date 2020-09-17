package main

import (
	"bytes"
	"fmt"

	"github.com/vmihailenco/msgpack/v4"
)

type S struct {
	N int8        `msgpack:"n"`
	I interface{} `msgpack:"i"`
}

func main() {
	s := S{1, "abc"}
	buf := bytes.NewBuffer([]byte{})
	enc := msgpack.NewEncoder(buf)
	_ = enc.Encode(s)

	bs := buf.Bytes()
	for i, b := range bs {
		fmt.Printf("%X", b)
		if i != len(bs)-1 {
			fmt.Print(" ")
		} else {
			fmt.Println("")
		}
	}
}
