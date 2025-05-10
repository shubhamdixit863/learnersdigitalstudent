package services

import "sync"

type Inventory struct {
	stock int
	mu sync.Mutex
}

func NewInventory(initialStock int) *Inventory {
	return &Inventory{stock: initialStock}
}

func (inv *Inventory) DeductStock() bool {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	if inv.stock > 0 {
		inv.stock--
		return true
	}
	return false
}

func (inv *Inventory) GetStock() int {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	return inv.stock
}