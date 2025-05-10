package main

import (
	"fmt"
	"inventory_management_system/inventory"
	"inventory_management_system/product"
)

func main() {
	fmt.Println("Inventory Management System")

	inv := inventory.NewInventory()

	prod1 := product.NewProduct(101, "Laptop", "Electronics", 75000.0, 10)
	prod2 := product.NewProduct(102, "Mouse", "Electronics", 250.0, 50)
	prod3 := product.NewProduct(103, "Notebook", "Stationery", 50.0, 100)

	inv.AddProduct(prod1)
	inv.AddProduct(prod2)
	inv.AddProduct(prod3)

	inv.ListProducts()

	inv.SearchByName("Laptop")
	inv.SearchByCategory("Electronics")

	inv.UpdateProduct(101, 72000.0, 8)

	inv.ListProducts()

	inv.DeleteProduct(102)

	inv.ListProducts()
}
