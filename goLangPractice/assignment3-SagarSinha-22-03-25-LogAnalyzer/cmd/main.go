package main

import (
	"fmt"
	"goRoutinesHandsOn/internal/services"
	"sync"
)

const chunkSize = 5
const fileName = "sample.log"
const resultStr = "Log Level Counts: "

func main() {
	var wg sync.WaitGroup
	results := make(chan map[string]int, 10)
	finalCounts := make(map[string]int)

	services.Reader(fileName, chunkSize, &wg, results)

	for r := range results {
		services.Aggregator(&finalCounts, r)
	}

	fmt.Println(resultStr)

	for log, count := range finalCounts {
		fmt.Println(log, ": ", count)
	}
}
