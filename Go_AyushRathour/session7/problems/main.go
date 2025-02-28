package main

import "fmt"

type Human struct {
	name     string
	age      int
	Isplaced bool
	address  string
}

// method here// value receiver
func (h Human) sleep() {
	fmt.Println("Hello ", h.name)
}

func main(){
	//object
	h:=&Human{
		name: "Ayush",
		age : 22
		Isplaced: true,
		address: "Jsr"
	}
	hPpointer: =new(Human)
	fmt.Println(hPpointer)
	//hPpointer.address="Mumbai"

	//hPpointer :=&h //pointer to object


	h.Sleep()
}