package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Mode represents the processing mode.
type Mode int

const (
	LineFilter Mode = iota
	WordCount
	APICall
	RetryableAPICall
)

const (
	filterKeyword = "error"
	apiURL        = "https://httpbin.org/post"
	maxRetries    = 3
	apiTimeout    = 2 * time.Second
)

func main() {
	// Parse command-line flags.
	dirPath := flag.String("dir", "./textfiles", "Directory path containing .txt files")
	modeStr := flag.String("mode", "linefilter", "Processing mode: linefilter, wordcount, apicall, retryapi")
	flag.Parse()

	if *dirPath == "" {
		fmt.Println("Please specify a directory using -dir flag.")
		os.Exit(1)
	}

	mode := parseMode(*modeStr)

	// Create a central channel for fan-in pattern.
	linesChan := make(chan string, 100)
	var wg sync.WaitGroup

	// Fan-Out: Walk through the directory and start a goroutine for each .txt file.
	err := filepath.WalkDir(*dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".txt") {
			return nil
		}
		wg.Add(1)
		go readFileLines(path, linesChan, &wg)
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	// Close the lines channel after all file readers are done.
	go func() {
		wg.Wait()
		close(linesChan)
	}()

	// Fan-In: Process the incoming lines based on the chosen mode.
	processLines(linesChan, mode)
}

// parseMode converts a string into a Mode constant.
func parseMode(modeStr string) Mode {
	switch strings.ToLower(modeStr) {
	case "linefilter":
		return LineFilter
	case "wordcount":
		return WordCount
	case "apicall":
		return APICall
	case "retryapi":
		return RetryableAPICall
	default:
		return LineFilter
	}
}

// readFileLines reads a file line by line and sends each line to the provided channel.
func readFileLines(path string, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file %s: %v\n", path, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", path, err)
	}
}

// processLines consumes lines from the channel and processes them based on the selected mode.
func processLines(in <-chan string, mode Mode) {
	switch mode {
	case LineFilter:
		for line := range in {
			if strings.Contains(line, filterKeyword) {
				fmt.Println("[Filtered]", line)
			}
		}
	case WordCount:
		totalWords := 0
		for line := range in {
			totalWords += len(strings.Fields(line))
		}
		fmt.Printf("Total word count: %d\n", totalWords)
	case APICall:
		for line := range in {
			callAPI(line)
		}
	case RetryableAPICall:
		for line := range in {
			callWithRetry(line, maxRetries)
		}
	}
}

// callAPI sends the given data to the API endpoint and prints the response status.
func callAPI(data string) {
	resp, err := http.Post(apiURL, "text/plain", bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Printf("API call error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("API status: %s\n", resp.Status)
}

// callWithRetry sends data to the API endpoint, retrying upon failure with exponential backoff.
func callWithRetry(data string, retries int) {
	// var err error
	for attempt := 0; attempt <= retries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), apiTimeout)
		req, reqErr := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewBuffer([]byte(data)))
		if reqErr != nil {
			cancel()
			fmt.Printf("Request creation error: %v\n", reqErr)
			return
		}
		req.Header.Set("Content-Type", "text/plain")

		resp, err := http.DefaultClient.Do(req)
		cancel() // Ensure cancellation after the request is done.

		if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			fmt.Printf("API call succeeded after %d attempt(s)\n", attempt+1)
			return
		}
		if resp != nil {
			resp.Body.Close()
		}
		// Exponential backoff before next attempt.
		time.Sleep(time.Duration(attempt+1) * time.Second)
	}
	fmt.Printf("API call failed after %d attempts. Data: %s\n", retries+1, data)
}
