package internal

import (
	"encoding/csv"
	"os"
	"strconv"
)


type Results struct {
	ProductRevenue           map[string]float64
	MostPurchasedProduct      string
	HighestSpendingCustomer   string
	CategoryRevenue           map[string]float64
}


func ProcessTransactions(filePath string) (Results, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Results{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return Results{}, err
	}

	
	productRevenue := make(map[string]float64)
	productQuantity := make(map[string]int)
	customerSpend := make(map[string]float64)
	categoryRevenue := make(map[string]float64)

	
	for _, record := range records[1:] {
		product := record[1]
		category := record[2]
		customerID := record[3]
		quantity, _ := strconv.Atoi(record[4])
		price, _ := strconv.ParseFloat(record[5], 64)
		total := float64(quantity) * price

		
		productRevenue[product] += total
		productQuantity[product] += quantity
		customerSpend[customerID] += total
		categoryRevenue[category] += total
	}

	
	mostPurchasedProduct := ""
	maxQuantity := 0
	for product, quantity := range productQuantity {
		if quantity > maxQuantity {
			maxQuantity = quantity
			mostPurchasedProduct = product
		}
	}

	
	highestSpendingCustomer := ""
	maxSpend := 0.0
	for customer, spend := range customerSpend {
		if spend > maxSpend {
			maxSpend = spend
			highestSpendingCustomer = customer
		}
	}

	return Results{
		ProductRevenue:         productRevenue,
		MostPurchasedProduct:   mostPurchasedProduct,
		HighestSpendingCustomer: highestSpendingCustomer,
		CategoryRevenue:        categoryRevenue,
	}, nil
}
