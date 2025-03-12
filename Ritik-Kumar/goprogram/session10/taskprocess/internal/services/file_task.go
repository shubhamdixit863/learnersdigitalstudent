package services

import (
    "errors"
    "fmt"
    "os"
)

type FileProcessingTask struct {
    FileName string
}

func (f *FileProcessingTask) Run() error {
    file, err := os.Open(f.FileName)
    if err != nil {
        return &TaskError{"FileProcessingTask", errors.New("failed to open file")}
    }
    defer file.Close()
    
    fmt.Println("Processing file:", f.FileName)
    return nil
}
