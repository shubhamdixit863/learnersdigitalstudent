package main

import (
	"log"
	"test_21_02_25_struct/internals/services"
	"time"
)

const (
	errorAdding    = "Error adding item: %v\n"
	TotalS         = "Invoice total: %.2f\n"
	ErrorExporting = "Error exporting invoice: %v\n"
	SuccessExport  = "Invoice successfully exported to %s\n"
)

const (
	invoiceName  = "invoice1"
	CustomerName = "Om Agrawal"
)

const jsonname = "invoice1.json"

func main() {
	// Create a new invoice
	purchaseDate := time.Now()
	invoice := services.NewStandardInvoice(invoiceName, CustomerName, purchaseDate)

	err := invoice.AddItem("Laptop", 1, 1299.99)
	if err != nil {
		log.Fatalf(errorAdding, err)
		return
	}

	err2 := invoice.AddItem("Mouse", 2, 25.50)
	if err2 != nil {
		log.Fatalf(errorAdding, err)
		return
	}

	err3 := invoice.AddItem("Keyboard", 1, 85.75)
	if err3 != nil {
		log.Fatalf(errorAdding, err)
		return
	}

	total := invoice.CalculateTotal()
	log.Printf(TotalS, total)

	err4 := invoice.ExportToJSON(jsonname)
	if err4 != nil {
		log.Fatalf(ErrorExporting, err)
		return
	}
	log.Printf(SuccessExport, jsonname)

}
