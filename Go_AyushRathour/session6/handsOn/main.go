package main

import (
	"fmt"
	"handsOn/insert"
)

func main() {
	

	for {
		var s string
		fmt.Println("Type put or get and exit to end")
		fmt.Scanln(&s)

		if s == "put" {
			Insert()
		} else if s == "get" {
			fmt.Println(Display())
		} else {
			return
		}
	}

}

func Insert() {
	var key int
	var value int
	fmt.Println("Enter the key value")
	fmt.Scanln(&key, &value)
	insert.Put(key, value)

}
func Display() int {

	var key int
	fmt.Println("Enter the key")
	fmt.Scanln(&key)

	return insert.Get(key)

}
