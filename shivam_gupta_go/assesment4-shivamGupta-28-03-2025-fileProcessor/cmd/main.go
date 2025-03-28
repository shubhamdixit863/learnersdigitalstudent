package main

import (
	"fileProcessor/internal/services"
	"fileProcessor/internal/utils"
	"fmt"
)


const (errScan ="Error scanning directory:"
      dirPath = "../internal/fileDir"
)

func main() {

	files, err := utils.GetTxtFiles(dirPath)
	if err != nil {
		fmt.Println(errScan, err)
		return
	}

	services.ProcessFiles(files)
}