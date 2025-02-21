package main

import (
	"fmt"
	"reflect"
)

func main() {
	var input string
	fmt.Println("Enter a single character: ")
	fmt.Scanln(&input)
	fmt.Println(reflect.TypeOf(input))

	// Convert the first character of the string to a rune
	if len(input) > 0 {
		c := rune(input[0])
		fmt.Println(reflect.TypeOf(c))
		i := charToInt(c)
		fmt.Println(i)
	} else {
		fmt.Println("No character entered.")
	}
}

func charToInt(c rune) int {
	i := int(c)
	return i
}
