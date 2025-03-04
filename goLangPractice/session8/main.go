package main

import (
	"fmt"
	"session8/filehandling"
)

type Data struct {
	Name string `json: "name, omitempty"` //--- Tags
	Age  int    `json: "age, omitempty"`
}

// value receiver do not modify the underlying struct
func (d Data) DoubleAge() int {
	d.Age = 2 * d.Age
	return d.Age
}

// Pointer receiver modify the underlying struct
func (d *Data) DoubleAgePointer() int {
	d.Age = 2 * d.Age
	return d.Age
}

func main() {
	// data := Data{
	// 	Name: "John",
	// 	Age:  23,
	// }

	// c := data.DoubleAge()
	// fmt.Println(c)
	// fmt.Println(data.age)

	// j := data.DoubleAgePointer()
	// fmt.Println(j)
	// fmt.Println(data.age)

	// s, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(s))

	// s := []byte(`{"name": "Sagar", "age": 21}`)
	// d := &Data{}
	// err := json.Unmarshal(s, d)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(d)

	file, err := filehandling.Readfile("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(file)
}
