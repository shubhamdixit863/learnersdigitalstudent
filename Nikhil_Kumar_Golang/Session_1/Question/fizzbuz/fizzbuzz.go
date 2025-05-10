package main

import (
	"fmt"
)

func fizzbuzz(number int){
	if number % 3 == 0 && number % 5 == 0 {
		fmt.Println("fizzbuzz")
	}else if number % 3 == 0 {
		fmt.Println("Fizz")
	}else if number % 5 == 0 {
		fmt.Println("buzz")
	}else{
		fmt.Println("wrong input")
	}
}

func main() {
	fizzbuzz(3)
	fizzbuzz(5)
	fizzbuzz(15)
	fizzbuzz(4)

}