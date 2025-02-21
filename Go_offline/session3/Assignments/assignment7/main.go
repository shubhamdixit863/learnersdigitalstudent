// Given a slice of string numbers ([]string{"10", "20", "30", "40"}), write a program to convert them into integers and calculate their total sum.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := []string{"10", "20", "30", "40"}
	sum := sumOfNumbers(numbers)
	fmt.Println("Sum of numbers is: ", sum)
}

func sumOfNumbers(numbers []string) int {
	sum := 0
	for i:=0; i<len(numbers); i++ {
		i, err := strconv.Atoi(numbers[i])
		if err != nil {
			fmt.Println("Conversion not possible", err)
		}
		sum += i
	}
	return sum
}
