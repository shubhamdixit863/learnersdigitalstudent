package main

import (
	"fmt"
	"session7/practical/inventory_management/inventory"
)

func main() {
	for {
		fmt.Println("Inventory Management System")
		fmt.Println("1 Add Product")
		fmt.Println("2️ Update Product")
		fmt.Println("3️ Delete Product")
		fmt.Println("4️ List Products")
		fmt.Println("5️ Exit")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			inventory.Add()
		case 2:
			inventory.Update()
		case 3:
			inventory.Delete()
		case 4:
			inventory.List()
		case 5:
			fmt.Println("🚪 Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}
