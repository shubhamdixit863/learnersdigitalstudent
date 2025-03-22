package main

import (
	"errors"
	"fmt"
)

type CustomeError struct {
	mess string
}

func (c *CustomeError) Error() string {
	return " CustomeError  ekmk e3d"
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can not divide by 0")

	}
	return a / b, nil
}
func ReturnCustome() error {
	return &CustomeError{}
}
func main() {

	customErr := ReturnCustome()
	fmt.Println("Error from ReturnCustome:", customErr)
}
