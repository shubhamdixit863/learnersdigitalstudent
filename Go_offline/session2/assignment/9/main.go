// Write a function that checks whether a given string is a palindrome.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Printf("The string is a palindrome: %t\n", isPalindrome(s))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i:=0; i<len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}