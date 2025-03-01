package main

import "fmt"

func main() {

	var choice string
	fmt.Println("Enter *:multiplication, +:addition, -:substraction and /:division")
	fmt.Scan(choice)

	var num1 int
	fmt.Println("Enter number one: ")
	fmt.Scan(num1)

	var num2 int
	fmt.Println("Enter number second: ")
	fmt.Scan(num2)

	if choice == "+" {
		var ans int
		ans = num1 + num2
		fmt.Println("Result: ", ans)
	} else if choice == "-" {
		var ans int
		ans = num1 - num2
		fmt.Println("Result: ", ans)
	} else if choice == "*" {
		var ans int
		ans = num1 * num2
		fmt.Println("Result: ", ans)
	} else if choice == "/" {
		var ans int
		ans = num1 / num2
		fmt.Println("Result: ", ans)
	}
}
