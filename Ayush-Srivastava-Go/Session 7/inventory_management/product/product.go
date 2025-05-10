package product

import "fmt"

type Product struct {
	ID       string
	Name     string
	Category string
	Price    float64
	Quantity int
}

func (p *Product) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Category: %s, Price: %.2f, Quantity: %d", p.ID, p.Name, p.Category, p.Price, p.Quantity)
}