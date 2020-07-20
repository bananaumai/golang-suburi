package main

import "fmt"

func main() {
	count()
}

func count() {
	i := 0
	defer func() {
		fmt.Println(i)
	}()

	for ; i < 10; i++ {
		i++
	}
}
