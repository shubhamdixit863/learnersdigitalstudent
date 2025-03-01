package inventory

import (
	"fmt"
	"session7/problems/product"
)

var Products []product.Product

func AddProduct(id int, name string, category string, price int, quantity int) {
	p := product.NewProduct(id, name, category, price, quantity)
	Products = append(Products, p)
	fmt.Println("Product added:", p)
}

func UpdateProduct(id int, price int, quantity int) {
	for i, p := range Products {
		if p.Id == id {
			Products[i].Price = price
			Products[i].Quantity = quantity
			fmt.Println("Product updated:", Products[i])
			return
		}
	}
	fmt.Println("Product not found")
}

func DeleteProduct(id int) {
	for i, p := range Products {
		if p.Id == id {
			Products = append(Products[:i], Products[i+1:]...)
			fmt.Println("Product deleted with ID:", id)
			return
		}
	}
	fmt.Println("Product not found")
}

func SearchByName(name string) {
	for _, p := range Products {
		if p.Name == name {
			fmt.Println(p)
		}
	}
}

func SearchByCategory(category string) {
	for _, p := range Products {
		if p.Category == category {
			fmt.Println(p)
		}
	}
}

func ListProducts() {
	for _, p := range Products {
		fmt.Println(p)
	}
}
