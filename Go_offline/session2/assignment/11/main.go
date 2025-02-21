// Use a for loop to calculate the sum of the first N natural numbers.

package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Printf("The sum of the first %d natural numbers is %d\n", n, sum(n))
}

func sum(n int) int {
	sum := 0
	for i:=1; i<=n; i++ {
		sum += i
	}
	return sum
}
