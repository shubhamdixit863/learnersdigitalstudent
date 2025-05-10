// inventory/inventory.go
package inventory

import (
	"Inventory_Management_System/product"
	"strings"
)

type Inventory struct {
	products     []product.Product
	productsByID map[string]int
}

func NewInventory() *Inventory {
	return &Inventory{
		products:     make([]product.Product, 0),
		productsByID: make(map[string]int),
	}
}

func (inv *Inventory) AddProduct(p product.Product) {
	inv.products = append(inv.products, p)
	inv.productsByID[p.ID] = len(inv.products) - 1
}

func (inv *Inventory) GetProductByID(id string) (product.Product, bool) {
	if idx, exists := inv.productsByID[id]; exists {
		return inv.products[idx], true
	}
	return product.Product{}, false
}

func (inv *Inventory) UpdateProduct(p product.Product) bool {
	if idx, exists := inv.productsByID[p.ID]; exists {
		inv.products[idx] = p
		return true
	}
	return false
}

func (inv *Inventory) DeleteProduct(id string) bool {
	idx, exists := inv.productsByID[id]
	if !exists {
		return false
	}

	lastIdx := len(inv.products) - 1
	if idx < lastIdx {
		inv.products[idx] = inv.products[lastIdx]
		inv.productsByID[inv.products[idx].ID] = idx
	}
	inv.products = inv.products[:lastIdx]
	delete(inv.productsByID, id)
	return true
}

func (inv *Inventory) SearchByName(name string) []product.Product {
	var result []product.Product
	lowerName := strings.ToLower(name)

	for _, p := range inv.products {
		if strings.Contains(strings.ToLower(p.Name), lowerName) {
			result = append(result, p)
		}
	}
	return result
}

func (inv *Inventory) ListAllProducts() []product.Product {
	result := make([]product.Product, len(inv.products))
	copy(result, inv.products)
	return result
}
