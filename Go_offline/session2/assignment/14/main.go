// Use a for loop to reverse a given string.

package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(reverse(s))
}


func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}