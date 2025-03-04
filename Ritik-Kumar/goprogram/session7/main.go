package main

import (
	"fmt"
	"session7/product"
)

func main() {
	var products []product.Inventory

	for {
		fmt.Println("\n Inventory Mangement System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Product")
		fmt.Println("3. Search Product")
		fmt.Println("4. Delete Product")
		fmt.Println("5. Exit")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var id,price, quantity int
			var names string
			fmt.Print("Enter ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Name: ")
			fmt.Scanln(&names)
			fmt.Print("Enter Price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter Quantity: ")
			fmt.Scanln(&quantity)
			newProduct := product.AddProduct(id, names, price, quantity)
			products = append(products, newProduct)
		case 2:
			var id int
			fmt.Println("Enter Product Id")
			fmt.Scanln(&id)
			product.UpdateProduct(products,id)

		case 3:
			var name string
			fmt.Println("Search by name")
			fmt.Scanln(&name)
			product.SearchProduct(products,name)

		case 4:
			var id int
			fmt.Print("Enter Product ID to delete: ")
			fmt.Scanln(&id)
			product.DeleteProduct(&products, id)
		
		case 5:
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid choice, try again!")
		}
	}
}






// Develop an Inventory Management System in Go that enables users to:

// 1. Add new products to the inventory.

// 2. Update product details such as price and quantity.

// 3. Search for products by name or category.

// 4. Delete products from the inventory.

// 5. List all products with their details.



// Requirements:

// • Utilize structs to represent products.

// • Employ slices to maintain a collection of products.

// • Implement maps for efficient product search by unique identifiers.

// • Define methods and functions to perform operations like adding, updating, searching, deleting, and listing products.

// • Organize the code into multiple packages (inventory, product, main) to ensure modularity and maintainability.



// Project Structure:

// inventory-management/

// ├── main.go

// ├── inventory/

// │   ├── inventory.go

// ├── product/

// │   ├── product.go

// Functional Specifications:

// 1. Product Management:

// • Add Product: Allow users to add new products with details such as ID, name, category, price, and quantity.

// • Update Product: Enable updating of product information, including price and quantity.

// • Delete Product: Provide functionality to remove a product from the inventory.

// 2. Search Functionality:

// • By Name: Search and retrieve products matching a given name.

// • By Category: List all products under a specific category.

// 3. Inventory Listing:

// • List All Products: Display all products with their details, including ID, name, category, price, and available quantity.

