//you have to write a program that takes user details as input: name, age, address
//convert the age into days.
//print out the whole details

package main

import "fmt"

func main() {
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
