// File: models/models.go
package models

// Document represents a crawled and processed document
type Document struct {
	URL      string
	Title    string
	Content  string
	Keywords []string
}

// SearchResult represents a single search result
type SearchResult struct {
	URL     string
	Title   string
	Snippet string
	Score   float64
}

// InvertedIndex represents the inverted index data structure
type InvertedIndex struct {
	// Map from term to document frequencies
	// The inner map is from document URL to term frequency in that document
	Index map[string]map[string]int
	// Document frequency - number of documents containing each term
	DocFreq map[string]int
	// Total number of documents in the index
	DocCount int
}

// NewInvertedIndex creates a new inverted index
func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		Index:    make(map[string]map[string]int),
		DocFreq:  make(map[string]int),
		DocCount: 0,
	}
}
