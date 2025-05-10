package validate

import (
	"assignment2/answers"
	"assignment2/score"
	"fmt"
)

func ValidateInput(allowedOptions ...int) {
	if allowedOptions[0] == 1 || allowedOptions[0] == 2 || allowedOptions[0] == 3 || allowedOptions[0] == 4 {
		correctOrNot(answers.CheckAnswer(allowedOptions[0], allowedOptions[1]))
	} else {
		fmt.Printf("\nInvalid Option\n")
	}
}
func correctOrNot(a bool) {
	if a {
		fmt.Printf("\nCorrect\n")
		score.IncrementScore()
	} else {
		fmt.Printf("\nIncorrect\n")
	}
}
