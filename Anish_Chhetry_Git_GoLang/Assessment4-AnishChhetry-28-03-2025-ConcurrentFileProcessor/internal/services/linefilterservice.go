package services

import (
	"log"
	"strings"
)

func LineFilter(storage map[string]string, SearchWord string) {
	for file, words := range storage {
		wordSlice := strings.Split(words, spaceString)
		for _, word := range wordSlice {
			if word == SearchWord {
				log.Println(words, inFileString, file)
			}
		}
	}
}
