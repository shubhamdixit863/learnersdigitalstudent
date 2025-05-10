package processor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

// RetryableAPICall sends lines to an API and retries on failure
func RetryableAPICall(linesCh <-chan string) {
	for line := range linesCh {
		success := false
		for attempt := 1; attempt <= 3; attempt++ {
			resp, err := http.Post("https://httpbin.org/post", "text/plain", bytes.NewBuffer([]byte(line)))
			if err != nil {
				fmt.Printf("Attempt %d: API Request Failed: %v\n", attempt, err)
				time.Sleep(time.Duration(math.Pow(2, float64(attempt))) * time.Second) // Exponential backoff
				continue
			}
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("API Response:", string(body))
			resp.Body.Close()
			success = true
			break
		}
		if !success {
			fmt.Println("Failed to send line after 3 attempts:", line)
		}
	}
}
