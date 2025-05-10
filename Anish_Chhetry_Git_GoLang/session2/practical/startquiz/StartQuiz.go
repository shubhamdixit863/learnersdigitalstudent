package startquiz

import (
	"assignment2/questions"
	"fmt"
)

func Startquiz() {
	fmt.Printf("\nWelcome to the quiz!\n")
	questions.Question(1)
	questions.Question(2)
	questions.Question(3)
	questions.Question(4)
	questions.Question(5)
}
