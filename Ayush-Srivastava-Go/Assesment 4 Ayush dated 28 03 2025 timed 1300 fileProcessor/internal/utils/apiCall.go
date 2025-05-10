package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

//SendAPI sends the data to the API with retires
func SendAPI(data string, retries int) string {

	for attempt := 1; attempt <= retries; attempt++ {
		
		res, err := http.Post(url, ContentType, strings.NewReader(data))

		if err == nil && res.StatusCode == http.StatusOK {
			return APISuccess
		}

		fmt.Println(APIRetying, attempt)
		time.Sleep(2 * time.Second)
	}
	return APIFailed
}