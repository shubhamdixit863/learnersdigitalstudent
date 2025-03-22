package orders

import "sync"

type Inventory struct {
	stock map[string]int
	mutex sync.Mutex
}

// NewInventory initializes an inventory
func NewInventory() *Inventory {
	return &Inventory{stock: make(map[string]int)}
}

// SetStock initializes stock for a product
func (i *Inventory) SetStock(product string, quantity int) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.stock[product] = quantity
}

// GetStock retrieves stock quantity of a product
func (i *Inventory) GetStock(product string) int {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	return i.stock[product]
}

// UpdateStock updates stock if enough quantity is available
func (i *Inventory) UpdateStock(product string, quantity int) bool {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.stock[product] >= quantity {
		i.stock[product] -= quantity
		return true
	}
	return false
}
