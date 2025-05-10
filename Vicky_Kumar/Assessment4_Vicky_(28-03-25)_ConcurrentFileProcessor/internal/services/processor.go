package services

import (
	"bytes"
	"concurrent_file_processor/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type Processor struct {
	mode    string
	keyword string
}

func NewProcessor(mode string, keyword string) *Processor {
	return &Processor{mode: mode,
		keyword: keyword}
}

func (p *Processor) Process(linesChan chan string, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	switch p.mode {
	case "filter":
		p.processFilter(linesChan, resultChan)
	case "wordcount":
		p.processWordCount(linesChan, resultChan)
	case "api":
		p.processAPICall(linesChan, resultChan)
	case "retry":
		p.processAPIRetry(linesChan, resultChan)
	default:
		log.Println("Unknown mode. No processing done.")
	}
}

func (p *Processor) processFilter(linesChan chan string, resultChan chan string) {

	for line := range linesChan {
		if strings.Contains(line, p.keyword) {
			resultChan <- line
		}
	}
	close(resultChan)
}

func (p *Processor) processWordCount(linesChan chan string, resultChan chan string) {
	totalWords := 0
	for line := range linesChan {
		words := strings.Split(line, " ")
		totalWords += len(words)
	}
	resultChan <- fmt.Sprintf(utils.WordCount, totalWords)
	close(resultChan)
}

func (p *Processor) processAPICall(linesChan chan string, resultChan chan string) {
	var wg sync.WaitGroup
	for line := range linesChan {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			resp, err := callAPI(l)
			if err != nil {
				resultChan <- fmt.Sprintf(utils.ErrApiProcessing, l, err)
			} else {
				resultChan <- fmt.Sprintf(utils.ApiResponse, l, resp)
			}
		}(line)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
}

func (p *Processor) processAPIRetry(linesChan chan string, resultChan chan string) {
	var wg sync.WaitGroup
	for line := range linesChan {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				resp, err := callAPI(l)
				if err != nil {
					log.Printf(utils.ErrApiRetry, l)
				} else {
					resultChan <- fmt.Sprintf(utils.ApiResponse, l, resp)
					break
				}

			}
		}(line)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

}

func callAPI(line string) (string, error) {
	url := "https://httpbin.org/post"
	payload := map[string]string{"line": line}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	return resp.Status, nil
}
