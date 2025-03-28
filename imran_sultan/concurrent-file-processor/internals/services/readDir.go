package services

import (
	"concurrent_file_processing/internals/utils"
	"fmt"
	"log"
	"os"
)

func ReadDirFunc() {
	dir, err := os.ReadDir(utils.DirName)
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
	data, err := os.ReadFile(utils.DirName + "/" + file.Name())
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(data)

}
