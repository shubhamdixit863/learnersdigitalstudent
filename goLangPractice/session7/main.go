package main

import "fmt"

type Human struct {
	name     string
	age      int
	isPlaced bool
	address  string
}

func (h Human) sleep() {
	fmt.Println("Good night", h.name)
}

func main() {
	h := &Human{
		name:     "Sagar",
		age:      21,
		isPlaced: true,
		address:  "Bihar",
	}
	// hPointer := new(Human)

	fmt.Println(&h.address)
	h.sleep()
}
