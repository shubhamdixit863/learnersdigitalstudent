package services

import (
	"fmt"
	"io"
	"net/http"
	"session15/storage"
	"strings"
)

func CrawlServices(baseurl string) {
	//storage.VisitedUrl[baseurl] = 1
	defer storage.Wg.Done()
	data, err := http.Get(baseurl)

	if err != nil {
		fmt.Println("Error fetching data:", err)
		//return
	}
	defer data.Body.Close() // Ensures the response body is closed properly

	body, err := io.ReadAll(data.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		//return
	}
	StringData := string(body)
	ExtractHrefs(StringData, baseurl)
	splited := strings.Split(CleanHTML(StringData), " ")
	fmt.Println(splited)
	SearchServices(splited, baseurl)
	//fmt.Println(string(body))
}

//723.63
