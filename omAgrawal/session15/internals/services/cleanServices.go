package services

import (
	"regexp"
	"strings"
)

//func CleanHTML(input string) string {
//	fmt.Println(input)
//	unescaped := html.UnescapeString(input) // Decode entities like &lt; and &gt;
//	fmt.Println(unescaped)
//	re := regexp.MustCompile(`<[^>]*>`) // Match HTML tags
//	cleaned := re.ReplaceAllString(unescaped, "")
//	fmt.Println(cleaned)
//	return strings.ReplaceAll(cleaned, "\n", " ") // Replace newlines with spaces
//}

func CleanHTML(input string) string {
	//unescaped := html.UnescapeString(input) // Decode entities like &lt; and &gt;
	re := regexp.MustCompile(`<[^>]*>`) // Match HTML tags
	cleaned := re.ReplaceAllString(input, "")

	return strings.ReplaceAll(cleaned, "\n", " ") // Replace newlines with spaces
}
