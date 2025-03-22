package services

import (
	"bufio"
	"filemonitor/internal/utils"
	"fmt"
	"os"
)

func ProcessFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf(utils.FileOpenErr, err)
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf(utils.FileReadErr, err)
	}

	fmt.Printf("File: %s, Lines: %d\n", filePath, lineCount)
	return nil
}