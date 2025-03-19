package services

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Crawler interface {
	Crawl(pageURL string) ([]string, error)
	GiveHTML(pageURL string) (string, error)
	IndexService(text string, document string)
	ExtractText(htmlContent string) string
	CleanText(text string) string
	Search(searchFor string)
}

type CrawlerService struct {
	InvertedIndex map[string]map[string]int
}

func NewCrawlerService() *CrawlerService {
	return &CrawlerService{InvertedIndex: make(map[string]map[string]int)}
}
func (cr *CrawlerService) Crawl(pageURL string) ([]string, error) {

	htmlContent, err := cr.GiveHTML(pageURL)
	if err != nil {
		return nil, err
	}
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return nil, err
	}

	links := make([]string, 0)
	start := 0

	for {
		aTagIndex := strings.Index(htmlContent[start:], "<a")
		if aTagIndex == -1 {
			break
		}
		hrefIndex := strings.Index(htmlContent[start:], "href=\"")
		if hrefIndex == -1 {
			break
		}
		start += hrefIndex + 6

		endQuote := strings.Index(htmlContent[start:], "\"")
		if endQuote == -1 {
			break
		}

		link := htmlContent[start : start+endQuote]

		start += endQuote

		parsedLink, err := url.Parse(link)
		if err == nil {
			absLink := baseURL.ResolveReference(parsedLink).String()
			if !contains(links, absLink) {
				links = append(links, absLink)
			}
		}
	}

	return links, nil
}

func (cr *CrawlerService) GiveHTML(pageURL string) (string, error) {
	data, err := http.Get(pageURL)
	if err != nil {
		return "", err
	}
	defer data.Body.Close()

	htmlBytes, err := io.ReadAll(data.Body)
	if err != nil {
		return "", err
	}
	return string(htmlBytes), nil
}
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
