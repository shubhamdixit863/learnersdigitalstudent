// Write a program that takes any variable and prints its type using %T in fmt.Printf.

package main

import "fmt"

func main() {
	a:=0
	fmt.Scanln(&a)
	fmt.Printf("The type of the variable is %T\n", a)
}

