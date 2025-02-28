package main

import (
	"fmt"
	"session7/inventry"
)

func main() {
	for {
		var choose int
		fmt.Println("Press")
		fmt.Println("1 for creating")
		fmt.Println("2 for List")
		fmt.Println("3 for catogery list")
		fmt.Println("4 for update")
		fmt.Println("5 for Delete")
		fmt.Println("6 for exit")
		fmt.Scanln(&choose)
		switch choose {
		case 1:
			fmt.Println("Enter name of product")
			var name string
			fmt.Scanln(&name)
			fmt.Println("Enter Id of product")
			var id string
			fmt.Scanln(&id)
			fmt.Println("Enter Qnt of product")
			var qnt int
			fmt.Scanln(&qnt)
			fmt.Println("Enter catogery of product")
			var ctg string
			fmt.Scanln(&ctg)
			inventry.Create_product(name, id, qnt, ctg)
		case 2:
			inventry.List()
		case 3:
			inventry.Catogery_List()
		case 4:
			fmt.Println("Enter name of product you want to update")
			var name string
			fmt.Scanln(&name)
			inventry.Update(name)
		case 5:
			fmt.Println("Enter name of product you want to delete")
			var name string
			fmt.Scanln(&name)
			inventry.Delete(name)
		case 6:
			return

		}
	}

}
