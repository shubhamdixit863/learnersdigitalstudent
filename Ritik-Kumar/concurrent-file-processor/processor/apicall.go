package processor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// APICall sends each line to an external API
func APICall(linesCh <-chan string) {
	for line := range linesCh {
		resp, err := http.Post("https://httpbin.org/post", "text/plain", bytes.NewBuffer([]byte(line)))
		if err != nil {
			fmt.Println("API Request Failed:", err)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("API Response:", string(body))
		resp.Body.Close()
	}
}
