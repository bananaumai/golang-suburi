package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	apiAddr := os.Args[1]

	readCloser, err := fetchDataFromApi(context.Background(), apiAddr)
	defer func() { _ = readCloser.Close() }()
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(readCloser)
	if err != nil {
		panic(err)
	}

	fmt.Printf("data size: %d", len(data))
}

func fetchDataFromApi(ctx context.Context, apiAddr string) (io.ReadCloser, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiAddr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send http request: %w", err)
	}

	return resp.Body, nil
}
