package startquiz

import (
	"fmt"
	"quiz/question"
	"quiz/viewscore"
)

func StartQuiz() {
	for i := 0; i < len(question.Question); i++ {
		printQuestion(i+1, question.Question[i], question.Options[4*i], question.Options[4*i+1], question.Options[4*i+2], question.Options[4*i+3])
		var userAnswer int
		fmt.Println("Your answer: ")
		fmt.Scanln(&userAnswer)
		correct := checkAnswer(userAnswer, question.CorrectAnswers[i])
		if correct {
			fmt.Println("Correct!")
			viewscore.Total = viewscore.Total + 1
		} else {
			fmt.Println("Wrong!")
		}
	}
	fmt.Println("Quiz Completed!")

}

func printQuestion(questionNo int, question string, option1 string, option2 string, option3 string, option4 string) {
	fmt.Printf("Question %d: %s\n", questionNo, question)
	fmt.Println("1)", option1, "\t\t", "2)", option2)
	fmt.Println("3)", option3, "\t\t", "4)", option4)
}
func checkAnswer(userAnswer int, correctAnswer int) bool {
	return userAnswer == correctAnswer
}

// printQuestion(1, "In what year did the Great October Socialist Revolution take place?", )
// var userAnswer int
// fmt.Scanln(&userAnswer)
// for {
// 	if userAnswer < 0 || userAnswer > 4 {
// 		fmt.Scanln(&userAnswer)
// 	} else {
// 		break
// 	}
// }
// correct := checkAnswer(userAnswer, 1)
// if correct {
// 	fmt.Println("Correct!")
// 	Total = Total + 1
// }
// printQuestion(2, "What is the largest lake in the world?", )
// fmt.Scanln(&userAnswer)
// for {
// 	if userAnswer < 0 || userAnswer > 4 {
// 		fmt.Scanln(&userAnswer)
// 	} else {
// 		break
// 	}
// }
// correct = checkAnswer(userAnswer, 1)
// if correct {
// 	Total = Total + 1
// }
