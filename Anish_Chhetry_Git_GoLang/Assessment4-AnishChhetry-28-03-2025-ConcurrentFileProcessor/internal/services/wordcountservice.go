package services

import (
	"log"
	"strings"
)

func WordCount(storage map[string]string) {
	for _, words := range storage {
		wordSlice := strings.Split(words, spaceString)
		log.Println(len(wordSlice))
	}
}
