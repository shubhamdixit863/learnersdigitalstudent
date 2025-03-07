package errorss

import (
	"errors"
	"log"
)

func HandleError(err error) {
	var vlderr *ValidationError
	if errors.As(err, &vlderr) {
		log.Println("Validation error occurred:", err)
	} else {
		log.Println("Task error:", err)
	}
}
