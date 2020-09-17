package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type T struct {
	S string `json:"s"`
	I int    `json:"i"`
}

func (t T) String() string {
	return fmt.Sprintf("{I: %d, S: %d}", t.I, len(t.S))
}

type MyWriter struct {
	io.Writer
}

func (w *MyWriter) Write(p []byte) (int, error) {
	size := len(p)
	fmt.Printf("received %d bytes: ", size)
	//for _, b := range p {
	//	fmt.Printf("%02X ", b)
	//}
	fmt.Printf("\n")
	return w.Writer.Write(p)
}

type MyReader struct {
	io.Reader
	buf []byte
}

func (r *MyReader) Read(p []byte) (int, error) {
	panic("not implemented")
}

func main() {
	var err error

	builder := strings.Builder{}
	for i := 0; i < 1024*1024*512; i++ {
		builder.WriteString("a")
	}
	jsonBS := []byte(fmt.Sprintf(`{"s":"%s", "i": 10}`, builder.String()))
	r := bytes.NewReader(jsonBS)
	dec := json.NewDecoder(r)
	var t T
	err = dec.Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", t)

	buf := bytes.Buffer{}
	w := &MyWriter{&buf}
	enc := json.NewEncoder(w)
	err = enc.Encode(t)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", buf.String())
}
