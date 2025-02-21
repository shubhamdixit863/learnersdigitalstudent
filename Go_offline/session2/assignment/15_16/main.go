package main

import "fmt"

func main(){
	var n int
	fmt.Scanln(&n)
	isPrime(n)
	isEven(n)
}

func isPrime(n int) {
	if n < 2 {
		fmt.Printf("%d is Not a prime number \n", n)
		return
	}
	for i:=2; i*i<=n; i++ {
		if n%i == 0 {
			fmt.Printf(" %d is Not a prime number \n", n)
			return
		}
	}
	fmt.Printf("%d Prime number \n", n)
}

func isEven(n int) {
	if n%2 == 0 {
		fmt.Printf("%d is Even number \n", n)
	} else {
		fmt.Printf("%d is Odd number \n", n)
	}
}