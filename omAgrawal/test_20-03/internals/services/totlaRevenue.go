package services

import (
	"fmt"
	"strconv"
)

func TotalRevenue(records [][]string) {

	TotalRevenuePerProduct := make(map[string]float64)

	for _, record := range records {
		price, _ := strconv.ParseFloat(record[5], 64)
		qnty, _ := strconv.ParseFloat(record[4], 64)
		TotalRevenuePerProduct[record[2]] = price*qnty + TotalRevenuePerProduct[record[2]]
	}
	//fmt.Println(TotalRevenuePerProduct)
	fmt.Println("Total Revenue Per Product is")
	fmt.Println("Product  TotalRevenue")
	for v, q := range TotalRevenuePerProduct {
		fmt.Println(v, q)
	}
	fmt.Println("")
}

func TotalRevenuePerCategory(records [][]string) {

	TotalRevenuePerProduct := make(map[string]float64)

	for _, record := range records {
		price, _ := strconv.ParseFloat(record[5], 64)
		qnty, _ := strconv.ParseFloat(record[4], 64)
		TotalRevenuePerProduct[record[3]] = price*qnty + TotalRevenuePerProduct[record[3]]
	}
	//fmt.Println(TotalRevenuePerProduct)
	fmt.Println("Total Revenue Per Category is")
	fmt.Println("Category  TotalRevenue")
	for v, q := range TotalRevenuePerProduct {
		fmt.Println(v, q)
	}
	fmt.Println("")
}

func HighestSpendingCustomer(records [][]string) {
	Map := make(map[string]float64)
	for _, record := range records {
		price, _ := strconv.ParseFloat(record[5], 64)
		qnty, _ := strconv.ParseFloat(record[4], 64)
		Map[record[1]] += price * qnty
	}

	//fmt.Println("Customer and the total Spend by them")
	//fmt.Println(Map)
	fmt.Println("")
	fmt.Println("Total Revenue Per Customer is")
	fmt.Println("Customer  TotalRevenue")
	for v, q := range Map {
		fmt.Println(v, "   ", q)
	}
	fmt.Println("")
	fmt.Println("")

}
