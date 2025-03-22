package main

import (
	"io"
	"net/http"
)

func main() {
	data, err := http.Get("")
	if err!=nil{
		return
	}
	body:=data.Body

	all,err:= io.ReadAll(body)
	if err!=nil{
		return
	}

	html.EscapeString(string(all))

}

// use html package to get string


// three services
// crawl service : traverse the link recursively and downloads the data :  search url and put it in queue, if it is traversed not visit again
// extract service:  extract the words from the text  storing text: Inverted index map[string]map[string]int
// clean service
// search service