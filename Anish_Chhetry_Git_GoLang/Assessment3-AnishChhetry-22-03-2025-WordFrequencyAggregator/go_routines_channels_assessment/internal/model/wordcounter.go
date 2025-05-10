package model

type WordCounter interface {
	CountWords(text string) map[string]int
}
