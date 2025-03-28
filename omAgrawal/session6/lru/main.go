package main

import (
	"fmt"
	"lru/data"
)

func main() {
	fmt.Println("welcome to LRU cache ")
	fmt.Println("------------Enter the cache capacity--------------------")

	var v int
	fmt.Scanln(&v)
	data.Capacity = v

	fmt.Println("the lru cache is initialised")

	for true {
		var choice int
		fmt.Println("Enter your Choice")
		fmt.Println("1. Put \n2. Get ")

		fmt.Scanln(&choice)

		if choice == 1 {
			var key int
			var value int

			fmt.Println("Enter your key")
			fmt.Scanln(&key)

			fmt.Println("Enter your value")
			fmt.Scanln(&value)

			data.Put(key, value)
		}

		if choice == 2 {
			var key int
			fmt.Println("Enter your key")
			fmt.Scanln(&key)
			data.Get(key)
		}
	}

}
