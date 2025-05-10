package main

import "fmt"

func main() {
	var name string
	var age int
	var address string
	
	fmt.Println("Enter your name:")
	fmt.Scan(&name)

	fmt.Println("Enter your age:")
	fmt.Scan(&age)
	
	fmt.Println()
	var age_in_days int = age * 365
	
	fmt.Print("Enter your address:")
	fmt.Scan(&address)
	
	fmt.Println()
	fmt.Println("**Preson Details**")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Age in days: " , age_in_days)
	fmt.Println("Address:", address)

}