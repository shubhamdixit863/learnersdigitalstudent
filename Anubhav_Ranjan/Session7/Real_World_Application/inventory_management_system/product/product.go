package product

type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
	Quantity int
}

func NewProduct(id int, name string, category string, price float64, quantity int) *Product {
	return &Product{ID: id, Name: name, Category: category, Price: price, Quantity: quantity}
}
