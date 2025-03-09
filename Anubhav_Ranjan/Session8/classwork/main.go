package main

import (
	"classwork/filehandling"
	"fmt"
)

// type Data struct {
// 	Name string
// 	Age  int
// }

// func (d Data) doubleAge() int {
// 	return d.Age * 2
// }

// func (d *Data) doubleAge_2()int{
// 	d.Age =d.Age*2
// 	return d.Age
// }

func main() {
	// data := Data{Name: "Anubhav", Age: 22}
	// c := data.doubleAge()
	// fmt.Println(c,data.Age)
	// d:=data.doubleAge_2()
	// fmt.Println(d,c,data.Age)

	data, err := filehandling.Readfile("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
