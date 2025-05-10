package main

import "fmt"

type Human struct {
	name     string
	age      int
	isPlaced bool
	address  string
}

func (h Human) Say() {
	fmt.Printf("Good morning %s", h.name)
}

func main() {

	a := Human{
		name:     "Ayush",
		age:      23,
		isPlaced: true,
		address:  "Jaunpur",
	}

	a.Say()
}