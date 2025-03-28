package main

import (
	"fileprocessor/internal/processor"
	"fileprocessor/internal/utils"
	"fmt"
)

func main() {

	directory := utils.SearchedDirectory
	
	for _, mode := range utils.Modes {

		err := processor.ProcessFiles(directory, mode)
		if err != nil {
			fmt.Println(utils.FileProcessingError, err)
			fmt.Println(utils.ModeProcessingError, mode)
			return
		}
	}
}