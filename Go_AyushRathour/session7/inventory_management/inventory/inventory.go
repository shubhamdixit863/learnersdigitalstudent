package inventory

import (
	"fmt"
	"inventory_management/product"
)

type Inventory struct {
	products map[string]product.Product
}

func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[string]product.Product),
	}
}


func (inv *Inventory) AddProduct(prod product.Product) {
	inv.products[prod.ID] = prod
	fmt.Println("Product added successfully!")
}

func (inv *Inventory) UpdateProduct(id string, price float64, quantity int) {
	if prod, exists := inv.products[id]; exists {
		prod.Price = price
		prod.Quantity = quantity
		inv.products[id] = prod
		fmt.Println("Product updated successfully!")
	} else {
		fmt.Println("Product not found!")
	}
}

func (inv *Inventory) DeleteProduct(id string) {
	if _, exists := inv.products[id]; exists {
		delete(inv.products, id)
		fmt.Println("Product deleted successfully!")
	} else {
		fmt.Println("Product not found!")
	}
}

func (inv *Inventory) SearchByName(name string) []product.Product {
	var results []product.Product
	for _, prod := range inv.products {
		if prod.Name == name {
			results = append(results, prod)
		}
	}
	return results
}

func (inv *Inventory) SearchByCategory(category string) []product.Product {
	var results []product.Product
	for _, prod := range inv.products {
		if prod.Category == category {
			results = append(results, prod)
		}
	}
	return results
}
func (inv *Inventory) ListAllProducts() {
	fmt.Println("\n--- Inventory List ---")
	for _, prod := range inv.products {
		fmt.Printf("ID: %s | Name: %s | Category: %s | Price: %.2f | Quantity: %d\n",
			prod.ID, prod.Name, prod.Category, prod.Price, prod.Quantity)
	}
}
