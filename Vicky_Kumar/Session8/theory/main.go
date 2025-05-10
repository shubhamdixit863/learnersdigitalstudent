package main

import (
	"fmt"
	"theory/filehandling"
)

type Data struct {
	name string
	age  int
}

func (d Data) Double() int {
	d.age = 2 * d.age
	return 2 * d.age
}
func (d *Data) DoubleAgePointer() int {
	d.age = 2 * d.age
	return d.age
}
func main() {
	// data := Data{
	// 	name: "John",
	// 	age:  23,
	// }
	// // fmt.Println(data.Double())
	// // fmt.Println(data.age)

	// j := data.DoubleAgePointer()
	// fmt.Println(j)
	// fmt.Println(data.age)

	file, err := filehandling.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)
}
