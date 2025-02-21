// Write a function that converts a temperature from Celsius to Fahrenheit. The formula is F = (C * 9/5) + 32.

package main

import "fmt"

func main() {
	var c float64
	fmt.Scanln(&c)
	fmt.Printf("Temperature in Fahrenheit is %f\n", temp(c))
}


func temp(celsius float64) float64 {
	return (celsius * 9/5) + 32
}