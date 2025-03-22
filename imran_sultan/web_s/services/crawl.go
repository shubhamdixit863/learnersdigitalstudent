package services

import (
	"fmt"
	"web_s/storage"
)

func Solve(links []string, ind int) {
	if ind >= len(links) {
		return
	}
	url := links[ind]
	ind = ind + 1
	text, err := GetPlainTextFromURL(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	IterateText(text, url)
	fmt.Println(text)
	Solve(links, ind)

}
func StartCrawl() {
	WordTolink("https://usf-cs272-s25.github.io/top10/")
	Extract("https://usf-cs272-s25.github.io/top10/")
	i := 0
	Solve(storage.Links, i)

}
