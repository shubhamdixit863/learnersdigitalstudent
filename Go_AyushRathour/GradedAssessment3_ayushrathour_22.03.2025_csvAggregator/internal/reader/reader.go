package reader

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(filePath string, chunkSize int) ([][][2]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	//fmt.Println(records)

	var chunks [][][2]string
	for i := 0; i < len(records); i += chunkSize {
		end := i + chunkSize
		if end > len(records) {
			end = len(records)
		}

		chunk := make([][2]string, end-i)
		for j, record := range records[i:end] {
			if len(record) < 2 {
				continue
			}
			chunk[j] = [2]string{record[0], record[1]}
		}
		chunks = append(chunks, chunk)
	}
	return chunks, nil
}
