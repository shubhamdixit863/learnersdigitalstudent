package services

import (
    "fmt"
    "strings"
    "time"
)

func ProcessLines_Filter(in <-chan string, keyword string) {
    fmt.Println(" Filtered Lines:")
    for line := range in {
        if strings.Contains(line, keyword) {
            fmt.Println(line)
        }
    }
}

func ProcessLines_WordCount(in <-chan string) {
    count := 0
    for line := range in {
        count += len(strings.Fields(line))
    }
    fmt.Println(" Total Word Count:", count)
}

func ProcessLines_APICall(in <-chan string) {
    for line := range in {
        fmt.Printf("API called with line: %s \n", line)
    }
}

func ProcessLines_RetryAPICall(in <-chan string) {
    for line := range in {
        for i := 1; i <= 3; i++ {
            fmt.Printf("Retry %d for line: %s\n", i, line)
            time.Sleep(200 * time.Millisecond)
        }
    }
}
