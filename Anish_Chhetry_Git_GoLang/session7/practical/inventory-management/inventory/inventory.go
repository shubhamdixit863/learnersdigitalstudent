package inventory

import "fmt"

type Inventory struct {
	Name     map[int]string
	Category map[int]string
	Price    map[int]int
	Quantity map[int]int
}

func CreateInventory() Inventory {
	return Inventory{
		Name:     make(map[int]string),
		Category: make(map[int]string),
		Price:    make(map[int]int),
		Quantity: make(map[int]int),
	}
}
func Search(id int, inv Inventory) (int, string, string, int, int) {
	name := inv.Name[id]
	category := inv.Category[id]
	price := inv.Price[id]
	quantity := inv.Quantity[id]
	return id, name, category, price, quantity
}
func ListProducts(inv Inventory) {
	for id := range inv.Name {
		fmt.Print(id, " ")
		fmt.Print(inv.Name[id], " ")
		fmt.Print(inv.Category[id], " ")
		fmt.Print(inv.Price[id], " ")
		fmt.Print(inv.Quantity[id], "\n")
	}
}
