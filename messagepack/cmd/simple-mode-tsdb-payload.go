package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/vmihailenco/msgpack/v4"
)

func main() {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, "2020-01-01T00:00:00.000Z")
	if err != nil {
		os.Exit(1)
	}
	timpestamp := t.UnixNano()
	seriesIDs := []string{"series1", "series2"}
	values := []float64{1.0, 1.0}

	writer := new(bytes.Buffer)
	encoder := msgpack.NewEncoder(writer)
	encoder.EncodeString("DATA")
	encoder.EncodeString("1.0")
	encoder.EncodeUint(uint64(len(seriesIDs)))
	for _, id := range seriesIDs {
		encoder.EncodeString(id)
	}
	encoder.EncodeInt(timpestamp)
	for _, v := range values {
		encoder.EncodeFloat64(v)
	}

	bytes := writer.Bytes()
	length := len(bytes)
	for i, b := range bytes {
		fmt.Printf("%02X", b)
		if i != length-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("")
}
