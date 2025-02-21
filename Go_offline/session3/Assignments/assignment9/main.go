// Read an integer, a float, and a string from the user. Convert the integer to a float and the float to an integer, then display the converted values.

package main

import (

	"fmt"
)

func main() {
	var i int
	var f float64
	var s string
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&i)
	fmt.Println("Enter a float: ")
	fmt.Scanln(&f)
	fmt.Println("Enter a string: ")
	fmt.Scanln(&s)
	fmt.Printf("Integer to float: %.2f\n", float64(i))
	fmt.Printf("Float to integer: %d\n", int(f))
	fmt.Printf("String: %s\n", s)
}