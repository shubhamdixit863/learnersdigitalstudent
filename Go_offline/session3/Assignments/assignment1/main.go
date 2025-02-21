// Convert String Inputs to Integers and Calculate Sum:

package main

import (
	"fmt"
	"strconv" // for string to int conversion
)

func main() {
	var num1, num2 string
	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)
	a, err:=strconv.Atoi(num1)
	if err != nil {
		fmt.Println("Conversion not possible", err)
	}


	b, err:=strconv.Atoi(num2)
	if err != nil {
		fmt.Println("Conversion not possible", err)
	}

	c := a + b
	fmt.Println("Sum of two numbers is: ", c)
}
