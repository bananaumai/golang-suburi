package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Foo struct {
	F string
}

func main() {
	mF := make(map[string]interface{})
	mF["f"] = "bar"

	var sF Foo
	if err := mapstructure.Decode(mF, &sF); err != nil {
		panic(err)
	}

	fmt.Println(sF.F)

}
