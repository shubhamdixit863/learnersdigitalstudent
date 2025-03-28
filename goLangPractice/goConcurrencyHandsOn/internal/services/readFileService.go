package services

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func ReadLines(filePath string, lines chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lines <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
