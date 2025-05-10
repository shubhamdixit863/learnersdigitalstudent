package main

import (
	"fmt"
	"os"
)

func main() {
	//1. you have to write a function called  fizzbuzz that takes an
	//argument if it is divisible by 3 return fizz if by 5 return buzz if by both return fizzbuzz
	fmt.Println(fizzbuzz(3))
	fmt.Println(fizzbuzz(5))
	fmt.Println(fizzbuzz(15))

	//2. create a backend application that takes following things:
	//first name
	//second age
	//third balance
	fmt.Println("Name:", os.Args[1])
	fmt.Println("Age:", os.Args[2])
	fmt.Println("Balance:", os.Args[3])

	//3. you have to write a program that takes user details as input: name, age, address
	//convert the age into days.
	//print out the whole details

	var age int
	var name string
	var address string
	var company string
	var parent string

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &age)
	fmt.Print("Enter your address: ")
	fmt.Scanf("%s", &address)
	fmt.Print("Enter your company name: ")
	fmt.Scanf("%s", &company)
	fmt.Print("Enter your Parent's name: ")
	fmt.Scanf("%s", &parent)
	fmt.Println("You are", name, "son/daughter of", parent, "from", address, ", you are", age*365, "days old", "and you work at", company, ".")
}

func fizzbuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return "None"
	}
}
