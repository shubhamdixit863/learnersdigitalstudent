package services

import (
	"errors"
	"fmt"
	"os"
)

type FileProcessing struct {
	Filepath string
}

func NewFile() Task {
	return &FileProcessing{
		Filepath: "",
	}
}

func (f FileProcessing) Run() error {

	data, err := os.ReadFile("../../data.txt")
	if errors.Is(err, os.ErrNotExist) {

		return err

	}

	fmt.Println(string(data))
	return nil

}
