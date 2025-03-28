package reader

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const(
	directoryreaderror ="Error reading directory:"
	txtsuffix =".txt"
	fileopenerror="Error opening file:"
	filereaderror="Error reading file:"
)
func ReadFiles(dirPath string, linesChan chan<- string, wg *sync.WaitGroup) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(directoryreaderror, err)
		return
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), txtsuffix) {
			wg.Add(1)
			go readFile(filepath.Join(dirPath, file.Name()), linesChan, wg)
		}
	}
}

func readFile(filePath string, linesChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(fileopenerror, filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesChan <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(filereaderror, filePath, err)
	}
}
