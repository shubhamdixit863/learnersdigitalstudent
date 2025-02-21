// Accept a single character and determine if it is a vowel or a consonant.

package main

import "fmt"

func main() {
	var c string
	fmt.Scanln(&c)
	if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
		fmt.Printf("%s is a vowel\n", c)
	} else {
		fmt.Printf("%s is a consonant\n", c)
	}
}