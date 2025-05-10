package main

import "fmt"

type Human struct {
	name    string
	age     int
	address string
}

func (h Human) sleep() {
	fmt.Println("Good night", h.name)
}

func main() {
	h := Human{
		name:    "Anubhav",
		age:     22,
		address: "Patna, Bihar",
	}
	fmt.Println(&h.name)
	// hPointer := new(Human)
	// fmt.Println((hPointer))
	h.sleep()
}
