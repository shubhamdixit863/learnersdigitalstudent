package slice

import (
	"fmt"
	"practical3/operations"
)

func SliceList(slice []int) {
	var start, end int
	var final []int
	for i := 0; i == 0; {

		fmt.Println("Enter start index: ")
		fmt.Scanf("%d\n", &start)
		fmt.Println("Enter end index: ")
		fmt.Scanf("%d\n", &end)
		if start <= end {
			i++
		} else {
			fmt.Println("starting index should be less than or equal to end index")
		}
	}

	final = slice[start:end]
	fmt.Println(final)
	operations.Sum(final)
	operations.Average(final)
	operations.Max(final)
}
