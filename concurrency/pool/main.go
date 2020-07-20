package main

import "sync"

var (
	bytesPool = sync.Pool{
		New: func() interface{} {
			return []byte{}
		},
	}
)

func main() {

}
