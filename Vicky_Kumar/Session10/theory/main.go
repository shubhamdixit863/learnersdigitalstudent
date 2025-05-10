package main

import (
	"errors"
	"log"
)

type CustomError struct {
	message string
}

func (c *CustomError) Error() string {
	return "CustomError"
}
func Divide(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("Cannot divide by zero")
		return 0, &CustomError{err.Error()}
	}
	return a / b, nil
}
func ReturnCustomError() error {
	return &CustomError{}
}

func main() {
	b, err := Divide(6, 0)
	if err != nil {
		log.Println(err)
	}
	log.Print(b)
}
