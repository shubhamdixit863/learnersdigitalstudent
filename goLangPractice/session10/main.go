package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type myCustomeError struct {
}

func (c *myCustomeError) Error() string {
	return "CustomeError"
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func ReturnCustomError() error {
	return &myCustomeError{}
}

func ReadFile(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("Error in Reading File")
	}

	log.Println(file)
}

func Recover() {
	r := recover()
	if r != nil {
		log.Println("Recovered from panic: ", r)
	}

	fmt.Println("Recover Completed")
}

var err = errors.New("SAME ERROR")

func ReturnError() error {
	return fmt.Errorf("%v", err)
}

func ReturnError2() error {
	return io.EOF
}

func main() {
	// b, err := Divide(6, 2)

	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println(b)
	// defer Recover()

	// ReadFile("file.txt")

	e1 := ReturnError2()

	if errors.Is(e1, io.EOF) {
		log.Println("They are same error")
	}

}
