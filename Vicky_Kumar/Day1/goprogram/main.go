package main

import (
	"fmt"
)

func main() {
	// var name string
	// var age int
	// address := ""
	// fmt.Println("Enter your name")
	// fmt.Scanln(&name)
	// fmt.Println("Enter your age: ")
	// fmt.Scanln(&age)
	// fmt.Println("Enter your address: ")
	// fmt.Scanln(&address)
	// ageInDays := age * 365
	// fmt.Println("Your name is", name, "and your age in days is", ageInDays, "and you address is", address)

	var operation string
	fmt.Println("Enter the operation you want to perform")
	fmt.Scanln(&operation)
	var num1 float32
	var num2 float32
	fmt.Println("Enter first number")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number")
	fmt.Scanln(&num2)
	var result float32
	if operation == "+" {
		result = num1 + num2
	} else if operation == "-" {
		result = num1 - num2
	} else if operation == "*" {
		result = num1 * num2
	} else if operation == "/" {
		result = num1 / num2
	} else {
		fmt.Println("Invalid operation")
	}
	fmt.Println(num1, operation, num2, "=", result)
}
