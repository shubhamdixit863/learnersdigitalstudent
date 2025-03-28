package services

import (
	"bytes"
	"log"
	"net/http"
	"omagrawalAssesment4-28-03-25/internals/models"
	"omagrawalAssesment4-28-03-25/internals/utils"
	"strings"
)

const (
	LineFIlterResponse = "File Name : %v Line Number %d  Line Data %s \n"
	WordCountPrint     = "File Name : %v Line Number %d  WordCount %s \n"
	ApiREsponsePrint   = "File Name : %v Line Number %d Statuc Code %s \n"
)

func LineFilter(data *models.ProcessedData) {
	textdata := data.Data

	splitWord := strings.Split(textdata, " ")
	for _, w := range splitWord {
		if w == utils.SplitWord {
			//log.Println(data.FileName, "line No. ", data.LineNumber, " line data : ", data.Data)
			log.Printf(LineFIlterResponse, data.FileName, data.LineNumber, data.Data)
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
	log.Printf(WordCountPrint, data.FileName, data.LineNumber, WordCountv)
}
func PostToApi(data *models.ProcessedData) {
	textdata := data.Data
	d := bytes.NewBuffer([]byte(textdata))

	post, err := http.Post(utils.ApiPath, "text/plain", d)
	if err != nil {
		return
	}
	log.Printf(ApiREsponsePrint, data.FileName, data.LineNumber, post.Status)
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
			log.Printf(ApiREsponsePrint, data.FileName, data.LineNumber, post.Status)
			return
		}
	}
}
