package services

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "sync"
)

func GenerateLines(dir string, out chan<- string) {
    var wg sync.WaitGroup

    files, _ := filepath.Glob(filepath.Join(dir, "*.txt"))

    for _, file := range files {
        wg.Add(1)
        go func(f string) {
            defer wg.Done()
            readFileLines(f, out)
        }(file)
    }

    wg.Wait()
    close(out)
}

func readFileLines(filePath string, out chan<- string) {
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        out <- scanner.Text()
    }
}
