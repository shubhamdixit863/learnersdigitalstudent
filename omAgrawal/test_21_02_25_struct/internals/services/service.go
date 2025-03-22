package services

import (
	"encoding/json"
	"fmt"
	"os"
	"test_21_02_25_struct/internals/models"
	"time"
)

const (
	ErrorFailedToExport  = "failed to create file: %w"
	FailedToExport       = "failed to marshal invoice to JSON: %w"
	ShoulGreaterThanZero = "should be Greater than zero"
)

type InvoiceGenerator interface {
	AddItem(name string, quantity int, price float64) error
	CalculateTotal() float64
	ExportToJSON(filename string) error
}
type StandardInvoice struct {
	invoice models.Invoice
}

// NewStandardInvoice creates a new StandardInvoice
func NewStandardInvoice(invoiceNumber, customerName string, purchaseDate time.Time) *StandardInvoice {
	return &StandardInvoice{
		invoice: models.Invoice{
			InvoiceNumber: invoiceNumber,
			CustomerName:  customerName,
			PurchaseDate:  purchaseDate,
			Items:         make([]models.Item, 0),
			Total:         0.0,
		},
	}
}

func (si *StandardInvoice) AddItem(name string, quantity int, price float64) error {
	if quantity <= 0 {
		return fmt.Errorf(ShoulGreaterThanZero)
	}
	if price <= 0 {
		return fmt.Errorf(ShoulGreaterThanZero)
	}

	subtotal := float64(quantity) * price
	item := models.Item{
		Name:     name,
		Quantity: quantity,
		Price:    price,
		Subtotal: subtotal,
	}
	si.invoice.Items = append(si.invoice.Items, item)
	return nil
}

func (si *StandardInvoice) CalculateTotal() float64 {
	var total float64
	for _, item := range si.invoice.Items {
		total += item.Subtotal
	}
	si.invoice.Total = total
	return total
}

func (si *StandardInvoice) ExportToJSON(filename string) error {
	si.CalculateTotal()
	jsonData, err := json.MarshalIndent(si.invoice, "", "  ")
	if err != nil {
		return fmt.Errorf(FailedToExport, err)
	}
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf(ErrorFailedToExport, err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf(ErrorFailedToExport, err)
	}
	return nil
}
func (si *StandardInvoice) GetInvoice() {
	fmt.Println(si.invoice)
}
