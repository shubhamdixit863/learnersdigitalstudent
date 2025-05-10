package main

import (
	"fmt"
	"inventory-management/inventory"
)

var options = []string{
	"1. Add Product",
	"2. Update Products",
	"3. Search Product",
	"4. Delete Product",
	"5. View All Products",
	"6. Exit"}

func main() {

	fmt.Println("\t\tProduct Inventory Manager:")
	for {
		for i := 0; i < len(options); i++ {
			fmt.Println(options[i])
		}
		fmt.Printf("\n")
		fmt.Print("Enter Choice: ")
		var choice int
		fmt.Scanln(&choice)
		fmt.Print("\033[H\033[2J")
		switch choice {
		case 1:
			var (
				id       int
				name     string
				category string
				price    float64
				quantity int
			)
			fmt.Print("Enter Product Id: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Product Name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter Product Category: ")
			fmt.Scanln(&category)
			fmt.Print("Enter Product Price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter Product Quantity: ")
			fmt.Scanln(&quantity)
			inventory.AddProduct(id, name, category, price, quantity)
			fmt.Print("\033[H\033[2J")
			fmt.Println("Product Added Successfully..!")
		case 2:
			var (
				id       int
				name     string
				category string
				price    float64
				quantity int
			)
			fmt.Print("Enter Product Id to Update: ")
			fmt.Scanln(&id)
			fmt.Print("Enter new Product Name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter new Product Category: ")
			fmt.Scanln(&category)
			fmt.Print("Enter new Product Price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter new Product Quantity: ")
			fmt.Scanln(&quantity)
			inventory.UpdateProduct(id, name, category, price, quantity)
			fmt.Print("\033[H\033[2J")
			fmt.Println("Product Updated Successfully..!")
		case 3:
			inventory.SearchProduct()
		case 4:
			var id int
			fmt.Print("Enter Product Id to Delete: ")
			fmt.Scanln(&id)
			inventory.DeleteProduct(id)
			fmt.Print("\033[H\033[2J")
			fmt.Println("Product Deleted Successfully..!")
		case 5:
			inventory.ListAllProduct()
			fmt.Printf("\n\n")
		case 6:
			return
		default:
			fmt.Print("\033[H\033[2J")
			fmt.Println("Please Enter a valid choice!!")

		}
	}
}
