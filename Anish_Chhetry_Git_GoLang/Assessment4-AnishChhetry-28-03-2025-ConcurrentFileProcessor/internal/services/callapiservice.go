package services

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func RetryCallAPI(storage map[string]string, url string) {
	for _, content := range storage {
		for i := int0; i < int3; i++ {
			response, err := CallAPI(url, content)
			if err == nil {
				log.Println(response)
				break
			}
		}
	}
}

func CallAPI(url, data string) (string, error) {
	respose, err := http.Post(url, emptyString, strings.NewReader(data))
	if err != nil {
		return emptyString, err
	}
	body, err := io.ReadAll(respose.Body)
	if err != nil {
		return emptyString, err
	}
	return string(body), nil

}
