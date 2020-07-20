package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		a int64
		b int64
		x float64
		y float64
		foo string
	)

	flag.Int64Var(&a, "a", 0, "a")
	flag.Int64Var(&b, "b", 0, "b")
	flag.Float64Var(&x, "x", 0.0, "x")
	flag.Float64Var(&y, "y", 0.0, "y")
	flag.StringVar(&foo, "foo", "foo", "foo")
	flag.Parse()
	args := flag.Args()

	fmt.Printf("a: %d\n", a)
	fmt.Printf("b: %d\n", b)
	fmt.Printf("x: %f\n", x)
	fmt.Printf("y: %f\n", y)
	fmt.Printf("foo: %s\n", foo)
	fmt.Printf("args: %v\n", args)
}
