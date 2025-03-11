package main

import "fmt"

func main() {

	go foo()
}
func foo() {
	fmt.Println("Hello world")

}
