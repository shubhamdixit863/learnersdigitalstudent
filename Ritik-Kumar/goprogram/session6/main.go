package main

import (
	"fmt"
	"session6/lru"
)

func main() {
	cache := lru.NewLRUCache()

	for {
		fmt.Println("\n LRU Cache System")
		fmt.Println("1. Add Student Score")
		fmt.Println("2. Update Student Score")
		fmt.Println("3. Mark Student as Visited")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var key, value int
			fmt.Print("Enter Student ID: ")
			fmt.Scanln(&key)
			fmt.Print("Enter Score: ")
			fmt.Scanln(&value)
			cache.Add(key, value)

		case 2:
			var key, value int
			fmt.Print("Enter Student ID to Update: ")
			fmt.Scanln(&key)
			fmt.Print("Enter New Score: ")
			fmt.Scanln(&value)
			cache.Update(key, value)

		case 3:
			var key int
			fmt.Print("Enter Student ID to Mark as Visited: ")
			fmt.Scanln(&key)
			cache.Visited(key)

		case 4:
			var key int
			fmt.Print("Enter Student ID to Delete: ")
			fmt.Scanln(&key)
			cache.Delete(key)

		case 5:
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid choice, try again!")
		}
	}
}
