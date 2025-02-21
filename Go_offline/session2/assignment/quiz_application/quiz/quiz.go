
// startQuiz() function

// Ask 5 questions using a for loop.

// Take user input for answers.

// Update the global score variable.



// checkAnswer(correctOption int, userAnswer int) bool function

// Takes the correct option and user answer.

// Returns true if correct, false otherwise.



// viewScore() function

// Displays the total score.



// validateInput(allowedOptions ...int) int variadic function

// Keeps prompting until the user enters a valid option.

// Bonus Features:

// Add a timer for each question (challenge feature).

// Randomize question order.

// Add hints for tricky questions.



package quiz

import (
	"fmt"
	"math/rand"
	
)

var score int

func Start() {
	score = 0
	
	for i := 0; i < 5; i++ {
		question, options, correctOption := getQuestion()
		fmt.Println(question)
		for i, option := range options {
			fmt.Printf("%d. %s\n", i+1, option)
		}
		var userAnswer int
		fmt.Scanln(&userAnswer)
		
		if checkAnswer(correctOption, userAnswer) {
			fmt.Printf("Your answer is %d \n", userAnswer)
		    fmt.Println("Correct!")
			score++
		}else{
			fmt.Printf("Your answer is %d \n", userAnswer)
			fmt.Println("Incorrect!")
		}
	}
}

func getQuestion() (string, []string, int) {
	
	questions := []string{
		"What is the capital of France?",
		"Who is the author of the Harry Potter series?",
		"Which planet is known as the Red Planet?",
		"Which country is known as the Land of the Rising Sun?",
		"Which is the largest mammal?",
	}
	
	options := [][]string{
		{"Paris", "Berlin", "London", "Madrid"},
		{"J.K. Rowling", "J.R.R. Tolkien", "George R.R. Martin", "Stephen King"},
		{"Mars", "Venus", "Jupiter", "Saturn"},
		{"Japan", "China", "Korea", "Vietnam"},
		{"Blue Whale", "Elephant", "Giraffe", "Hippopotamus"},
	}
	
	correctOptions := []int{1, 1, 1, 1, 1}
	
	randomIndex := rand.Intn(len(questions))
	
	return questions[randomIndex], options[randomIndex], correctOptions[randomIndex]
}

func checkAnswer(correctOption int, userAnswer int) bool {
	return correctOption == userAnswer
}


func ViewScore() {
	fmt.Printf("Your score is %d/5\n",score)

}


