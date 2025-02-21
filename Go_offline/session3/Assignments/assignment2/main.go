// Create a function that converts Fahrenheit to Celsius. Take temperature input from the user and display the result with two decimal places.

package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64
	fmt.Println("Enter temperature in Fahrenheit: ")
	fmt.Scanln(&fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Temperature in Celsius: %.2f\n", celsius)
}
