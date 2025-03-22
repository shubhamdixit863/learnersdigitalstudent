package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Configurable parameters
const (
	chunkSize = 2          // Number of lines per chunk
	fileName  = "book.txt" // Input file name
)

// readChunks reads the file and sends chunks of lines to the channel
func readChunks(filePath string, chunkSize int, out chan<- []string) error {
	file, err := os.Open(filePath)
	if err != nil {
		close(out)
		return err
	}
	defer file.Close()
	defer close(out) // Ensures channel is closed properly

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

	return scanner.Err()
}

// countWords counts word frequencies in a given chunk
func countWords(lines []string, out chan<- map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()

	wordFreq := make(map[string]int)
	for _, line := range lines {
		words := strings.Fields(strings.ToLower(line))
		for _, word := range words {
			word = strings.Trim(word, ".,!?;:\"'()[]{}")
			wordFreq[word]++
		}
	}
	out <- wordFreq
}

// mergeResults aggregates word frequencies from all goroutines
func mergeResults(in <-chan map[string]int) map[string]int {
	finalFreq := make(map[string]int)

	for freq := range in {
		for word, count := range freq {
			finalFreq[word] += count
		}
	}

	return finalFreq
}

func main() {
	chunks := make(chan []string)
	results := make(chan map[string]int)

	var wg sync.WaitGroup

	// Launch goroutine to read chunks
	if err := readChunks(fileName, chunkSize, chunks); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Process chunks concurrently
	for chunk := range chunks {
		wg.Add(1)
		go countWords(chunk, results, &wg)
	}

	// Close results channel after all goroutines finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Merge results
	wordCounts := mergeResults(results)

	// Print word frequencies
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
