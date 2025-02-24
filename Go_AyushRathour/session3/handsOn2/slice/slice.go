package slice

import (
	"fmt"
)

func Slice(arr[]int, start int, end int)[]int{
	if start < 0 {
		fmt.Println("Invalid start index")
		return nil
	}
	if end > len(arr) {
		fmt.Println("Invalid end index")
		return nil
	}
	if start > end {
		fmt.Println("Invalid start and end index")
		return nil
	}
	return arr[start:end]
}