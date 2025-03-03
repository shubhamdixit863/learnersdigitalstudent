package main

import (
	"fmt"
	"inventory_Management/product"
)

func main() {
	fmt.Println("Welcome")
	for true {

		fmt.Println("1. Create\n2. List Product\n3. Search by Category\n4. Search by name \n5.Delete product by id ")
		fmt.Println("enter you choice")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("enter the details id name category quantity price")
			var id, quantity, price int
			var name, category string
			fmt.Scanln(&id, &name, &category, &quantity, &price)
			product.NewProduct(id, name, category, quantity, price)
		}
		if choice == 2 {
			product.ListProduct()

		}
		if choice == 3 {
			var cat string
			fmt.Println("enter category")
			fmt.Scanln(&cat)
			product.SearchProductCategory(cat)

		}
		

		if choice == 4 {
			var name string
			fmt.Println("enter name")
			fmt.Scanln( &name)
			product.SearchProductName(name)

		}
		if choice == 5 {
			var id int
fmt.Println("enter id")
			fmt.Scanln( &id)
			product.DeleteProduct(id)
		}
	}
	// product.NewProduct(0, "ss", "book", 5, 70)
	// product.NewProduct(0, "kk", "book", 5, 70)
	// product.NewProduct(0, "tt", "pen", 5, 70)
	// product.ListProduct()
	// product.SearchProductCategory("book")
	// product.SearchProductName("ss")
	// product.DeleteProduct(0)
}
