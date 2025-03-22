package main

import (
	"fmt"
	"quiz/exit"
	"quiz/questions"
)

func main() {

	for {
		fmt.Println("\n===Start Quiz===")
		fmt.Println("1. Start Quiz")
		fmt.Println("2. View Score")
		fmt.Println("3. Exit")
		fmt.Println("Enter your choice: ")

		choice := questions.Validate(1,2,3)

		switch choice{
		case 1:
			questions.StartQuiz()
		case 2:
			questions.ViewScore()
		case 3:
			exit.SayExit()
		}
	}

}