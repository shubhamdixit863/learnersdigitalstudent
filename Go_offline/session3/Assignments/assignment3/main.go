// Take three integer inputs from the user, convert them to floats, and calculate the average using a function./

package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)
	fmt.Println("Enter third number: ")
	fmt.Scanln(&num3)
	a:= float64(num1)
	b:= float64(num2)
	c:= float64(num3)

	average := calculateAverage(a, b, c)
	fmt.Printf("Average of three numbers is: %.2f\n", average)


}

func calculateAverage(a, b, c float64) float64 {
	return (a + b + c) / 3
}

