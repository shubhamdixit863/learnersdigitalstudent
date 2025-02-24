package main

import (
	"fmt"
	"quiz-game/quiz"
)

func main() {
	for {
		fmt.Println("\nQuiz Game")
		fmt.Println("1. Start Quiz")
		fmt.Println("2. View Score")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			c:=quiz.StartQuiz()
			fmt.Println("Score is" ,c)
		case 2:
			// quiz.ViewScore()
		case 3:
			fmt.Println("Thank you for playing!")
			return
		}
	}
}


