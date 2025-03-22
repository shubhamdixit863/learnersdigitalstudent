package inventry

import (
	"fmt"
	"session7/product"
)

var arr []*product.Product
var mp = make(map[string]string)

func Create_product(name string, id string, quantity int, catogery string) {
	item := product.StudentDetails(name, id, quantity, catogery)
	arr = append(arr, item)
	mp[item.Name] = item.Product_catogery
}
func Update(name string) {
	for i := 0; i < len(arr); i++ {
		if arr[i].Name == name {
			fmt.Println(arr[i])
			var catogery string
			fmt.Println("Enter the updated ")
			fmt.Scanln(&catogery)
			arr[i].Product_catogery = catogery
			break

		}
	}
}
func List() {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
func Catogery_List() {
	for key, value := range mp {
		fmt.Println(key, value)
	}
}
func Delete(name string) {
	for i := 0; i < len(arr); i++ {
		if arr[i].Name == name {

			arr = append(arr[:i], arr[i+1:]...)

			break

		}
	}
}
