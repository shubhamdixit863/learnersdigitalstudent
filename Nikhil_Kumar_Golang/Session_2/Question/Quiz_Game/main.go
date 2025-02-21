/*Create a Quiz Game in Go that:
1. Displays a menu with options:
Start Quiz
View Score
Exit
2. Start Quiz:
Ask the user 5 multiple-choice questions.
For each question:
Display the question and 4 options.
Take user input (1-4) for the answer.
Validate the answer and update the score.
3. View Score:
 Display the userâ€™s total score.
4. Exit:
Quit the application with a thank-you message.
 Concepts Covered:

Data types: Integers, strings, booleans.
Functions: Modularize quiz logic.
Variadic function: To validate multiple correct answers.
For loops: To iterate through questions.
If-else: To check answers and handle user choices.
User input: For answers and menu selection.
Menu Example:
=== Quiz Game ===
1. Start Quiz
2. View Score
3. Exit
Enter your choice: 1

Question 1: What is the capital of France?
1) Berlin   2) Madrid
3) Paris    4) Rome
Your answer: 3
Correct!

Question 2: 2 + 2 = ?
1) 3   2) 4
3) 5   4) 6
Your answer: 2
Correct!

=== Quiz Completed! ===
Your Score: 2 out of 5
 Functional Requirements:
Menu Handling
Use a for loop to keep showing the menu until the user exits.
startQuiz() function
Ask 5 questions using a for loop.
Take user input for answers.
Update the global score variable.
checkAnswer(correctOption int, userAnswer int) bool function
Takes the correct option and user answer.
Returns true if correct, false otherwise.
viewScore() function
Displays the total score.
validateInput(allowedOptions ...int) int variadic function
Keeps prompting until the user enters a valid option.
 Bonus Features:
Add a timer for each question (challenge feature).
Randomize question order.
Add hints for tricky questions.
*/

package main

import (
	"Quiz_Game/display"
	"Quiz_Game/start"
	"fmt"
)
func main()  {
	var answer int
	for{
	fmt.Println(display.Display())
	fmt.Scan(&answer)
	if answer == 1 {
		start.Start()
	}else if answer == 2 {
		fmt.Println(start.Score)
	}else if answer == 3 {
		break
	}
}
}