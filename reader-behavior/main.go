package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func main() {

	rc := ioutil.NopCloser(bytes.NewReader([]byte("abc")))

	// ensure that the reader contains the proper contents
	log.Printf("rc is something like %s", rc)

	// a slice into which to read
	buf := make([]byte, 0, 10)
	buf = append(buf, 0xab)

	// read the slice
	if _, err := rc.Read(buf); err != nil {
		log.Printf("error: %s", err)
	}

	// check to make sure you have it
	log.Printf("Read '%s'", buf)
}
