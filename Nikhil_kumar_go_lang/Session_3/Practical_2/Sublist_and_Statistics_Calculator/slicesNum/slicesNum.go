package slicesNum

import (
	"fmt"
)

func SliceNumbers(numbers []int, start, end int) []int {
	if start < 0 || end > len(numbers) || start >= end {
		fmt.Println("Error: Invalid indices")
		return nil
	}
	return numbers[start:end]
}
