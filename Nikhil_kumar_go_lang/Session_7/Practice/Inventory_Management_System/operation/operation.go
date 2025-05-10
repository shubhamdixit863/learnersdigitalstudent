package operation

import (
	"Inventory_Management_System/inventory"
	"Inventory_Management_System/product"
	"fmt"
)

func AddProduct(inv *inventory.Inventory) {
	var newProduct product.Product
	var input string

	fmt.Print("Enter product ID: ")
	fmt.Scan(&newProduct.ID)

	if _, exists := inv.GetProductByID(newProduct.ID); exists {
		fmt.Println("Product with this ID already exists!")
		return
	}

	fmt.Scanln()

	fmt.Print("Enter product name: ")
	fmt.Scanln(&input)
	newProduct.Name = input

	fmt.Print("Enter product category: ")
	fmt.Scanln(&input)
	newProduct.Category = input

	fmt.Print("Enter product price: ")
	fmt.Scan(&newProduct.Price)

	fmt.Print("Enter product quantity: ")
	fmt.Scan(&newProduct.Quantity)

	inv.AddProduct(newProduct)
	fmt.Println("Product added successfully!")
}

func UpdateProduct(inv *inventory.Inventory) {
	var id string
	var price float64
	var quantity int
	// var input string

	fmt.Print("Enter product ID to update: ")
	fmt.Scan(&id)

	prod, exists := inv.GetProductByID(id)
	if !exists {
		fmt.Println("Product not found!")
		return
	}

	fmt.Printf("Current product details: %+v\n", prod)

	fmt.Print("Enter new price : ")
	fmt.Scan(&price)
	if price != 0 {
		prod.Price = price
	}

	fmt.Print("Enter new quantity : ")
	fmt.Scan(&quantity)
	if quantity != -1 {
		prod.Quantity = quantity
	}

	inv.UpdateProduct(prod)
	fmt.Println("Product updated successfully!")
}

func SearchProducts(inv *inventory.Inventory) {
	var searchTerm string

	fmt.Println("Search by name")
	fmt.Scanln()

	fmt.Print("Enter product name to search: ")
	fmt.Scanln(&searchTerm)
	products := inv.SearchByName(searchTerm)
	DisplayProducts(products)

}

func DeleteProduct(inv *inventory.Inventory) {
	var id string
	fmt.Print("Enter product ID to delete: ")
	fmt.Scan(&id)

	if inv.DeleteProduct(id) {
		fmt.Println("Product deleted successfully!")
	} else {
		fmt.Println("Product not found!")
	}
}

func ListAllProducts(inv *inventory.Inventory) {
	products := inv.ListAllProducts()
	DisplayProducts(products)
}

func DisplayProducts(products []product.Product) {
	if len(products) == 0 {
		fmt.Println("No products found.")
		return
	}

	fmt.Println("\n===== Product List =====")
	fmt.Printf("%s %s %s %s %s\n", "ID", "Name", "Category", "Price", "Quantity")

	for _, p := range products {
		fmt.Printf("%s %s %s %.2f %d\n",
			p.ID, p.Name, p.Category, p.Price, p.Quantity)
	}
}


