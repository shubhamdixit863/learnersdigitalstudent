package main

import ("fmt")

func main(){
	fmt.Println("Welcome to the Calculator Application")
	fmt.Println("Enter Add for Addition")
	fmt.Println("Enter Sub for Subtraction")
	fmt.Println("Enter Mult for Multiplication")
	fmt.Println("Enter Div for Division")

	var operation string
	fmt.Scanln(&operation)

	var a int
	fmt.Println("Enter first number:")
	fmt.Scanln(&a)

	var b int
	fmt.Println("Enter second number:")
	fmt.Scanln(&b)

	fmt.Println("Result: ", Calculator(operation, a, b))
}

    