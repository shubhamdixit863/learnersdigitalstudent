package services

import (
	"fmt"
	"os"
	"task/internal/utils"
)

type FileProcessing struct {
	FileName string
}

func (f *FileProcessing) Run() error {
	fmt.Println("Processing file", f.FileName)

    data, err := os.OpenFile(f.FileName, os.O_RDONLY, 0644)
    if err != nil {
        return &utils.TaskError{Message: "Failed to open file"}
    }

    defer data.Close()
    return nil
}

