package main

import "fmt"

func main() {

	var name string
	var age int
	var address string
	var company string
	var parent string
	

	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)
	fmt.Println("Enter your address:")
	fmt.Scanln(&address)
	fmt.Println("Enter your company:")
	fmt.Scanln(&company)
	fmt.Println("Enter your parent's name:")
	fmt.Scanln( &parent)

	var ageInDays int = age * 365

	fmt.Printf("Your name is %s, with age in days %d. You live in %s. Your compnay name is %s. Your parent's name is %s.", name, ageInDays, address, company, parent)
}