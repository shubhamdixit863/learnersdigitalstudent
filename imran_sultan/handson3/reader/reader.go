package reader

import (
	"bufio"
	"fmt"
	"os"
)

const errorMessage = "Error reading file: %v\n"

func ReadChunks(filePath string, chunkSize int, out chan<- []string) {
	defer close(out)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf(errorMessage, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) >= chunkSize {
			out <- lines
			lines = nil
		}
	}

	if len(lines) > 0 {
		out <- lines
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf(errorMessage, err)
	}
}
