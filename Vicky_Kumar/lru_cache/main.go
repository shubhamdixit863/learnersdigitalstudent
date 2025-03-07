package main

import (
	"fmt"
	"lru_cache/operations"
)

func main() {
	for {

		fmt.Println("\n\n1. Put")
		fmt.Println("2. Get")
		fmt.Println("3. Print entire cache")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice:")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var key int
			var value string
			fmt.Print("\033[H\033[2J")
			fmt.Print("Enter key (int): ")
			fmt.Scanln(&key)
			fmt.Print("Enter value (string): ")
			fmt.Scanln(&value)
			operations.Put(key, value)
			fmt.Println("Added..!")

		case 2:
			var key int
			fmt.Print("\033[H\033[2J")
			fmt.Print("Enter key: ")
			fmt.Scanln(&key)
			value := operations.Get(key)
			if value == "-1" {
				fmt.Println("Key not found..!")
			} else {
				fmt.Println("Value:", value)
			}

		case 3:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Cache is:\n", operations.Mp)

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("\nInvalid choice!")
		}
	}
}
