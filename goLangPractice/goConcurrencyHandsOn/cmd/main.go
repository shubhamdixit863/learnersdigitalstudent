package main

import (
	"fmt"
	"goConcurrencyHandsOn/internal/services"
	"goConcurrencyHandsOn/internal/utilities"
	"sync"
)

func main() {
	lines := make(chan string)
	var wg sync.WaitGroup

	services.ScanDirectory(utilities.FilePath, lines, &wg)

	mode, keyword := selectMode()

	services.ProcessLines(lines, mode, keyword)
}

func selectMode() (int, string) {
	var mode int
	var keyword string

	fmt.Println(utilities.TakeInput)
	fmt.Println(utilities.Modes)
	fmt.Scan(&mode)

	if mode == 1 {
		fmt.Println(utilities.Keyword)
		fmt.Scan(&keyword)
		return mode, keyword
	}

	return mode, ""
}
