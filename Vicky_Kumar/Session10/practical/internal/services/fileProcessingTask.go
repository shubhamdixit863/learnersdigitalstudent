package services

import (
	"fmt"
	"os"
	"practical/internal/errorss"
	"strings"
)

type FileProcessing struct {
	Filepath string
	Data     string
}

func NewFileProcessing(filepath string) Task {
	return &FileProcessing{Filepath: filepath, Data: ""}
}

func (f *FileProcessing) Run() error {

	file, err := os.Open(f.Filepath)
	if err != nil {
		return &errorss.TaskError{Msg: fmt.Sprintf("failed to open file: %v", err)}
	}
	defer file.Close()
	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		return &errorss.TaskError{Msg: fmt.Sprintf("failed to read file: %v", err)}
	}

	updatedContent := strings.ToUpper(string(data))
	f.Data = updatedContent

	err = os.WriteFile(f.Filepath, []byte(updatedContent), 0644)
	if err != nil {
		return &errorss.TaskError{Msg: fmt.Sprintf("failed to write to file: %v", err)}
	}
	fmt.Println("File Processed Successfully!")
	return nil
}
