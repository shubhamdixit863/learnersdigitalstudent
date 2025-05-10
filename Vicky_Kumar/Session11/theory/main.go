package main

import "fmt"

func foo() {
	fmt.Println("Hello")
}
func main() {
	go foo()
}
