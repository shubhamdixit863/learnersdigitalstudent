// File: services/extract/extractor.go
package extract

import (
	"regexp"
	"strings"
	"web_crawler_app/internal/services/crawl"

	"golang.org/x/net/html"
)

// ExtractService extracts text and metadata from HTML content
type ExtractService struct {
	stopWords map[string]bool
}

// ExtractedData represents extracted data from a document
type ExtractedData struct {
	URL         string
	Title       string
	TextContent string
	Keywords    []string
}

// NewExtractService creates a new extract service
func NewExtractService() *ExtractService {
	return &ExtractService{
		stopWords: loadStopWords(),
	}
}

// Extract extracts text and metadata from HTML content
func (es *ExtractService) Extract(result crawl.CrawlResult) (ExtractedData, error) {
	if result.Error != nil {
		return ExtractedData{}, result.Error
	}

	// Parse HTML
	doc, err := html.Parse(strings.NewReader(result.Content))
	if err != nil {
		return ExtractedData{}, err
	}

	// Extract title
	title := extractTitle(doc)

	// Extract text content
	textContent := extractText(doc)

	// Extract keywords
	keywords := es.extractKeywords(textContent)

	return ExtractedData{
		URL:         result.URL,
		Title:       title,
		TextContent: textContent,
		Keywords:    keywords,
	}, nil
}

// extractTitle extracts the title from HTML
func extractTitle(n *html.Node) string {
	var title string

	var extractTitleFunc func(*html.Node)
	extractTitleFunc = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			if n.FirstChild != nil {
				title = n.FirstChild.Data
				return
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractTitleFunc(c)
		}
	}

	extractTitleFunc(n)
	return title
}

// extractText extracts text content from HTML
func extractText(n *html.Node) string {
	var text strings.Builder

	var extractTextFunc func(*html.Node)
	extractTextFunc = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "script", "style", "noscript", "iframe", "nav", "footer":
				return
			}
		}

		if n.Type == html.TextNode {
			text.WriteString(n.Data + " ")
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractTextFunc(c)
		}
	}

	extractTextFunc(n)
	return cleanText(text.String())
}

// cleanText cleans extracted text
func cleanText(text string) string {
	// Replace multiple spaces with a single space
	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")

	// Trim space
	return strings.TrimSpace(text)
}

// extractKeywords extracts keywords from text
func (es *ExtractService) extractKeywords(text string) []string {
	// Split text into words
	words := strings.Fields(strings.ToLower(text))

	// Filter out stop words and short words
	keywords := make([]string, 0)
	for _, word := range words {
		// Clean the word
		word = cleanWord(word)

		// Skip empty, short words, and stop words
		if len(word) <= 2 || es.stopWords[word] {
			continue
		}

		keywords = append(keywords, word)
	}

	return keywords
}

// cleanWord cleans a word by removing non-alphanumeric characters
func cleanWord(word string) string {
	// Remove non-alphanumeric characters
	re := regexp.MustCompile(`[^\p{L}\p{N}]`)
	word = re.ReplaceAllString(word, "")

	return word
}

// loadStopWords loads a list of common stop words
func loadStopWords() map[string]bool {
	stopWordsList := []string{
		"a", "about", "above", "after", "again", "against", "all", "am", "an", "and",
		"any", "are", "aren't", "as", "at", "be", "because", "been", "before", "being",
		"below", "between", "both", "but", "by", "can't", "cannot", "could", "couldn't",
		"did", "didn't", "do", "does", "doesn't", "doing", "don't", "down", "during",
		"each", "few", "for", "from", "further", "had", "hadn't", "has", "hasn't", "have",
		"haven't", "having", "he", "he'd", "he'll", "he's", "her", "here", "here's",
		"hers", "herself", "him", "himself", "his", "how", "how's", "i", "i'd", "i'll",
		"i'm", "i've", "if", "in", "into", "is", "isn't", "it", "it's", "its", "itself",
		"let's", "me", "more", "most", "mustn't", "my", "myself", "no", "nor", "not",
		"of", "off", "on", "once", "only", "or", "other", "ought", "our", "ours", "ourselves",
		"out", "over", "own", "same", "shan't", "she", "she'd", "she'll", "she's", "should",
		"shouldn't", "so", "some", "such", "than", "that", "that's", "the", "their", "theirs",
		"them", "themselves", "then", "there", "there's", "these", "they", "they'd", "they'll",
		"they're", "they've", "this", "those", "through", "to", "too", "under", "until", "up",
		"very", "was", "wasn't", "we", "we'd", "we'll", "we're", "we've", "were", "weren't",
		"what", "what's", "when", "when's", "where", "where's", "which", "while", "who", "who's",
		"whom", "why", "why's", "with", "won't", "would", "wouldn't", "you", "you'd", "you'll",
		"you're", "you've", "your", "yours", "yourself", "yourselves",
	}

	stopWords := make(map[string]bool)
	for _, word := range stopWordsList {
		stopWords[word] = true
	}

	return stopWords
}
