package main

import "fmt"

type human struct {
	name string
	age  int
}

func (h human) Sleep() {
	h.name = "kkk"
	fmt.Println("sleeep", h.name, h.age)

}

func main() {
	h1 := &human{name: "om", age: 20}
	h2 := human{name: "som", age: 30}
	h2.Sleep()
	h1.Sleep()

	u := new(human)
	u.Sleep()
}
