package apiHandler

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

// CallAPI function  Calls API for each line
func CallAPI(lineChan <-chan string) map[string]string {
	results := make(map[string]string)
	for line := range lineChan {
		status := sendRequest(line)
		results[line] = status
	}
	return results
}

// CallAPIWithRetry function Calls API with retries (up to 3 times with backoff)
func CallAPIWithRetry(lineChan <-chan string) map[string]string {
	results := make(map[string]string)
	for line := range lineChan {
		status := sendRequestWithRetry(line, 3)
		results[line] = status
	}
	return results
}

// Sends a single HTTP POST request
func sendRequest(line string) string {
	url := "https://httpbin.org/post"
	reqBody := []byte(fmt.Sprintf(`{"data":"%s"}`, line))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "Failed"
	}
	defer resp.Body.Close()
	return resp.Status
}

// Sends an API request with retry logic
func sendRequestWithRetry(line string, retries int) string {
	for i := 0; i < retries; i++ {
		status := sendRequest(line)
		if status != "Failed" {
			return status
		}
		time.Sleep(time.Duration(2*i) * time.Second) // Exponential backoff
	}
	return "Failed after retries"
}
