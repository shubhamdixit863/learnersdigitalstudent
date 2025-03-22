package models

import (
	"time"
)

type Item struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`
}

type Invoice struct {
	InvoiceNumber string    `json:"invoice_number"`
	CustomerName  string    `json:"customer_name"`
	PurchaseDate  time.Time `json:"purchase_date"`
	Items         []Item    `json:"items"`
	Total         float64   `json:"total"`
}
