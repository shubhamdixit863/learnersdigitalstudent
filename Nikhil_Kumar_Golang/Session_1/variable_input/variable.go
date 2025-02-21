package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name and age: ")
	fmt.Scanf("%d %s", &name, &age)
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Scan(&name, &age)	
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}