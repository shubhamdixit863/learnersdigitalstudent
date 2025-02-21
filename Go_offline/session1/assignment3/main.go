package main

import("fmt")

func main(){
	fmt.Println("Enter Name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter Age:")
	var age int
	fmt.Scanln(&age)

	fmt.Println("Enter City:")
	var city string
	fmt.Scanln(&city)

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age*365, "days")
	fmt.Println("City: ", city)
}