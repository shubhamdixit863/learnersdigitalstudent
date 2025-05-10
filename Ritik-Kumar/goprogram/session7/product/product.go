package product

import "fmt"

type Inventory struct {
	Productid int
	Name      string
	Price     int
	Quantity  int
}

func AddProduct(ids int, name string, price int, quantity int) Inventory {
	return Inventory{
		Productid: ids,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
	}
}

func UpdateProduct(products []Inventory, id int) {
	for i, p := range products {
		if p.Productid == id {
			fmt.Println("Enter New Price:")
			var price int
			fmt.Scanln(&price)

			fmt.Println("Enter New Quantity:")
			var quantity int
			fmt.Scanln(&quantity)

			products[i].Price = price
			products[i].Quantity = quantity
			fmt.Println("Product updated successfully!")
			return
		}
	}
	fmt.Println("Item not present")
}


func SearchProduct(products []Inventory,name string){
	for _,p:= range products{
		if(p.Name==name){
			fmt.Println(p.Productid,p.Name,p.Price,p.Quantity)
		}else{
			fmt.Println("Item not found")
		}
	}
}
func DeleteProduct(products *[]Inventory, id int) {
	index := -1

	for i, p := range *products {
		if p.Productid == id {
			index = i
			break
		}
	}
	if index != -1 {
		*products = append((*products)[:index], (*products)[index+1:]...)
		fmt.Println("Product deleted successfully!")
	} else {
		fmt.Println("Product not found!")
	}
}
