Develop an Inventory Management System in Go that enables users to:

1. Add new products to the inventory.

2. Update product details such as price and quantity.

3. Search for products by name or category.

4. Delete products from the inventory.

5. List all products with their details.



Requirements:

• Utilize structs to represent products.

• Employ slices to maintain a collection of products.

• Implement maps for efficient product search by unique identifiers.

• Define methods and functions to perform operations like adding, updating, searching, deleting, and listing products.

• Organize the code into multiple packages (inventory, product, main) to ensure modularity and maintainability.



Project Structure:

inventory-management/

├── main.go

├── inventory/

│   ├── inventory.go

├── product/

│   ├── product.go

Functional Specifications:

1. Product Management:

• Add Product: Allow users to add new products with details such as ID, name, category, price, and quantity.

• Update Product: Enable updating of product information, including price and quantity.

• Delete Product: Provide functionality to remove a product from the inventory.

2. Search Functionality:

• By Name: Search and retrieve products matching a given name.

• By Category: List all products under a specific category.

3. Inventory Listing:

• List All Products: Display all products with their details, including ID, name, category, price, and available quantity.