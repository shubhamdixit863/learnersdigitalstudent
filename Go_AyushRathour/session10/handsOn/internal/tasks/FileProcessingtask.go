package tasks

import (
	"fmt"
	"os"
)

type FileProcessingTask struct {
	Filename string
	Text     string
}

func NewFileProcessingTask(filename, text string) *FileProcessingTask {
	return &FileProcessingTask{
		Filename: filename,
		Text:     text,
	}
}

func (task *FileProcessingTask) Run() error {

	e := CreateFile(task.Filename, task.Text)
	if e != nil {
		return e
	}
	return ReadFile(task.Filename)
}

func CreateFile(filename, text string) error {
	fmt.Println("Writing file")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to Create file %s: %w", filename, err)
	}
	defer file.Close()

	length, err := file.WriteString(text)
	if err != nil {
		return fmt.Errorf("failed to Write file %s: %w", filename, err)
	}
	fmt.Printf("File name: %s\n", file.Name())
	fmt.Printf("File length: %d bytes\n", length)

	return nil
}

func ReadFile(filename string) error {
	fmt.Println("\nReading file")
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to Read file %s: %w", filename, err)
	}

	fmt.Println("File name: " + filename)
	fmt.Printf("File size: %d\n", len(data))
	fmt.Printf("File content: %s\n", data)

	return nil

}
