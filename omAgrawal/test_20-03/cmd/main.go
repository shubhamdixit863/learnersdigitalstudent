package main

import (
	"fmt"
	"test_20-03/internals/services"
)

func main() {
	recodes := services.FileReader()
	NewRecord := recodes[1:]

	fmt.Println("the data")
	for _, NewRecord := range NewRecord {
		fmt.Println(NewRecord)
	}
	fmt.Println("")

	services.TotalRevenue(NewRecord)
	services.MostPurchasedProduct(NewRecord)
	services.HighestSpendingCustomer(NewRecord)
	services.TotalRevenuePerCategory(NewRecord)

}
