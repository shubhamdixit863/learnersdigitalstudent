package inventory

import (
	"fmt"
	"inventory_management_system/product"
)

type Inventory struct {
	products map[int]*product.Product
}

func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[int]*product.Product),
	}
}

func (inv *Inventory) AddProduct(prod *product.Product) {
	if _, exists := inv.products[prod.ID]; exists {
		fmt.Println("Product with this ID already exists.")
		return
	}
	inv.products[prod.ID] = prod
	fmt.Println("Product added successfully.")
}

func (inv *Inventory) UpdateProduct(id int, price float64, quantity int) {
	if prod, exists := inv.products[id]; exists {
		prod.Price = price
		prod.Quantity = quantity
		fmt.Println("Product updated successfully.")
	} else {
		fmt.Println("Product not found.")
	}
}

func (inv *Inventory) DeleteProduct(id int) {
	if _, exists := inv.products[id]; exists {
		delete(inv.products, id)
		fmt.Println("Product deleted successfully.")
	} else {
		fmt.Println("Product not found.")
	}
}

func (inv *Inventory) SearchByName(name string) {
	found := false
	for _, prod := range inv.products {
		if prod.Name == name {
			fmt.Printf("Found: %+v\n", *prod)
			found = true
		}
	}
	if !found {
		fmt.Println("No product found with this name.")
	}
}

func (inv *Inventory) SearchByCategory(category string) {
	found := false
	for _, prod := range inv.products {
		if prod.Category == category {
			fmt.Printf("Found: %+v\n", *prod)
			found = true
		}
	}
	if !found {
		fmt.Println("No products found in this category.")
	}
}

func (inv *Inventory) ListProducts() {
	if len(inv.products) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	fmt.Println("\nInventory List:")
	for _, prod := range inv.products {
		fmt.Printf("ID: %d, Name: %s, Category: %s, Price: %.2f, Quantity: %d\n",
			prod.ID, prod.Name, prod.Category, prod.Price, prod.Quantity)
	}
}
