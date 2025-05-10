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
		name:    "Vicky",
		age:     21,
		address: "Jamshedpur",
	}
	fmt.Println(&h.name)
	// hPointer := new(Human)
	// fmt.Println((hPointer))
	h.sleep()
}
