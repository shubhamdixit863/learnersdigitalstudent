// Accept three integers and determine the largest using if-else conditions.

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	if a > b && a > c {
		fmt.Printf("%d is the largest number\n", a)
	} else if b > a && b > c {
		fmt.Printf("%d is the largest number\n", b)
	} else {
		fmt.Printf("%d is the largest number\n", c)
	}
}