// Write a function that takes two integers and swaps their values without using a third variable.

package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Println("Before swapping: ", a, b)
	a=a+b
	b=a-b
	a=a-b
	fmt.Println("After swapping: ", a, b)
   
}