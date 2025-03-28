package main

import (
	"concurrent-file-processor/config"
	"concurrent-file-processor/constants"
	"concurrent-file-processor/processor"
	"concurrent-file-processor/utils"
	"fmt"
	"os"
	"sync"
)

func main() {
	// Parse command-line arguments
	args := config.ParseArgs()

	// Discover .txt files
	files, err := utils.FindTextFiles(args.DirPath)
	if err != nil {
		fmt.Println(constants.ErrReadingDir, err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println(constants.NoTxtFiles)
		os.Exit(1)
	}

	// Fan-Out: Read files concurrently
	linesCh := make(chan string, 100)
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go utils.ReadFileLines(file, linesCh, &wg)
	}
	go func() {
		wg.Wait()
		close(linesCh)
	}()

	// Fan-In: Process based on mode
	processor.ProcessLines(linesCh, args.Mode, args.Keyword)
}
