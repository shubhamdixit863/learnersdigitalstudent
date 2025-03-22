package main

import (
	"E-commerceAnalyticsDashboard/internal"
	"fmt"
)

func main() {
	filePath := "D:/NexturnAssessments/Assessment01/E-commerceAnalyticsDashboard/data/data.csv"

	results, err := internal.ProcessTransactions(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Total Revenue Per Product:", results.ProductRevenue)
	fmt.Println("Most Purchased Product:", results.MostPurchasedProduct)
	fmt.Println("Highest Spending Customer:", results.HighestSpendingCustomer)
	fmt.Println("Aggregated Revenue Per Category:", results.CategoryRevenue)
}
