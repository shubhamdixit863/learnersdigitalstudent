package services

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const(
	filteredline ="Filtered Line:"
	totalwords = "Total Words:"
	posturl ="https://httpbin.org/post"
	json ="application/json"
	apierror ="API Error:"
	apiresponse ="API Response:"
	failedmessage="Failed to send after retries:"
	retrymessage="Retry %d failed: %v\n"

)

func FilterLines(linesChan <-chan string, keyword string) {
	for line := range linesChan {
		if strings.Contains(line, keyword) {
			fmt.Println(filteredline, line)
		}
	}
}

//reciever
func CountWords(linesChan <-chan string) {
	totalWords := 0
	for line := range linesChan {
		totalWords += len(strings.Fields(line))
	}
	fmt.Println(totalwords, totalWords)
}


func SendToAPI(linesChan <-chan string) {
	for line := range linesChan {
		resp, err := http.Post(posturl, json, bytes.NewBuffer([]byte(line)))
		if err != nil {
			fmt.Println(apierror, err)
			continue
		}
		fmt.Println(apiresponse, resp.Status)
		resp.Body.Close()
	}
}


func SendToAPIWithRetry(linesChan <-chan string) {
	for line := range linesChan {
		retryAPI(line, 3)
	}
}

func retryAPI(line string, maxRetries int) {
	for i := 0; i < maxRetries; i++ {
		resp, err := http.Post(posturl, json, bytes.NewBuffer([]byte(line)))
		if err == nil {
			fmt.Println(apiresponse, resp.Status)
			resp.Body.Close()
			return
		}
		fmt.Printf(retrymessage, i+1, err)
		time.Sleep(time.Second * 2)
	}
	fmt.Println(failedmessage, line)
}
