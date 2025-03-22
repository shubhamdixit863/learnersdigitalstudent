package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sehandson1/services"
	"strconv"
)

func main() {
	file, err := os.Open("data/stock_prices.csv")
	myMap := make(map[string][]float64)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	records = records[1:]

	for _, eachrecord := range records {
		n, _ := strconv.ParseFloat(eachrecord[2], 64)
		myMap[eachrecord[1]] = append(myMap[eachrecord[1]], n)

	}
	services.CalculateAvg(myMap)
	services.MInMax(myMap)
	services.MostVolatile(myMap)
}
