package main

import "fmt"

type Human struct {
	name     string
	age      int
	isplaced bool
}

func (h Human) sleep() {
	fmt.Println("good night!",h.name)
}

func main() {
h:=&Human{
	name:"shivam",
	age:22,
	isplaced: true,
}


hPointer :=new(Human) // new memory allocation
hPointer.name="annay"


hPointer.sleep()
h.sleep()
}


