package services

import (
	"bufio"
	"fileProcessor/internal/models"
	"fmt"
	"os"
	"sync"
)

const (
	err21 ="Error opening file:"
	err22 = "Error reading file:"
)

func ProcessFiles(files []string) {
	lineChannel := make(chan models.LineWithMetadata, 100)
	var wg sync.WaitGroup


	for _, file := range files {
		wg.Add(1)
		go readFile(file, lineChannel, &wg)
	}


	go func() {
		wg.Wait()
		close(lineChannel)
	}()

	
	ProcessLines(lineChannel)
}

func readFile(filePath string, lineChannel chan<- models.LineWithMetadata, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err21, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		lineChannel <- models.LineWithMetadata{
			FileName: filePath,
			LineNum:  lineNum,
			Line:     scanner.Text(),
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err22, err)
	}
}
