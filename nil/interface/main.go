package main

import "fmt"

type (
	i interface {
		foo()
	}
	s struct{}
)

func (s *s) foo() {}

func main() {
	vI := cI()
	fmt.Printf("vI(cI) : %v\n", vI)
	if vI != nil {
		fmt.Printf("not nil !? vI(cI) : %v\n", vI)
	} else {
		fmt.Printf("yes nil -- vI(cI) : %v\n", vI)
	}

	vI = cS()
	fmt.Printf("vI(cS) : %+v\n", vI)
	if vI != nil {
		fmt.Printf("not nil !? vI(cS) : %v\n", vI)
	} else {
		fmt.Printf("yes nil -- vI(cS) : %v\n", vI)
	}

	vS := cS()
	fmt.Printf("vS : %v\n", vS)
	if vS != nil {
		fmt.Printf("not nil !? vS : %v\n", vS)
	} else {
		fmt.Printf("yes nil -- vS : %v\n", vS)
	}
}

func cS() *s {
	return nil
}

func cI() i {
	return nil
}
