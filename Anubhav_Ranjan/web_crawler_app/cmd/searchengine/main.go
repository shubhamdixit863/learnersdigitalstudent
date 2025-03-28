package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"web_crawler_app/internal/config"
	"web_crawler_app/internal/services/crawl"
	"web_crawler_app/internal/services/extract"
	"web_crawler_app/internal/services/indexer"
	"web_crawler_app/internal/services/search"
)

func main() {
	// Initialize configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize services
	crawlService := crawl.NewCrawlService(cfg.CrawlDepth, cfg.MaxConcurrentRequests)
	extractService := extract.NewExtractService()
	indexService := indexer.NewIndexService(cfg.IndexFilePath)
	searchService := search.NewSearchService(indexService)

	// Try to load existing index
	err = indexService.LoadIndex()
	if err != nil {
		log.Printf("Could not load existing index, starting fresh: %v", err)
	} else {
		log.Println("Loaded existing index")
	}

	// Check if we need to crawl
	shouldCrawl := true
	if len(os.Args) > 1 && os.Args[1] == "search" {
		shouldCrawl = false
	}

	if shouldCrawl {
		// Create channels for the pipeline
		crawledDataChan := make(chan crawl.CrawlResult, cfg.ChannelBufferSize)
		extractedDataChan := make(chan extract.ExtractedData, cfg.ChannelBufferSize)

		// WaitGroup to wait for all goroutines to finish
		var wg sync.WaitGroup

		// Start crawling
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(crawledDataChan)

			log.Println("Starting crawling...")
			err := crawlService.CrawlUrls(cfg.StartUrls, crawledDataChan)
			if err != nil {
				log.Printf("Error during crawling: %v", err)
			}
			log.Println("Crawling complete")
		}()

		// Start extraction
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(extractedDataChan)

			log.Println("Starting extraction...")
			for result := range crawledDataChan {
				if result.Error != nil {
					log.Printf("Error crawling %s: %v", result.URL, result.Error)
					continue
				}

				extractedData, err := extractService.Extract(result)
				if err != nil {
					log.Printf("Error extracting data from %s: %v", result.URL, err)
					continue
				}
				extractedDataChan <- extractedData
			}
			log.Println("Extraction complete")
		}()

		// Build index
		wg.Add(1)
		go func() {
			defer wg.Done()

			log.Println("Starting indexing...")
			for data := range extractedDataChan {
				err := indexService.IndexDocument(data)
				if err != nil {
					log.Printf("Error indexing document %s: %v", data.URL, err)
				}
			}
			log.Println("Indexing complete")
		}()

		// Wait for all goroutines to finish
		wg.Wait()

		// Save index to disk
		log.Println("Saving index...")
		if err := indexService.SaveIndex(); err != nil {
			log.Fatalf("Failed to save index: %v", err)
		}
		log.Println("Index saved successfully")
	}

	// Start interactive search
	fmt.Println("\nSearch Engine - Interactive Mode")
	fmt.Println("================================")
	runSearchInterface(searchService)
}

func runSearchInterface(searchService *search.SearchService) {
	for {
		fmt.Print("\nEnter search query (or 'exit' to quit): ")
		var query string
		fmt.Scanln(&query)

		if query == "exit" {
			break
		}

		results, err := searchService.Search(query)
		if err != nil {
			fmt.Printf("Error during search: %v\n", err)
			continue
		}

		if len(results) == 0 {
			fmt.Println("No results found")
			continue
		}

		fmt.Printf("Found %d results:\n", len(results))
		for i, result := range results {
			fmt.Printf("\n%d. %s\n", i+1, result.Title)
			fmt.Printf("   URL: %s\n", result.URL)
			fmt.Printf("   Score: %.4f\n", result.Score)
			if result.Snippet != "" {
				fmt.Printf("   Snippet: %s\n", result.Snippet)
			}

			if i >= 9 {
				fmt.Println("\n...more results available")
				break
			}
		}
	}

	fmt.Println("Goodbye!")
}
