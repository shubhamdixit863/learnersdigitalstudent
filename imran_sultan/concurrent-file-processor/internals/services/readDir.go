package services

import (
	"concurrent_file_processing/internals/utils"
	"fmt"
	"log"
	"os"
)

const (
	DirName = "text"
)

func ReadDirFunc() {
	dir, err := os.ReadDir(DirName)
	if err != nil {
		log.Fatal(err)
		return
	}
	for d, val := range dir {
		fmt.Println(d, val)
		ReadFileFunc(val)
		utils.FileInput <- val
	}
	close(utils.FileInput)

}

func ReadFileFunc(file os.DirEntry) string {
	data, err := os.ReadFile(DirName + "/" + file.Name())
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(data)

}
