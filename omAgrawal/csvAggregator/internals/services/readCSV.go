package services

import (
	"csvAggregator/internals/models"
	"encoding/csv"
	"log"
	"os"
)

func FileReader() [][]string {
	file, err := os.Open(models.FileName)
	if err != nil {
		log.Println(err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(models.FileReadError, err)

	}
	log.Printf(models.FileReadSuccesfull)
	return records

}
