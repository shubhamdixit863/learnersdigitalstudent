//// package main
////
//// import (
////
////	"fmt"
////	"html"
////	"io"
////	"net/http"
////	"net/url"
////	"regexp"
////	"strings"
////
//// )
////
//// //https://usf-cs272-s25.github.io/top10/
////
////	type Data struct {
////		UserId int    `json:"userId"`
////		Title  string `json:"title"`
////	}
////
////	func main() {
////		baseurl := "https://usf-cs272-s25.github.io/top10/"
////		fmt.Println(baseurl)
////		data, err := http.Get("https://usf-cs272-s25.github.io/top10/")
////		if err != nil {
////			return
////		}
////
////		body := data.Body
////
////		all, err := io.ReadAll(body)
////		if err != nil {
////			return
////		}
////
////		//var data1 Data
////		//
////		//err = json.Unmarshal(all, &data1)
////		//if err != nil {
////		//	return
////		//}
////		//
////		//parsedString := html.EscapeString(string(all))
////
////		a := extractHrefs(string(all))
////		fmt.Println(a)
////		allString := cleanHTML(string(all))
////		vacde := strings.ToLower(allString)
////		//fmt.Println(vacde)
////
////		mkk := strings.Split(vacde, " ")
////		//fmt.Println(mkk)
////
////		var ma = make(map[string]map[string]int)
////
////		for _, v := range mkk {
////			ma[v]["https://usf-cs272-s25.github.io/top10/"] = ma[v]["https://usf-cs272-s25.github.io/top10/"] + 1
////		}
////
////		//fmt.Println(ma)
////
//// }
////
//// //	func removeHTMLTags(input string) string {
//// //		re := regexp.MustCompile(`<[^>]*>`)
//// //
//// //		return re.ReplaceAllString(input, "")
//// //	}
////
////	func cleanHTML(input string) string {
////		//fmt.Println(input)
////		unescaped := html.UnescapeString(input) // Decode entities like &lt; and &gt;
////		re := regexp.MustCompile(`<[^>]*>`)     // Match HTML tags
////		cleaned := re.ReplaceAllString(unescaped, "")
////		return strings.ReplaceAll(cleaned, "\n", " ") // Replace newlines with spaces
////	}
////
////	func extractHrefs(input string) []string {
////		re := regexp.MustCompile(`href=["']([^"']+)["']`)
////		matches := re.FindAllStringSubmatch(input, -1)
////
////		var hrefs []string
////		for _, match := range matches {
////			hrefs = append(hrefs, match[1]) // Extract captured group
////		}
////		return hrefs
////	}
////
//// //  CrawlService  -- would traverse the link recursively ,downloads the data  // in backrgound as a go routine
//// //  ExtractService -- it will extract the words from the text  //t
//// // CleanService
//// //SearchService --it will search the data
//// // InvIndex map[string]map[string]int
////
//// //var mp = map[string]map[string]int{
//// //	"dracula": {
//// //		"http://one.com": 2,
//// //		"http://uo.co,":  78,
//// //	},
//// //}
////
////	func clean(host string, href string) string {
////		// Make sure host ends with slash
////		if !strings.HasSuffix(host, "/") {
////			host += "/"
////		}
////
////		// Parse href URL
////		parsedHref, err := url.Parse(href)
////		if err == nil && (parsedHref.Scheme == "http" || parsedHref.Scheme == "https") {
////			// If href is an absolute URL (http/https), return it as is
////			return parsedHref.String()
////		}
////
////		// use it with host if the href is relative URL
////		if parsedHref != nil {
////			return host + strings.TrimPrefix(parsedHref.Path, "/")
////		}
////		return host
////	}
////
//// //
//package main
//
//import (
//	"fmt"
//	"html"
//	"io"
//	"net/http"
//	"regexp"
//	"strings"
//)
//
//func main() {
//	baseurl := "https://usf-cs272-s25.github.io/top10/"
//	fmt.Println("Base URL:", baseurl)
//
//	data, err := http.Get(baseurl)
//	if err != nil {
//		fmt.Println("Error fetching data:", err)
//		return
//	}
//	defer data.Body.Close() // Ensures the response body is closed properly
//
//	body, err := io.ReadAll(data.Body)
//	if err != nil {
//		fmt.Println("Error reading response body:", err)
//		return
//	}
//
//	// Extract and display hrefs
//	links := extractHrefs(string(body), baseurl)
//	fmt.Println("Extracted HREFs:")
//	for _, link := range links {
//		fmt.Println(link)
//	}
//
//	// Clean HTML text
//	allString := cleanHTML(string(body))
//	vacde := strings.ToLower(allString)
//
//	mkk := strings.Fields(vacde) // Improved: Splits on spaces and handles multiple spaces
//	var ma = make(map[string]map[string]int)
//
//	for _, v := range mkk {
//		if ma[v] == nil {
//			ma[v] = make(map[string]int)
//		}
//		ma[v][baseurl]++
//	}
//
//	// Example: Display word counts
//	fmt.Println("\nSample Word Count:")
//	for word, count := range ma {
//		fmt.Printf("%s: %d\n", word, count[baseurl])
//		if count[baseurl] > 5 {
//			break // Display only a sample
//		}
//	}
//}
//
//// cleanHTML removes HTML tags, decodes entities, and replaces newlines with spaces
//func cleanHTML(input string) string {
//	unescaped := html.UnescapeString(input) // Decode entities like &lt; and &gt;
//	re := regexp.MustCompile(`<[^>]*>`)     // Match HTML tags
//	cleaned := re.ReplaceAllString(unescaped, "")
//	return strings.ReplaceAll(cleaned, "\n", " ") // Replace newlines with spaces
//}
//
//// extractHrefs extracts all href values and appends base URL for relative links
//func extractHrefs(input string, baseurl string) []string {
//	re := regexp.MustCompile(`href=["']([^"']+)["']`)
//	matches := re.FindAllStringSubmatch(input, -1)
//
//	var hrefs []string
//	for _, match := range matches {
//		link := match[1]
//		// Handle relative URLs
//		if !strings.HasPrefix(link, "http") {
//			link = baseurl + link
//		}
//		hrefs = append(hrefs, link)
//	}
//	return hrefs
//}

package main

import (
	"fmt"
	"session15/internals/services"
	"session15/storage"
	"time"
)

func main() {
	T := time.Now()
	//baseurl := "https://usf-cs272-s25.github.io/top10/The Project Gutenberg eBook of The Picture of Dorian Gray, by Oscar Wilde/index.html"
	baseurl := "https://usf-cs272-s25.github.io/top10/"
	//storage.UrlSlice
	storage.Wg.Add(1)
	services.CrawlServices(baseurl)

	fmt.Println("-------------000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println(storage.UrlSlice)
	for _, v := range storage.UrlSlice {
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		storage.Wg.Add(1)
		go services.CrawlServices(v)

	}
	storage.Wg.Wait()
	fmt.Println(storage.UrlSlice)

	//services.CrawlServices(storage.UrlSlice)
	fmt.Println("-------------000000000000000000000000000000000000000000000000000000000000000000000000")

	for i, v := range storage.Mappp {
		for c, x := range v {
			fmt.Println(i, c, x)
		}
	}

	fmt.Println("-------------000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println(time.Since(T))

}
