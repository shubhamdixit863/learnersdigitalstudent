package fileReader

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func ReadFiles(dirPath string, lineChan chan string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var wg sync.WaitGroup

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			wg.Add(1)
			go func(fileName string) {
				defer wg.Done()
				readFile(filepath.Join(dirPath, fileName), lineChan)
			}(file.Name())
		}
	}

	wg.Wait() // ✅ Ensure all files are read before closing the channel
}

// Read a single file and send lines to channel
func readFile(filePath string, lineChan chan string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineChan <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", filePath, err)
	}
}
