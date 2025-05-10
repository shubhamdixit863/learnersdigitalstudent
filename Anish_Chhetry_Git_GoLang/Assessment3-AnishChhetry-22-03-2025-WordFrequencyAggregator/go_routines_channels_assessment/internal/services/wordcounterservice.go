package services

import (
	"strings"
)

const (
	extras = ".,!?;:'\"()[]{}<>"
)

type SimpleWordCounter struct{}

func NewSimpleWordCounter() *SimpleWordCounter {
	return &SimpleWordCounter{}
}

func (swc SimpleWordCounter) CountWords(text string) map[string]int {
	wordFreq := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		lowerWord := strings.ToLower(strings.Trim(word, extras))

		if lowerWord != "" {
			wordFreq[lowerWord]++
		}
	}

	return wordFreq
}
