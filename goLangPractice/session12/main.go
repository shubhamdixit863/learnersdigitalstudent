package main

import "fmt"

func foo() {
	fmt.Println("Hello")
	fmt.Println("World")
}

func main() {
	go foo()
}
