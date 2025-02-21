
ðŸŽ¯ Problem Statement:



Create a Quiz Game in Go that:

1. Displays a menu with options:

â€¢ Start Quiz

â€¢ View Score

â€¢ Exit

2. Start Quiz:

â€¢ Ask the user 5 multiple-choice questions.

â€¢ For each question:

â€¢ Display the question and 4 options.

â€¢ Take user input (1-4) for the answer.

â€¢ Validate the answer and update the score.

3. View Score:

â€¢ Display the userâ€™s total score.

4. Exit:

â€¢ Quit the application with a thank-you message.

ðŸ’¡ Concepts Covered:



âœ… Data types: Integers, strings, booleans.

âœ… Functions: Modularize quiz logic.

âœ… Variadic function: To validate multiple correct answers.

âœ… For loops: To iterate through questions.

âœ… If-else: To check answers and handle user choices.

âœ… User input: For answers and menu selection.

ðŸ“ Menu Example:

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

ðŸ”¥ Functional Requirements:



âœ… Menu Handling

â€¢ Use a for loop to keep showing the menu until the user exits.



âœ… startQuiz() function

â€¢ Ask 5 questions using a for loop.

â€¢ Take user input for answers.

â€¢ Update the global score variable.



âœ… checkAnswer(correctOption int, userAnswer int) bool function

â€¢ Takes the correct option and user answer.

â€¢ Returns true if correct, false otherwise.



âœ… viewScore() function

â€¢ Displays the total score.



âœ… validateInput(allowedOptions ...int) int variadic function

â€¢ Keeps prompting until the user enters a valid option.

ðŸš¦ Bonus Features:

â€¢ Add a timer for each question (challenge feature).

â€¢ Randomize question order.

â€¢ Add hints for tricky questions.



This paste expires in <1 hour. Public IP access. Share whatever you see with others in seconds with Context.Terms of ServiceReport this