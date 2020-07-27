package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := regexp.MustCompile(`(\d+)\.packet$`)
	key := "foo/4.packet"
	match := reg.FindSubmatch([]byte(key))
	fmt.Printf("%q\n", match)
}
