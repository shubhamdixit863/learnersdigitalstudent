package product

type Product struct {
	Name             string
	Product_Id       string
	Product_Quantity int
	Product_catogery string
}

func StudentDetails(name string, id string, quantity int, catgory string) *Product {
	product := new(Product)
	product.Name = name
	product.Product_Id = id
	product.Product_Quantity = quantity
	product.Product_catogery = catgory
	return product
}

/*Output:
  Gabriel 1703065
  &{Gabriel 1703065}
*/
