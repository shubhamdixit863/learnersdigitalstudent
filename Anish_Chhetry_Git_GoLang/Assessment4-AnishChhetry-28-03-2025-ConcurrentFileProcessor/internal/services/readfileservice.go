package services

import (
	"log"
	"os"
)

func ReadFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	return string(file)
}
