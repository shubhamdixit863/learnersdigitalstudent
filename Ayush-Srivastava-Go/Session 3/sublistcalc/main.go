package main

import (
	"calculator/slice"
	"fmt"
)

func main() {
	var cont string
	arr := read()
	for i := 0; i == 0; {
		slice.SliceList(arr)
		for j := 0; j == 0; {
			fmt.Println("Do you want to slice again?(yes/no): ")
			fmt.Scanf("%s\n", &cont)
			if cont == "no" {
				fmt.Println("Exiting the application, Goodbye!")
				i++
				j++
			} else if cont == "yes" {
				j++
			} else {
				fmt.Println("Invalid input. Please enter yes or no.")
			}
		}

	}
}

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}
func read() []int {
	fmt.Println("Enter integers separated by space: ")
	x := input([]int{}, nil)
	return x
}