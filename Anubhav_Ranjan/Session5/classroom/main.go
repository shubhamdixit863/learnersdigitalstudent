package main

import "fmt"

func main() {
	// Reverse a String:
	str := "Hello"
	revStr := reverseString(str)
	fmt.Println("The Reverse string is ", revStr)

}

func reverseString(str string) string {
	n := len(str)
	str2 := []rune(str)
	for i := 0; i < n/2; i++ {
		str2[i], str2[n-i-1] = str2[n-i-1], str2[i]
	}

	return string(str2)

	// for i:=0; i<n/2; i++ {
	// 	str[i], str[n-i-1] = str[n-i-1], str[i]  // string is immutable..so we can't assign a new value once string is created..Although it works for Arr..
	// }
}
