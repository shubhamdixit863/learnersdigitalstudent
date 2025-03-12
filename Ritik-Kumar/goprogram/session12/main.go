package main

import "fmt"

func foo() {
	fmt.Println("hello")
}

func main(){
	go foo()
}