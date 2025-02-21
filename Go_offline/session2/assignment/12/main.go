// Write a program that prints the multiplication table of a given number.

package main

import "fmt"


func main() {
	var n int
	fmt.Scanln(&n)
	printMultiplicationTable(n)
}

func printMultiplicationTable(n int) {
	for i:=1; i<=10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}