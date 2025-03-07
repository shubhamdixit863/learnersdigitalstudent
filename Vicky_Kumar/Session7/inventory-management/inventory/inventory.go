package inventory

import (
	"fmt"
	"inventory-management/product"
)

func AddProduct(id int, name string, category string, price float64, quantity int) {
	prod := product.Product{
		Id: id, Name: name, Category: category, Price: price, Quantity: quantity}
	product.Products = append(product.Products, prod)

}

func UpdateProduct(id int, name string, category string, price float64, quantity int) {
	for i := 0; i < len(product.Products); i++ {
		if product.Products[i].Id == id {
			product.Products[i].Name = name
			product.Products[i].Category = category
			product.Products[i].Price = price
			product.Products[i].Quantity = quantity
			break
		}
	}
}
func DeleteProduct(id int) {
	for i := 0; i < len(product.Products); i++ {
		if product.Products[i].Id == id {
			product.Products = append(product.Products[:i], product.Products[i+1:]...)
			break
		}
	}
}

func SearchProduct() {
	fmt.Println("Enter 1 to Search by Name: ")
	fmt.Println("Enter 2 to Search by Category: ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		var name string
		fmt.Print("Enter name of product: ")
		fmt.Scanln(&name)
		for i := 0; i < len(product.Products); i++ {
			if product.Products[i].Name == name {
				fmt.Println("\nProduct Details: ")
				fmt.Printf("ID:\t%d\n", product.Products[i].Id)
				fmt.Printf("Name:\t%s\n", product.Products[i].Name)
				fmt.Printf("Category:\t%s\n", product.Products[i].Category)
				fmt.Printf("Price:\t%0.2f\n", product.Products[i].Price)
				fmt.Printf("Quantity:\t%d\n", product.Products[i].Quantity)
			}
		}
		fmt.Println("No product found")
	} else {
		var category string
		fmt.Print("Enter category of product: ")
		fmt.Scanln(&category)
		for i := 0; i < len(product.Products); i++ {
			if product.Products[i].Category == category {
				fmt.Println("\nProduct Details: ")
				fmt.Printf("ID:\t%d\n", product.Products[i].Id)
				fmt.Printf("Name:\t%s\n", product.Products[i].Name)
				fmt.Printf("Category:\t%s\n", product.Products[i].Category)
				fmt.Printf("Price:\t%0.2f\n", product.Products[i].Price)
				fmt.Printf("Quantity:\t%d\n", product.Products[i].Quantity)
			}
		}
		fmt.Println("No product found")
	}
}

func ListAllProduct() {
	if len(product.Products) == 0 {
		fmt.Println("No Product Found..!!")
		return
	}
	for i := 0; i < len(product.Products); i++ {
		fmt.Printf("Product %d: \n", i+1)
		fmt.Printf("ID: \t\t%d\n", product.Products[i].Id)
		fmt.Printf("Name: \t\t%s\n", product.Products[i].Name)
		fmt.Printf("Category: \t\t%s\n", product.Products[i].Category)
		fmt.Printf("Price: \t\t%.2f\n", product.Products[i].Price)
		fmt.Printf("Quantity: \t\t%d\n", product.Products[i].Quantity)
	}
}
