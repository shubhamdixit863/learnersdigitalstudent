// Print the first N terms of the Fibonacci series using a for loop.

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	printFibonacciSeries(n)
}

func printFibonacciSeries(n int) {
	a, b := 0, 1
	for i:=0; i<n; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}