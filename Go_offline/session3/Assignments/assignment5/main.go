// Take a float input from the user, convert it to an integer, and determine whether the resulting integer is even or odd.

package main

import (
	"fmt"
)

func main() {
	var f float64
	fmt.Println("Enter a float number: ")
	fmt.Scanln(&f)
	i := int(f)
	if i%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}