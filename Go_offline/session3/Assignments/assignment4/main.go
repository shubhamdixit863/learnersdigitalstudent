// Write a function that converts a boolean value to an integer (true -> 1, false -> 0). Take user input as "true" or "false" and convert it.

package main

import (
	"fmt"
)

func main() {
	var f bool
	fmt.Println("Enter boolean value: ")
	fmt.Scanln(&f)
	i := boolToInt(f)
	fmt.Println(i)
}

func boolToInt(f bool) int {
	if f {
		return 1
	} else {
		return 0
	}
}