package main

import (
	"WebCrawler/internal/services"
	"fmt"
	"log"
	"sync"
)

func main() {
	crawler := services.NewCrawlerService()
	mainLink := "https://usf-cs272-s25.github.io/top10/"

	links, errLink := crawler.Crawl(mainLink)
	if errLink != nil {
		fmt.Println(errLink)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	chCleanedText := make(chan string)
	chLink := make(chan string)
	for _, link := range links {
		wg.Add(1)

		go func(link string) {
			defer wg.Done()

			htmlContent, errHTML := crawler.GiveHTML(link)
			if errHTML != nil {
				fmt.Println(errHTML)
				return
			}
			extractedText := crawler.ExtractText(htmlContent)
			cleanedText := crawler.CleanText(extractedText)
			chCleanedText <- cleanedText
			chLink <- link

		}(link)
	}
	go crawler.IndexService(chCleanedText, chLink, &mu)
	wg.Wait()
	searchFor := ""
	YorN := ""
	for i := 0; i == 0; {
		fmt.Println("Do you want to search?(y or n) ")
		fmt.Scanf("%s\n", &YorN)
		if YorN == "y" {
			fmt.Println("What would you like to Search For: ")
			fmt.Scanf("%s\n", &searchFor)
			results, errResult := crawler.Search(searchFor, &mu)
			if errResult != nil {
				log.Println(errResult)
			}
			for i, result := range results {
				log.Println("Link", i+1, result)
			}
		} else if YorN == "n" {
			i++
		} else {
			fmt.Println("Please type y or n")
		}
	}
}
