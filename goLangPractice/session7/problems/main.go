package main

import (
	"fmt"
	"session7/problems/inventory"
)

func main() {
	inventory.AddProduct(1, "Laptop", "Electronics", 800, 10)
	inventory.AddProduct(2, "Phone", "Electronics", 500, 20)

	fmt.Println("All Products:")
	inventory.ListProducts()

	fmt.Println("\nSearch by Name: Phone")
	inventory.SearchByName("Phone")

	fmt.Println("\nSearch by Category: Electronics")
	inventory.SearchByCategory("Electronics")

	fmt.Println("\nUpdating Product")
	inventory.UpdateProduct(1, 900, 5)
	inventory.ListProducts()

	fmt.Println("\nDeleting Product")
	inventory.DeleteProduct(2)
	inventory.ListProducts()
}
