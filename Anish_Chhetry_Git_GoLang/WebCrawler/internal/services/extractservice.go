package services

import (
	"strings"
)

func (cr *CrawlerService) ExtractText(htmlContent string) string {
	bodyStart := strings.Index(htmlContent, "<body")
	if bodyStart == -1 {
		return ""
	}

	bodyStart = strings.Index(htmlContent[bodyStart:], ">")
	if bodyStart == -1 {
		return ""
	}
	bodyStart += strings.Index(htmlContent, "<body") + 1

	bodyEnd := strings.Index(htmlContent[bodyStart:], "</body>")
	if bodyEnd == -1 {
		return ""
	}

	bodyEnd += bodyStart

	bodyContent := htmlContent[bodyStart:bodyEnd]
	textContent := ""
	inTag := false

	for _, char := range bodyContent {

		if char == '<' {
			inTag = true
		} else if char == '>' {
			inTag = false
		} else if !inTag {
			textContent += string(char)
		}
	}

	return strings.TrimSpace(textContent)
}
