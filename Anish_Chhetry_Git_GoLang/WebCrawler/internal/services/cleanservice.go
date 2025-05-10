package services

import "strings"

func (cr *CrawlerService) CleanText(text string) string {
	cleanedText := ""
	spaceFound := false

	for _, char := range text {
		if char == ' ' || char == '\n' || char == '\t' {
			if !spaceFound {
				cleanedText += " "
				spaceFound = true
			}
		} else {
			cleanedText += string(char)
			spaceFound = false
		}
	}

	return strings.TrimSpace(cleanedText)
}
