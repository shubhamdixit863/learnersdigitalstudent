# TITLE :-  INVOICE GENERATOR

### **Features**

* Add customer information and purchase date
* Create invoices with multiple items
* Automatic calculation of item subtotals and invoice totals
* JSON export with proper formatting

### HOW to Build
go build .\cmd\main.go

### input 1
`invoice := services.NewStandardInvoice("invoice-1", "John Doe", purchaseDate)`
invoice number
Name of the customer

### input 2 
`err := invoice.AddItem("Laptop", 1, 1299.99)`
Name Of the product
Quantity 
price in float

### output 1
`total := invoice.CalculateTotal()
fmt.Printf(TotalS, total)`

Prints the total of all the products

### output 2

Creates "invoice.json" file for the user




