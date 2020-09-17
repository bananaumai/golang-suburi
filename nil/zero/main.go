package main

import "fmt"

func main() {
	var sl []int = nil
	fmt.Printf("%#v\n", sl)

	var mp map[string]interface{} = nil
	fmt.Printf("%#v\n", mp)

	var ch chan struct{} = nil
	fmt.Printf("%#v\n", ch)

	var fn func() = nil
	fmt.Printf("%#v\n", fn)

	type st struct{}
	var stp *st
	fmt.Printf("%#v\n", stp)

	var err error = nil
	fmt.Printf("%#v\n", err)
}
