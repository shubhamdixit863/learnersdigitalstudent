package services

import (
	"bufio"
	"bytes"
	"fmt"
	"wordfrequencycounter/internal/models"
)

const scanerror ="error while scanning words: %w"

func WordScanner(data []byte) ([]string, error) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		word := scanner.Text()
		cleanWord := models.CleanText(word)
		if cleanWord != "" {
			words = append(words, cleanWord)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf(scanerror, err)
	}

	return words, nil
}
