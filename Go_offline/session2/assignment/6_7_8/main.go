// Write a function that takes two integers and returns their sum.

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Printf("The sum of the two integers is %d\n", sum(a, b))

	fmt.Println(sumVariadic(1, 2, 3, 4, 5))
	fmt.Println(max(1, 2, 3, 4, 5))
}

func sum(a, b int) int {
	return a + b
}

// Create a variadic function that accepts any number of integers and returns their sum.
func sumVariadic(nums ...int) int {
	sum := 0
	for i:=0; i<len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

// Write a variadic function that returns the maximum value among the provided integers.
func max(nums ...int) int {
	max := nums[0]
	for i:=1; i<len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}


