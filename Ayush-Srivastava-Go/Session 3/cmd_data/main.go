package main

import "fmt"

func main() {
	var d int8
	fmt.Scanf("%d", &d)
	fmt.Println(d)

	// Multiple variable declaration
	var(
		a = 5
		b = 5
	)

	fmt.Println(a+b)

}