package slice

import (
	"fmt"
	"practical3/operations"
)

func SliceList(slice []int) {
	var a, b int
	var final []int
	fmt.Println("Enter start index: ")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Enter end index: ")
	fmt.Scanf("%d\n", &b)
	final = slice[a:b]
	fmt.Println(final)
	operations.Sum(final)
	operations.Average(final)
	operations.Max(final)
}
