package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/vmihailenco/msgpack/v4"
)

type HLS struct {
	SearchKey string        `msgpack:"searchKey"`
	MetaData  interface{}   `msgpack:"metaData,omitempty"`
	M3U8      string        `msgpack:"m3u8"`
	Files     []SegmentFile `msgpack:"files"`
}

type SegmentFile struct {
	Timestamp time.Time `msgpack:"timestamp"`
	Name      string    `msgpack:"name"`
	Blob      []byte    `msgpack:"blog"`
}

func main() {
	hls := HLS{
		SearchKey: "/banana/umai",
		MetaData:  nil,
		M3U8:      "playlist....",
		Files: []SegmentFile{
			{
				Timestamp: time.Now(),
				Name:      "s1",
				Blob:      []byte{1, 2, 3},
			},
			{
				Timestamp: time.Now(),
				Name:      "s2",
				Blob:      []byte{2, 3, 4},
			},
		},
	}

	buf := bytes.Buffer{}
	enc := msgpack.NewEncoder(&buf)
	_ = enc.Encode(hls)

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
