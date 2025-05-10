package utils

import (
	"bytes"
	"net/http"
	"time"
)

// SendToAPIWithRetry sends each line to an API with retries
// func SendToAPIWithRetry(lines <-chan string) {
// 	for line := range lines {
// 		resp, err := retryHTTPRequest(line, 3)
// 		if err != nil {
// 			fmt.Println("API Call Failed After Retries:", err)
// 		} else {
// 			fmt.Println("API Response:", resp.Status)
// 		}
// 	}
// }

// retryHTTPRequest retries sending data up to maxRetries
func retryHTTPRequest(data string, maxRetries int) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; i < maxRetries; i++ {
		resp, err = sendHTTPRequest(data)
		if err == nil {
			return resp, nil
		}
		time.Sleep(time.Duration(i+1) * time.Second) // Exponential backoff
	}
	return nil, err
}

// sendHTTPRequest sends data to an external API
func sendHTTPRequest(data string) (*http.Response, error) {
	reqBody := bytes.NewBuffer([]byte(data))
	resp, err := http.Post("https://httpbin.org/post", "text/plain", reqBody)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
