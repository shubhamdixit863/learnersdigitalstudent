// Create a Quiz Game in Go that:

// 1. Displays a menu with options:

// â€¢ Start Quiz

// â€¢ View Score

// â€¢ Exit

// 2. Start Quiz:

// â€¢ Ask the user 5 multiple-choice questions.

// â€¢ For each question:

// â€¢ Display the question and 4 options.

// â€¢ Take user input (1-4) for the answer.

// â€¢ Validate the answer and update the score.

// 3. View Score:

// â€¢ Display the userâ€™s total score.

// 4. Exit:

// â€¢ Quit the application with a thank-you message.

// main we have to use different packages



package main

import (
	"fmt"
	"quiz_application/quiz"
)

func main() {
	
	var choice int
	for {

		fmt.Println("===== Quiz Game =====")
		fmt.Println("1. Start Quiz")
		fmt.Println("2. View Score")
		fmt.Println("3. Exit")
		
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			quiz.Start()
		case 2:
			quiz.ViewScore()
		case 3:
			fmt.Println("Thank you for playing!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

