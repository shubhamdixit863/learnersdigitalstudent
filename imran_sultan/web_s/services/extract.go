package services

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"web_s/storage"
)

func GetHTMLFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
func ExtractLinksFromList(html string, baseURL string) []string {
	re := regexp.MustCompile(`<li>\s*<a\s+[^>]*href=["']([^"'>]+)["']`)
	matches := re.FindAllStringSubmatch(html, -1)

	var links []string
	for _, match := range matches {
		if len(match) > 1 {
			link := match[1]
			// Convert relative URLs to absolute
			if !strings.HasPrefix(link, "http") {
				link = baseURL + link
			}
			links = append(links, link)
		}
	}
	return links
}
func Extract(url string) {
	htmlContent, err := GetHTMLFromURL(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Println(htmlContent)

	storage.Links = ExtractLinksFromList(htmlContent, "https://usf-cs272-s25.github.io")

	fmt.Println("Extracted Links from Body:", len(storage.Links))
	for _, link := range storage.Links {
		fmt.Println(link)
	}
}
