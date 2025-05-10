package inventory

import (
	"ims/product"
)

type Inventory struct{
	products map[string]*product.Product // we store {id->product info}
}

func NewInventory() *Inventory{
	return &Inventory{products: make(map[string]*product.Product)}
}

func (inv *Inventory) AddProduct(p *product.Product){
	inv.products[p.ID] = p
}

func (inv *Inventory) UpdateProduct(id string, price float64, quantity int){
	if p, ok := inv.products[id]; ok{
		p.Price = price
		p.Quantity = quantity
	}
}

func (inv *Inventory) DeleteProduct(id string){
	delete(inv.products, id)
}

func (inv *Inventory) SearchByName(name string) []*product.Product {
	var result []*product.Product
	for _, p := range inv.products {
		if p.Name == name {
			result = append(result, p)
		}
	}
	return result
}

func (inv *Inventory) SearchByCategory(category string) []*product.Product{
	var res []*product.Product

	for _, p := range inv.products{
		if p.Category == category{
			res = append(res, p)
		}
	}
	return res
}

func (inv *Inventory) ListAllProducts() []*product.Product {

	var res []*product.Product

	for _, p := range inv.products{
		res = append(res, p)
	}

	return res
}