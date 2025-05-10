package crawl

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type CrawlService struct {
	maxDepth         int
	maxConcurrent    int
	visited          map[string]bool
	visitedMutex     sync.RWMutex
	concurrencyLimit chan struct{}
}

type CrawlResult struct {
	URL     string
	Content string
	Links   []string
	Depth   int
	Error   error
}

func NewCrawlService(maxDepth, maxConcurrent int) *CrawlService {
	return &CrawlService{
		maxDepth:         maxDepth,
		maxConcurrent:    maxConcurrent,
		visited:          make(map[string]bool),
		concurrencyLimit: make(chan struct{}, maxConcurrent),
	}
}

func (cs *CrawlService) CrawlUrls(startUrls []string, resultChan chan<- CrawlResult) error {
	if len(startUrls) == 0 {
		return errors.New("no URLs provided for crawling")
	}

	var wg sync.WaitGroup
	for _, urlStr := range startUrls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			cs.crawlURL(url, 0, resultChan)
		}(urlStr)
	}

	wg.Wait()
	return nil
}

func (cs *CrawlService) crawlURL(urlStr string, depth int, resultChan chan<- CrawlResult) {
	// Check if we've reached the maximum depth
	if depth > cs.maxDepth {
		return
	}

	// Normalize the URL
	normalizedURL := normalizeURL(urlStr)

	// Check if we've already visited this URL
	cs.visitedMutex.RLock()
	if cs.visited[normalizedURL] {
		cs.visitedMutex.RUnlock()
		return
	}
	cs.visitedMutex.RUnlock()

	// Mark as visited
	cs.visitedMutex.Lock()
	cs.visited[normalizedURL] = true
	cs.visitedMutex.Unlock()

	// Respect the concurrency limit
	cs.concurrencyLimit <- struct{}{}
	defer func() { <-cs.concurrencyLimit }()

	// Fetch the URL
	content, links, err := cs.fetchURL(normalizedURL)

	// Send the result
	result := CrawlResult{
		URL:     normalizedURL,
		Content: content,
		Links:   links,
		Depth:   depth,
		Error:   err,
	}

	resultChan <- result

	// If there was an error, don't crawl further
	if err != nil {
		return
	}

	// Recursively crawl the links
	var wg sync.WaitGroup
	for _, link := range links {
		absLink := makeAbsoluteURL(normalizedURL, link)
		if absLink == "" {
			continue
		}

		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			cs.crawlURL(link, depth+1, resultChan)
		}(absLink)
	}

	wg.Wait()
}

// fetchURL fetches the content of a URL and extracts links
func (cs *CrawlService) fetchURL(urlStr string) (string, []string, error) {
	// Create a client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Make the request
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return "", nil, err
	}

	// Set a user agent
	req.Header.Set("User-Agent", "GoSearchEngine/1.0")

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	// Check if the response is OK
	if resp.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Check if the content type is HTML
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", nil, fmt.Errorf("not HTML content: %s", contentType)
	}

	// Read the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	content := string(body)

	// Extract links
	links := extractLinks(content)

	return content, links, nil
}

// extractLinks extracts links from HTML content
func extractLinks(htmlContent string) []string {
	links := make([]string, 0)

	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return links
	}

	var extractFunc func(*html.Node)
	extractFunc = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
					break
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractFunc(c)
		}
	}

	extractFunc(doc)
	return links
}

// normalizeURL normalizes a URL by removing fragments and trailing slashes
func normalizeURL(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}

	// Remove fragment
	u.Fragment = ""

	// Remove trailing slash
	path := u.Path
	if len(path) > 1 && strings.HasSuffix(path, "/") {
		u.Path = path[:len(path)-1]
	}

	return u.String()
}

// makeAbsoluteURL converts a relative URL to an absolute URL
func makeAbsoluteURL(baseURL, relURL string) string {
	// If it's already absolute, return it
	if strings.HasPrefix(relURL, "http://") || strings.HasPrefix(relURL, "https://") {
		return relURL
	}

	// Parse the base URL
	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	// Parse the relative URL
	rel, err := url.Parse(relURL)
	if err != nil {
		return ""
	}

	// Resolve the relative URL against the base URL
	abs := base.ResolveReference(rel)

	// Only return HTTP and HTTPS URLs
	if abs.Scheme != "http" && abs.Scheme != "https" {
		return ""
	}

	return abs.String()
}
