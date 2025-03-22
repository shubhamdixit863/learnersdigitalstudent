package services

import (
	"fmt"
	"regexp"
	"session15/storage"
	"strings"
)

func ExtractHrefs(input string, baseurl string) {
	re := regexp.MustCompile(`href=["']([^"']+)["']`)
	matches := re.FindAllStringSubmatch(input, -1)

	var hrefs []string
	for _, match := range matches {
		link := match[1]
		// Handle relative URLs
		if !strings.HasPrefix(link, "http") {
			link = "https://usf-cs272-s25.github.io" + link
		}
		hrefs = append(hrefs, link)
		storage.Mu.Lock()
		storage.UrlSlice = append(storage.UrlSlice, link)
		storage.Mu.Unlock()
		fmt.Println()

	}
	//storage.UrlSlice = append(storage.UrlSlice, hrefs...)
	//return hrefs
}
