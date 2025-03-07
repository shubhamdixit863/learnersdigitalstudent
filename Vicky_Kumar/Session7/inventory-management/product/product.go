package product

type Product struct {
	Id       int
	Name     string
	Category string
	Price    float64
	Quantity int
}

var Products []Product
