package main

import (
	"fmt"
	"os"
)

type Data struct {
	name string
	age  int
}

// value recivers do not modify the data.age
// func (d Data) Doubleage() int {
// 	return 2 * d.age
// }

// pointer reciver will modify the data.age
func (d *Data) Doubleage() int {
	d.age = 2 * d.age
	return d.age
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Id         int
	Name       string
	Age        int
	Department string
	Salary     float64
}
type Address struct {
	Street  string
	City    string
	Zipcode string
	Skills  []string
}

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(data), err

}

func main() {
	// fmt.Println("hello world")
	// data := Data{
	// 	name: "imran",
	// 	age:  22,
	// }
	// var arr []Data
	// arr = append(arr, data)
	// fmt.Println(arr)

	// data.Doubleage()
	// fmt.Println(data.age)
	// p := Person{"rahul", 12}
	// jsondata, _ := json.Marshal(p)

	// fmt.Println(string(jsondata)) // Json Data:=> {"Name":"rahul","Age":12}
	// fmt.Println(jsondata)         // Marshal Data:=> [123 34 78 97 109 101 34 58 34 114 97 104 117 108 34 44 34 65 103 101 34 58 49 50 125]
	// var unmarshalledP Person
	// json.Unmarshal(jsondata, &unmarshalledP)
	// fmt.Println(unmarshalledP)
	var emp []Employee
	file, err := ReadFile("practical/employee.json")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(file)
}
