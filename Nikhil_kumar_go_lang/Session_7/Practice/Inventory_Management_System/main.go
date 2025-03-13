// main.go
package main

import (
	"Inventory_Management_System/display"
	"Inventory_Management_System/inventory"
	"Inventory_Management_System/operation"
	"Inventory_Management_System/product"
	"fmt"
)

func main() {
	inv := inventory.NewInventory()

	inv.AddProduct(product.Product{
		ID:       "P001",
		Name:     "Laptop",
		Category: "Electronics",
		Price:    999.99,
		Quantity: 10,
	})

	inv.AddProduct(product.Product{
		ID:       "P002",
		Name:     "Desk Chair",
		Category: "Furniture",
		Price:    149.99,
		Quantity: 20,
	})

	for {
		display.DisplayMenu()
		var option string
		fmt.Print("Select an option: ")
		fmt.Scan(&option)

		switch option {
		case "1":
			operation.AddProduct(inv)
		case "2":
			operation.UpdateProduct(inv)
		case "3":
			operation.SearchProducts(inv)
		case "4":
			operation.DeleteProduct(inv)
		case "5":
			operation.ListAllProducts(inv)
		case "6":
			fmt.Println("Exiting the Inventory Management System. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
