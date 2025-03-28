// File: services/indexer/indexer.go
package indexer

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"web_crawler_app/internal/models"
	"web_crawler_app/internal/services/extract"
)

// IndexService handles document indexing
type IndexService struct {
	index        *models.InvertedIndex
	indexFile    string
	mutex        sync.RWMutex
	documentInfo map[string]DocumentInfo
}

// DocumentInfo stores additional document information
type DocumentInfo struct {
	Title    string
	DocLen   int               // Document length in terms
	Snippets map[string]string // Term -> surrounding text
}

// NewIndexService creates a new index service
func NewIndexService(indexFile string) *IndexService {
	return &IndexService{
		index:        models.NewInvertedIndex(),
		indexFile:    indexFile,
		documentInfo: make(map[string]DocumentInfo),
	}
}

// IndexDocument indexes a document
func (is *IndexService) IndexDocument(data extract.ExtractedData) error {
	is.mutex.Lock()
	defer is.mutex.Unlock()

	// Create term frequency map for this document
	termFreq := make(map[string]int)

	// Count term frequencies
	for _, term := range data.Keywords {
		termFreq[term]++
	}

	// Update document info
	is.documentInfo[data.URL] = DocumentInfo{
		Title:    data.Title,
		DocLen:   len(data.Keywords),
		Snippets: extractSnippets(data.TextContent, termFreq),
	}

	// Update the inverted index
	for term, freq := range termFreq {
		// Create posting list if it doesn't exist
		if _, exists := is.index.Index[term]; !exists {
			is.index.Index[term] = make(map[string]int)
			is.index.DocFreq[term] = 0
		}

		// Add document to posting list
		is.index.Index[term][data.URL] = freq

		// Update document frequency
		is.index.DocFreq[term]++
	}

	// Update document count
	is.index.DocCount++

	return nil
}

// GetIndex returns the inverted index
func (is *IndexService) GetIndex() *models.InvertedIndex {
	is.mutex.RLock()
	defer is.mutex.RUnlock()

	return is.index
}

// GetDocumentInfo returns information about a document
func (is *IndexService) GetDocumentInfo(url string) (DocumentInfo, bool) {
	is.mutex.RLock()
	defer is.mutex.RUnlock()

	info, exists := is.documentInfo[url]
	return info, exists
}

// SaveIndex saves the index to a file
func (is *IndexService) SaveIndex() error {
	is.mutex.RLock()
	defer is.mutex.RUnlock()

	// Create index data for serialization
	indexData := struct {
		Index        map[string]map[string]int `json:"index"`
		DocFreq      map[string]int            `json:"doc_freq"`
		DocCount     int                       `json:"doc_count"`
		DocumentInfo map[string]DocumentInfo   `json:"document_info"`
	}{
		Index:        is.index.Index,
		DocFreq:      is.index.DocFreq,
		DocCount:     is.index.DocCount,
		DocumentInfo: is.documentInfo,
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(indexData, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(is.indexFile, data, 0644)
}

// LoadIndex loads the index from a file
func (is *IndexService) LoadIndex() error {
	is.mutex.Lock()
	defer is.mutex.Unlock()

	// Check if file exists
	if _, err := os.Stat(is.indexFile); os.IsNotExist(err) {
		return fmt.Errorf("index file does not exist: %s", is.indexFile)
	}

	// Read file
	data, err := os.ReadFile(is.indexFile)
	if err != nil {
		return err
	}

	// Unmarshal JSON
	var indexData struct {
		Index        map[string]map[string]int `json:"index"`
		DocFreq      map[string]int            `json:"doc_freq"`
		DocCount     int                       `json:"doc_count"`
		DocumentInfo map[string]DocumentInfo   `json:"document_info"`
	}

	err = json.Unmarshal(data, &indexData)
	if err != nil {
		return err
	}

	// Update index
	is.index.Index = indexData.Index
	is.index.DocFreq = indexData.DocFreq
	is.index.DocCount = indexData.DocCount
	is.documentInfo = indexData.DocumentInfo

	return nil
}

// extractSnippets extracts relevant snippets for terms
func extractSnippets(text string, terms map[string]int) map[string]string {
	snippets := make(map[string]string)

	// For each term, find a relevant snippet
	for term := range terms {
		idx := IndexOf(text, term)
		if idx == -1 {
			continue
		}

		// Extract snippet around the term
		start := max(0, idx-50)
		end := min(len(text), idx+len(term)+50)

		snippet := text[start:end]
		snippets[term] = snippet
	}

	return snippets
}

// IndexOf finds the first occurrence of a substring in a string (case insensitive)
func IndexOf(text, term string) int {
	text = strings.ToLower(text)
	term = strings.ToLower(term)
	return strings.Index(text, term)
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
