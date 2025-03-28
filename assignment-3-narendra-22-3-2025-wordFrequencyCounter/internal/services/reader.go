package services

import (
	"bufio"
	"fmt"
	"os"
)


func Reader(file *os.File, chunkChan chan<- string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		chunkChan <- scanner.Text() 
	}

	
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	close(chunkChan) 
}




