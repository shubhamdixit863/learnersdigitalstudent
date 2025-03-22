// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"regexp"
// 	"strings"

// 	"golang.org/x/net/html"
// )

// //https://usf-cs272-s25.github.io/top10/

// type Data struct {
// 	UserId int    `json:"userId"`
// 	Title  string `json:"title"`
// }

// func getPlainTextFromURL(url string) (string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	htmlContent := string(body)

// 	// Remove HTML tags using regex
// 	re := regexp.MustCompile("<[^>]*>")
// 	plainText := re.ReplaceAllString(htmlContent, " ")

// 	// Trim and normalize spaces
// 	plainText = strings.Join(strings.Fields(plainText), " ")

// 	return plainText, nil
// }

// // func main() {

// // 	data, err := http.Get("https://usf-cs272-s25.github.io/top10/")
// // 	if err != nil {
// // 		return
// // 	}

// // 	body := data.Body

// // 	all, err := io.ReadAll(body)
// // 	if err != nil {
// // 		return
// // 	}
// // 	fmt.Println(string(all))
// // 	fmt.Println("****************************************")

// // 	//var data1 Data
// // 	//
// // 	//err = json.Unmarshal(all, &data1)
// // 	//if err != nil {
// // 	//	return
// // 	//}
// // 	//
// // 	parsedString := html.EscapeString(string(all))
// // 	log.Println(parsedString)

// // }

// //  CrawlService  -- would traverse the link recursively ,downloads the data  // in backrgound as a go routine
// //  ExtractService -- it will extract the words from the text  //t
// // CleanService
// //SearchService --it will search the data
// // InvIndex map[string]map[string]int

// /*
// *

// 	func clean(host string, href string) string {
// 		// Make sure host ends with slash
// 		if !strings.HasSuffix(host, "/") {
// 			host += "/"
// 		}

// 		// Parse href URL
// 		parsedHref, err := url.Parse(href)
// 		if err == nil && (parsedHref.Scheme == "http" || parsedHref.Scheme == "https") {
// 			// If href is an absolute URL (http/https), return it as is
// 			return parsedHref.String()
// 		}

// 		// use it with host if the href is relative URL
// 		if parsedHref != nil {
// 			return host + strings.TrimPrefix(parsedHref.Path, "/")
// 		}
// 		return host
// 	}
// */
// func iterateText(text string, url string) {
// 	words := strings.Split(text, " ")
// 	for _, word := range words {
// 		// Initialize the inner map if the word is not present
// 		if _, exists := mp[word]; !exists {
// 			mp[word] = make(map[string]int)
// 		}

// 		// Increment the count for the given URL
// 		mp[word][url]++
// 	}
// }

// var mp = map[string]map[string]int{}

// // Extract all links from HTML content
// func extractLinksFromList(html string, baseURL string) []string {
// 	re := regexp.MustCompile(`<li>\s*<a\s+[^>]*href=["']([^"'>]+)["']`)
// 	matches := re.FindAllStringSubmatch(html, -1)

// 	var links []string
// 	for _, match := range matches {
// 		if len(match) > 1 {
// 			link := match[1]
// 			// Convert relative URLs to absolute
// 			if !strings.HasPrefix(link, "http") {
// 				link = baseURL + link
// 			}
// 			links = append(links, link)
// 		}
// 	}
// 	return links
// }
// func getHTMLFromURL(url string) (string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(body), nil
// }
// func solve(links []string, ind int) {
// 	if ind >= len(links) {
// 		return
// 	}
// 	url := links[ind]
// 	ind = ind + 1
// 	text, err := getPlainTextFromURL(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	iterateText(text, url)
// 	fmt.Println(text)
// 	solve(links, ind)
// 	// return

// }

// func extractLinks(url string) ([]string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var links []string
// 	z := html.NewTokenizer(resp.Body)

// 	for {
// 		tt := z.Next()
// 		switch tt {
// 		case html.ErrorToken:
// 			return links, nil
// 		case html.StartTagToken:
// 			t := z.Token()
// 			if t.Data == "a" {
// 				for _, a := range t.Attr {
// 					if a.Key == "href" {
// 						links = append(links, a.Val)
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func main() {
// 	url := "https://usf-cs272-s25.github.io/top10/"

// 	text, err := getPlainTextFromURL(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	iterateText(text, url)
// 	// url := "https://usf-cs272-s25.github.io/top10/Dracula%20%7C%20Project%20Gutenberg/index.html"
// 	htmlContent, err := getHTMLFromURL(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	// fmt.Println(htmlContent)

// 	links := extractLinksFromList(htmlContent, "https://usf-cs272-s25.github.io")

// 	fmt.Println("Extracted Links from Body:", len(links))
// 	for _, link := range links {
// 		fmt.Println(link)
// 	}

// 	// i := 0
// 	// solve(links, i)
// 	// for word, urls := range mp {
// 	// 	fmt.Printf("\"%s\": {\n", word)
// 	// 	for url, count := range urls {
// 	// 		fmt.Printf("\t\"%s\": %d,\n", url, count)
// 	// 	}
// 	// 	fmt.Println("},")
// 	// }

// 	// url := "https://usf-cs272-s25.github.io/top10/Dracula%20%7C%20Project%20Gutenberg/index.html"
// 	// links, err := extractLinks(url)
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// 	return
// 	// }
// 	// fmt.Println("Extracted Links:", links)
// }

// // package main

// // import (
// // 	"fmt"
// // 	"regexp"
// // 	"strings"
// // )

// // func extractLinksFromHTML(html string, baseURL string) []string {
// // 	re := regexp.MustCompile(`<a\s+[^>]*href=["']([^"'>]+)["']`)
// // 	matches := re.FindAllStringSubmatch(html, -1)

// // 	var links []string
// // 	for _, match := range matches {
// // 		if len(match) > 1 {
// // 			link := match[1]
// // 			// Convert relative URLs to absolute
// // 			if !strings.HasPrefix(link, "http") {
// // 				link = baseURL + link
// // 			}
// // 			links = append(links, link)
// // 		}
// // 	}
// // 	return links
// // }

// // func main() {
// // 	// html := `<html><body><a href="/page1">Page 1</a><a href="https://example.com/page2">Page 2</a></body></html>`
// // 	url := "https://usf-cs272-s25.github.io/top10/Dracula%20%7C%20Project%20Gutenberg/index.html"
// // 	htmlContent, err := getHTMLFromURL(url)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 		return
// // 	}
// // 	baseURL := "https://usf-cs272-s25.github.io"
// // 	links := extractLinksFromHTML(html, baseURL)
// // 	fmt.Println("Extracted Links:", links)
// // }
package main

import (
	"fmt"
	"web_s/services"
	"web_s/storage"
)

func main() {
	services.StartCrawl()
	for word, urls := range storage.Mp {
		fmt.Printf("\"%s\": {\n", word)
		for url, count := range urls {
			fmt.Printf("\t\"%s\": %d,\n", url, count)
		}
		fmt.Println("},")
	}
}
