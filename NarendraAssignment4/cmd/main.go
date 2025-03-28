package main

import (
    "flag"
    "fmt"
    "os"

    "assignment4/internal/services"
)

func main() {
    dir := "./txtfiles"

    mode := flag.String("mode", "filter", "Mode: filter | wordcount | apicall | retryapi")
    keyword := flag.String("keyword", "error", "Keyword for filter mode (optional)")
    flag.Parse()

    linesChan := make(chan string)

    go services.GenerateLines(dir, linesChan)

    switch *mode {
    case "filter":
        services.ProcessLines_Filter(linesChan, *keyword)
    case "wordcount":
        services.ProcessLines_WordCount(linesChan)
    case "apicall":
        services.ProcessLines_APICall(linesChan)
    case "retryapi":
        services.ProcessLines_RetryAPICall(linesChan)
    default:
        fmt.Println(" Unknown mode")
        os.Exit(1)
    }
}
