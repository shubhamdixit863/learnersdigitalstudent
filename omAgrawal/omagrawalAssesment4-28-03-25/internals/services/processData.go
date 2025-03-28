package services

import (
	"bytes"
	"log"
	"net/http"
	"omagrawalAssesment4-28-03-25/internals/models"
	"omagrawalAssesment4-28-03-25/internals/utils"
	"strings"
)

func LineFilter(data *models.ProcessedData) {
	textdata := data.Data

	splitWord := strings.Split(textdata, " ")
	for _, w := range splitWord {
		if w == utils.SplitWord {
			log.Println(data.FileName, "line No. ", data.LineNumber, " line data : ", data.Data)
		}

	}

}
func WordCount(data *models.ProcessedData) {
	textdata := data.Data

	splitWord := strings.Split(textdata, " ")
	WordCountv := 0
	for i, _ := range splitWord {
		WordCountv++
		i = i
	}
	log.Println(data.FileName, "line No. ", data.LineNumber, " Word Count : ", WordCountv)

}

func PostToApi(data *models.ProcessedData) {
	textdata := data.Data
	d := bytes.NewBuffer([]byte(textdata))

	post, err := http.Post(utils.ApiPath, "text/plain", d)
	if err != nil {
		return
	}
	log.Println(data.FileName, "line No. ", data.LineNumber, " Post Response ", post)

}

func RetryPostToApi(data *models.ProcessedData) {
	textdata := data.Data
	d := bytes.NewBuffer([]byte(textdata))
	Try := 0
	for Try < 3 {
		post, err := http.Post(utils.ApiPath, "text/plain", d)
		if err != nil {
			Try++
		} else {
			log.Println(data.FileName, "line No. ", data.LineNumber, " Post Response ", post)
			return
		}

	}

}
