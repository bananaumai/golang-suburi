package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/heavy", func(writer http.ResponseWriter, _ *http.Request) {
		huge := make([]byte, 1024*1024*256)
		_, _ = writer.Write(huge)
	})
	mux.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
		_, _ = writer.Write([]byte("hello"))
	})

	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		panic(err)
	}
}
