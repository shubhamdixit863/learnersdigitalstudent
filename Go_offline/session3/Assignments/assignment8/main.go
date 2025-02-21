// Take a string input representing a floating-point number, convert it to a float, and print the square of the number. Handle any conversion errors.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f string
	fmt.Println("Enter a float number: ")
	fmt.Scanln(&f)
	i, err := strconv.ParseFloat(f, 64)
	if err != nil {
		fmt.Println("Conversion not possible", err)
	}
	fmt.Println(i * i)
}