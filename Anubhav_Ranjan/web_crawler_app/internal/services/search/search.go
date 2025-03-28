// File: services/search/search.go
package search

import (
	"errors"
	"math"
	"sort"
	"strings"
	"web_crawler_app/internal/models"
	"web_crawler_app/internal/services/indexer"
)

// SearchService handles search operations
type SearchService struct {
	indexService *indexer.IndexService
}

// NewSearchService creates a new search service
func NewSearchService(indexService *indexer.IndexService) *SearchService {
	return &SearchService{
		indexService: indexService,
	}
}

// Search performs a search query and returns results
func (ss *SearchService) Search(query string) ([]models.SearchResult, error) {
	// Check if query is empty
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, errors.New("empty search query")
	}

	// Tokenize the query
	queryTerms := tokenizeQuery(query)
	if len(queryTerms) == 0 {
		return nil, errors.New("no valid search terms in query")
	}

	// Get the index
	index := ss.indexService.GetIndex()

	// Calculate scores using TF-IDF
	scores := ss.calculateScores(queryTerms, index)

	// Convert scores to search results
	results := ss.convertScoresToResults(scores, queryTerms)

	// Sort results by score (descending)
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	return results, nil
}

// tokenizeQuery tokenizes a search query
func tokenizeQuery(query string) []string {
	// Split query into terms
	terms := strings.Fields(strings.ToLower(query))

	// Clean terms
	cleanTerms := make([]string, 0)
	for _, term := range terms {
		// Remove non-alphanumeric characters
		term = strings.Map(func(r rune) rune {
			if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
				return r
			}
			return -1
		}, term)

		// Skip empty and short terms
		if len(term) <= 2 {
			continue
		}

		cleanTerms = append(cleanTerms, term)
	}

	return cleanTerms
}

// calculateScores calculates TF-IDF scores for documents
func (ss *SearchService) calculateScores(queryTerms []string, index *models.InvertedIndex) map[string]float64 {
	scores := make(map[string]float64)

	// For each query term
	for _, term := range queryTerms {
		// Skip if term is not in index
		if _, exists := index.Index[term]; !exists {
			continue
		}

		// Calculate IDF for this term
		idf := math.Log10(float64(index.DocCount) / float64(index.DocFreq[term]))

		// For each document containing this term
		for docURL, tf := range index.Index[term] {
			// Get document info
			docInfo, exists := ss.indexService.GetDocumentInfo(docURL)
			if !exists {
				continue
			}

			// Calculate normalized TF
			normalizedTF := float64(tf) / float64(docInfo.DocLen)

			// Calculate TF-IDF
			tfidf := normalizedTF * idf

			// Add to document score
			scores[docURL] += tfidf
		}
	}

	return scores
}

// convertScoresToResults converts score map to search results
func (ss *SearchService) convertScoresToResults(scores map[string]float64, queryTerms []string) []models.SearchResult {
	results := make([]models.SearchResult, 0, len(scores))

	for docURL, score := range scores {
		// Get document info
		docInfo, exists := ss.indexService.GetDocumentInfo(docURL)
		if !exists {
			continue
		}

		// Find best snippet
		snippet := findBestSnippet(docInfo.Snippets, queryTerms)

		// Create search result
		result := models.SearchResult{
			URL:     docURL,
			Title:   docInfo.Title,
			Snippet: snippet,
			Score:   score,
		}

		results = append(results, result)
	}

	return results
}

// findBestSnippet finds the best snippet for a search result
func findBestSnippet(snippets map[string]string, queryTerms []string) string {
	// Try to find a snippet containing the first query term
	for _, term := range queryTerms {
		if snippet, exists := snippets[term]; exists {
			return "..." + snippet + "..."
		}
	}

	// If no snippet found, return empty string
	return ""
}
