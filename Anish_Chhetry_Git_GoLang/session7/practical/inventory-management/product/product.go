package product

import "problems/inventory"

func Add(id int, name string, category string, price int, quantity int, inv inventory.Inventory) inventory.Inventory {
	inv.Name[id] = name
	inv.Category[id] = category
	inv.Price[id] = price
	inv.Quantity[id] = quantity
	return inv
}

func Update(id int, name string, category string, price int, quantity int, inv inventory.Inventory) inventory.Inventory {
	inv.Name[id] = name
	inv.Category[id] = category
	inv.Price[id] = price
	inv.Quantity[id] = quantity
	return inv
}
func Delete(id int, inv inventory.Inventory) inventory.Inventory {
	delete(inv.Name, id)
	delete(inv.Category, id)
	delete(inv.Price, id)
	delete(inv.Quantity, id)
	return inv
}
