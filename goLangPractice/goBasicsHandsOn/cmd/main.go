package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	revenuePerProduct := make(map[string]float64)
	quantityPerProduct := make(map[string]int)
	revenuePerCategory := make(map[string]float64)
	revenuePerCustomer := make(map[string]float64)

	file, err := os.Open("ecommerce_transactions.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	//fmt.Println("Data:", data)

	data = data[1:]
	for _, row := range data {
		customerID := row[1]
		product := row[2]
		category := row[3]
		quantity, _ := strconv.Atoi(row[4])
		price, _ := strconv.ParseFloat(row[5], 64)

		totalPrice := float64(quantity) * price

		revenuePerProduct[product] += totalPrice
		quantityPerProduct[product] += quantity
		revenuePerCategory[category] += totalPrice
		revenuePerCustomer[customerID] += totalPrice
	}

	//TASK - 1
	fmt.Println("Revenue per product:")
	totalRevenuePerProduct(revenuePerProduct)
	fmt.Println("-------------------------------------------------------------------------------------")

	//TASK - 2
	mostPurchasedProduct, maxQuantity := mostPurchasedProduct(quantityPerProduct)
	fmt.Printf("Most Purchased Product is: %s, total purchased quantity is: %d\n", mostPurchasedProduct, maxQuantity)
	fmt.Println("-------------------------------------------------------------------------------------")

	//TASK - 3
	highestSpendingCustomer, maxSpending := highestSpendingCustomer(revenuePerCustomer)
	fmt.Printf("CustomerId of customer spending most is: %s, total spending is: %.2f\n", highestSpendingCustomer, maxSpending)
	fmt.Println("-------------------------------------------------------------------------------------")

	//TASK - 4
	fmt.Println("Revenue per category:")
	revenuePerProductCategory(revenuePerCategory)

}

func totalRevenuePerProduct(revenuePerProduct map[string]float64) {
	for k, v := range revenuePerProduct {
		fmt.Println(k, " - ", v)
	}
}

func mostPurchasedProduct(quantityPerProduct map[string]int) (string, int) {
	maxQuantity := 0
	mostPurchasedProduct := ""

	for k, v := range quantityPerProduct {
		if v > maxQuantity {
			maxQuantity = v
			mostPurchasedProduct = k
		}
	}

	return mostPurchasedProduct, maxQuantity
}

func highestSpendingCustomer(revenuePerCustomer map[string]float64) (string, float64) {
	maxSpending := 0.0
	customerId := ""

	for k, v := range revenuePerCustomer {
		if v > maxSpending {
			maxSpending = v
			customerId = k
		}
	}

	return customerId, maxSpending
}

func revenuePerProductCategory(revenuePerCategory map[string]float64) {
	for k, v := range revenuePerCategory {
		fmt.Println(k, " - ", v)
	}
}
