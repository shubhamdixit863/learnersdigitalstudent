package services

import (
	"encoding/csv"
	"log"
	"os"
)

func FileReader() [][]string {
	file, err := os.Open("ecommerce_transactions.csv")
	if err != nil {
		log.Println(err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)

	}
	return records

}
