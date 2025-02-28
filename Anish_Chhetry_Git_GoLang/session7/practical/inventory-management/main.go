//Develop an Inventory Management System in Go that enables users to:

package main

import (
	"fmt"
	"problems/inventory"
	"problems/product"
)

func main() {
	inventoryAll := inventory.CreateInventory()
	inventoryAll = product.Add(1, "AirForce1", "Shoe", 20, 40, inventoryAll)
	inventoryAll = product.Add(2, "T100", "Brush", 2, 89, inventoryAll)
	inventoryAll = product.Add(3, "Nike", "Hat", 8, 77, inventoryAll)
	fmt.Println("After Adding:")
	inventory.ListProducts(inventoryAll)

	inventoryAll = product.Update(1, "AirForce1", "Shoe", 10, 60, inventoryAll)
	fmt.Println("\nAfter Updating:")
	inventory.ListProducts(inventoryAll)
	fmt.Println(inventory.Search(1, inventoryAll))

	inventoryAll = product.Delete(1, inventoryAll)
	fmt.Println("\nAfter Deleting:")
	inventory.ListProducts(inventoryAll)
	fmt.Println(inventory.Search(1, inventoryAll))

	inventoryAll = product.Add(4, "Nike", "Hat", 8, 77, inventoryAll)
	fmt.Println("\nAfter Adding:")
	inventory.ListProducts(inventoryAll)

}
