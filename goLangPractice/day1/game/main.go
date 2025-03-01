package main

import "fmt"

func FizzBuzz(num int) {
	if num%3 == 0 && num%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if num%3 == 0 {
		fmt.Println("Fizz")
	} else if num%5 == 0 {
		fmt.Println("Buzz")
	}
}

func main() {
	FizzBuzz(3)
	FizzBuzz(5)
	FizzBuzz(15)
}
