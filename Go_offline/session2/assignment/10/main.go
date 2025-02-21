// Implement a recursive function to find the factorial of a given number.

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Printf("The factorial of %d is %d\n", n, factorial(n))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

