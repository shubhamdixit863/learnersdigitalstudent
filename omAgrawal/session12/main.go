package main

import "fmt"

func main() {
	go foo()
	fmt.Println("fff")
	fmt.Println("fff")
	fmt.Println("fff")

}

func foo() {
	fmt.Println("fff")
}
