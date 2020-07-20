package main

import "fmt"

func main() {
	//s := make([]int, 0, 10)
	//s := make([]int, 10, 10)
	var s []int

	print(s)

	for i := 0; i < 20; i++ {
		s = append(s, i)
		print(s)
	}
}

func print(ns []int) {
	fmt.Printf("%d: %d: %v\n", len(ns), cap(ns), ns)
}
