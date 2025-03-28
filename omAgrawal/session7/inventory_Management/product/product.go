package product

import (
	"fmt"
)

type Product struct {
	id       int
	name     string
	category string
	quantity int
	price    int
}

var ProductList []Product

func NewProduct(id int, name string, category string, Quantity int, price int) Product {
	v := Product{id, name, category, Quantity, price}
	ProductList = append(ProductList, v)
	fmt.Println("new Product Created", v)
	return v

}
func UpdateProduct(id int, price int, quantity int) {
	ProductList[id].price = price
	ProductList[id].quantity = quantity
}
func SearchProductName(name string) {

	fmt.Println("Search result for", name)
	for _, v := range ProductList {

		if v.name == name {
			fmt.Println(v)
		}
	}

}

func SearchProductCategory(category string) {
	fmt.Println("The Products of the category ", category, "are")
	for i, v := range ProductList {

		if v.category == category {
			fmt.Println(i, v)
		}
	}

}
func DeleteProduct(id int) {

	ProductList = append(ProductList[:id], ProductList[id+1:]...)
	fmt.Println("product deleted Succesfully")

}

func ListProduct() {
	for i, v := range ProductList {

		fmt.Println("Product", i, v)

	}
}
