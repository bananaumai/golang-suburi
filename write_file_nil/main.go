package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file := os.Args[1]

	if err := ioutil.WriteFile(file, nil,  os.FileMode(0777)); err != nil {
		log.Fatalf("failed to write: %v", err)
	}

	err, data := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to read: %v", err)
	}

	fmt.Printf("data: %s\n", data)
}
