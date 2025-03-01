package product

type Product struct {
	Id       int
	Name     string
	Category string
	Price    int
	Quantity int
}

func NewProduct(id int, name string, category string, price int, quantity int) Product {
	return Product{Id: id, Name: name, Category: category, Price: price, Quantity: quantity}
}
