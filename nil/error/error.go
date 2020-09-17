package main

import "fmt"

type MyErr struct{}

func (e MyErr) Error() string {
	return "my-error"
}

func myerr() error {
	var err *MyErr
	return err
}

func err() error {
	return nil
}

const (
	c = 200
)

func main() {
	fmt.Println(err())
	fmt.Println(myerr())

	//printInt(c)
	printUint64(c)
}

func printInt(i int) {

}

func printUint64(i uint64) {

}
