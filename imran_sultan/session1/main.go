package main

import (
	"errors"
	"fmt"
)

type MyError struct{}

func (e MyError) Error() string {
	return "error"
}

func main() {
	var err error = MyError{}

	if errors.As(err, &MyError{}) {
		fmt.Println("Matched")
	} else {
		fmt.Println("Not matched")
	}
}
