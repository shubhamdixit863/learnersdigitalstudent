package main

import (
	"fmt"
	"os"
)

func main() {
	var name = os.Args[1]
	age := os.Args[2]
	bank_balance := os.Args[3]
	fmt.Println("Name is ", name, "age is ", age, "and his bank balance is ", bank_balance)
}
