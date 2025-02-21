// Accept two strings from the user and concatenate them using a function.

package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Printf("The concatenated string is %s\n", concat(a, b))
}

func concat(a, b string) string {
	return a + b
}
