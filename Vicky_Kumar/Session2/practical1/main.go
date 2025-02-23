package main

import (
	"fmt"
	"quiz/startquiz"
	"quiz/viewscore"
)

func main() {
	fmt.Println("_____Welcome to the Quiz Game_____")
	for {
		fmt.Println("Enter 1 to start quiz")
		fmt.Println("Enter 2 to view score")
		fmt.Println("Enter 3 to exit")
		var choice int
		fmt.Scanln(&choice)
		if choice == 1 {
			startquiz.StartQuiz()
		} else if choice == 2 {
			viewscore.ViewScore()
		} else {
			fmt.Println("Thank you for playing the game...!!")
			break
		}
	}
}
