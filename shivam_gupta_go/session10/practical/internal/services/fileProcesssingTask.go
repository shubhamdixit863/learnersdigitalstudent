package services

import (
	"errors"
	"fmt"
	"os"
)

type FileProcessing struct {
	Filepath string
}

func (a *FileProcessing) Run() error {
	// fmt.Println("enter file path")
	// var fp string
	// fmt.Scanln(&fp)

	fmt.Println("opening file.....")
	Data, err := os.ReadFile(a.Filepath)

	if err != nil {
		return errors.New("opening failed")
	}
	fmt.Println(string(Data))
	return err
}
