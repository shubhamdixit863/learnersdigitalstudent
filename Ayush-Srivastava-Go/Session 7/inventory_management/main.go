package main

import (
	"fmt"
	"ims/inventory"
	"ims/product"
)

func main() {
	inv := inventory.NewInventory()

	p1 := &product.Product{ID: "101", Name: "Laptop", Category: "Electronics", Price: 75000, Quantity: 10}
	p2 := &product.Product{ID: "102", Name: "Phone", Category: "Electronics", Price: 50000, Quantity: 20}
	p3 := &product.Product{ID: "103", Name: "Shoes", Category: "Fashion", Price: 3000, Quantity: 50}

	inv.AddProduct(p1)
	inv.AddProduct(p2)
	inv.AddProduct(p3)

	fmt.Println("All Products:")
	for _, p := range inv.ListAllProducts() {
		fmt.Println(p)
	}

	fmt.Println("\nSearching for product with name 'Laptop':")
	for _, p := range inv.SearchByName("Laptop") {
		fmt.Println(p)
	}

	fmt.Println("\nSearching for category 'Electronics':")
	for _, p := range inv.SearchByCategory("Electronics") {
		fmt.Println(p)
	}

	fmt.Println("\nUpdating Laptop price and quantity...")
	inv.UpdateProduct("101", 72000, 8)
	fmt.Println(inv.SearchByName("Laptop"))

	fmt.Println("\nDeleting product with ID 103 (Shoes)...")
	inv.DeleteProduct("103")
	fmt.Println("All Products after deletion:")
	for _, p := range inv.ListAllProducts() {
		fmt.Println(p)
	}
}
