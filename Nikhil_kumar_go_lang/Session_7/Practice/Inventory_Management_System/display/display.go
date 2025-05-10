package display

import "fmt"

func DisplayMenu() {
	fmt.Println("\n===== Inventory Management System =====")
	fmt.Println("1. Add new product")
	fmt.Println("2. Update product")
	fmt.Println("3. Search products")
	fmt.Println("4. Delete product")
	fmt.Println("5. List all products")
	fmt.Println("6. Exit")
	fmt.Println("======================================")
}