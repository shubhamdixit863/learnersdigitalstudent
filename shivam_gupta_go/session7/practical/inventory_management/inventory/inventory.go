package inventory

import (
	"fmt"
	"session7/practical/inventory_management/product"
)

var Inventory_array []product.Product
var ProductID int = 1

func Add() {
	var Product2 product.Product
	fmt.Println("add fucntion")
	fmt.Println("enter name : ")

	var name string
	fmt.Scanln(&name)

	fmt.Println("enter category")
	var c string
	fmt.Scanln(&c)

	fmt.Println("enter price ")
	var p float64

	fmt.Scanln(&p)

	fmt.Println("enter qauntity")
	var q int
	fmt.Scanln(&q)
	Product2.ID = ProductID
	Product2.Name = name
	Product2.Category = c
	Product2.Price = p
	Product2.Quantity += q

	ProductID++

	Inventory_array = append(Inventory_array, Product2)
	fmt.Println("added successfully")

}

func Update() {
	fmt.Println("enter name ")
	var name string
	fmt.Scanln(&name)
	for i, _ := range Inventory_array {
		if Inventory_array[i].Name == name {
			fmt.Println("enter price for update")
			var p float64
			fmt.Scanln(&p)
			Inventory_array[i].Price = p
			fmt.Println("enter quantity")
			var q int
			fmt.Scanln(&q)
			Inventory_array[i].Quantity = q
			return
		}
	}
	fmt.Println("not found")
}

func Delete() {
	fmt.Println("enter name ")
	var name string
	fmt.Scanln(&name)
	for i, _ := range Inventory_array {
		if Inventory_array[i].Name == name {
			Inventory_array = append(Inventory_array[:i], Inventory_array[i+1:]...)
			fmt.Println("deteted")
			return
		}
	}
	fmt.Println("not found")
}

func List() {
	fmt.Println(Inventory_array)
}
