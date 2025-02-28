package main

import (
	"fmt"
	"inventory_management/inventory"
	"inventory_management/product"
)

func main() {
	inv := inventory.NewInventory()

	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Product")
		fmt.Println("3. Delete Product")
		fmt.Println("4. Search by Name")
		fmt.Println("5. Search by Category")
		fmt.Println("6. List All Products")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, name, category string
			var price float64
			var quantity int

			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Category: ")
			fmt.Scan(&category)
			fmt.Print("Enter Price: ")
			fmt.Scan(&price)
			fmt.Print("Enter Quantity: ")
			fmt.Scan(&quantity)

			inv.AddProduct(product.Product{ID: id, Name: name, Category: category, Price: price, Quantity: quantity})

		case 2:
			var id string
			var price float64
			var quantity int

			fmt.Print("Enter Product ID to Update: ")
			fmt.Scan(&id)
			fmt.Print("Enter New Price: ")
			fmt.Scan(&price)
			fmt.Print("Enter New Quantity: ")
			fmt.Scan(&quantity)

			inv.UpdateProduct(id, price, quantity)

		case 3:
			var id string
			fmt.Print("Enter Product ID to Delete: ")
			fmt.Scan(&id)
			inv.DeleteProduct(id)

		case 4:
			var name string
			fmt.Print("Enter Product Name to Search: ")
			fmt.Scan(&name)
			results := inv.SearchByName(name)
			if len(results) > 0 {
				fmt.Println("\n--- Search Results ---")
				for _, prod := range results {
					fmt.Printf("ID: %s | Name: %s | Category: %s | Price: %.2f | Quantity: %d\n",
						prod.ID, prod.Name, prod.Category, prod.Price, prod.Quantity)
				}
			} else {
				fmt.Println("No products found!")
			}

		case 5:
			var category string
			fmt.Print("Enter Category to Search: ")
			fmt.Scan(&category)
			results := inv.SearchByCategory(category)
			if len(results) > 0 {
				fmt.Println("\n--- Search Results ---")
				for _, prod := range results {
					fmt.Printf("ID: %s | Name: %s | Category: %s | Price: %.2f | Quantity: %d\n",
						prod.ID, prod.Name, prod.Category, prod.Price, prod.Quantity)
				}
			} else {
				fmt.Println("No products found!")
			}

		case 6:
			inv.ListAllProducts()

		case 7:
			fmt.Println("Exiting the system. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
