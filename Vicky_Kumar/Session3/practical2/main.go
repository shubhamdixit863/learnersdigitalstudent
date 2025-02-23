package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Println("Enter number of elements:")
	fmt.Scanln(&n)
	var slc []int
	fmt.Println("Enter integers :")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanln(&num)
		slc = append(slc, num)
	}
	for {
		var start int
		var end int
		fmt.Println("Enter start index:")
		fmt.Scanln(&start)
		fmt.Println("Enter end index:")
		fmt.Scanln(&end)
		Slice(slc, start, end)
		fmt.Println("Do you want to slice again? (yes/no):")

		var choice string
		fmt.Scanln(&choice)
		if choice != "yes" {
			fmt.Println("Exiting the application. Goodbye!")
			break
		}
	}

}
func Slice(slc []int, start int, end int) {
	if start < 0 || end >= len(slc) || start > end {
		fmt.Println("Please enter proper start and end index!")
	}
	slcSub := slc[start:end]
	sum := 0
	var avg float64
	maxm := math.MinInt
	for i := 0; i < len(slcSub); i++ {
		sum += slcSub[i]
		maxm = max(maxm, slcSub[i])
	}
	avg = float64(sum) / float64((len(slc)))
	fmt.Println("\nSliced Sublist: ", slcSub)
	fmt.Println("Sum: ", sum)
	fmt.Println("Average: ", avg)
	fmt.Printf("Maximum: %d \n", maxm)
}
