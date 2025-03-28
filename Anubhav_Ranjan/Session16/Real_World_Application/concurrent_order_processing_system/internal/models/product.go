package models

import "sync"

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
	mu    sync.Mutex
}
