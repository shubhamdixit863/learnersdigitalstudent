package mode

import (
	"concurrentFileProcessor/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func disp() {
	log.Println(utils.SELECT_MODE)
	log.Println(utils.LINE_FILTER)
	log.Println(utils.WORD_FILTER)
	log.Println(utils.API_CALL)
	log.Println(utils.RETRYABLE_API_CALL)
}

func Mode(linesChan chan string) {

	disp()

	var mode int
	fmt.Scanln(&mode)

	var lines []string
	for line := range linesChan {
		lines = append(lines, line)
	}

	switch mode {
	case 1:
		option_one(lines)
	case 2:
		option_two(lines)
	case 3:
		option_three(lines)
	case 4:
		option_four(lines)
	default:
		log.Println(utils.INV_MODE)
	}
}

func option_one(lines []string) {
	log.Println(utils.ENTER_KEYWORD)
	var keyword string
	fmt.Scanln(&keyword)
	for _, line := range lines {
		if strings.Contains(line, keyword) {
			log.Println(line)
		}
	}

}

func option_two(lines []string) {
	wordCount := 0
	for _, line := range lines {
		wordCount += len(strings.Fields(line))
	}
	log.Println(utils.TOTAL_WORDS, wordCount)
}

func option_three(lines []string) {
	for _, line := range lines {
		resp, err := http.Post(utils.API_URL, utils.CONTANT_TYPE, strings.NewReader(line))
		if err != nil {
			log.Println(utils.API_ERR, err)
		} else {
			log.Println(utils.API_RES, resp.Status)
			resp.Body.Close()
		}
	}
}

func option_four(lines []string) {
	for _, line := range lines {
		var resp *http.Response
		var err error
		maxRetries := 3

		for i := 0; i < maxRetries; i++ {
			resp, err = http.Post(utils.API_URL, utils.CONTANT_TYPE, strings.NewReader(line))
			if err == nil {
				log.Println(utils.API_RES, resp.Status)
				resp.Body.Close()
				break
			}

			log.Printf(utils.RETRY_API, i+1, line, err)
			time.Sleep(time.Duration(1<<i) * time.Second)
		}

		if err != nil {
			log.Println(utils.FAIL_RETRY, err)
		}
	}
}
