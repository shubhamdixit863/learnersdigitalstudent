package models

import "time"

type OrderItem struct {
	ProductID int
	Quantity  int
}

type Order struct {
	ID        int
	Items     []OrderItem
	Status    string
	Timestamp time.Time
}
