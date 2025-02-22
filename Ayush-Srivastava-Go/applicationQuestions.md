# Problem Statement: Build a Student Score Management CLI Application in Golang (Without Structs)

Overview:

Create a Command-Line Interface (CLI) application in Golang that allows users to add, view, update, and delete student scores .

    1.	Add a Student:

    •	Prompt the user to enter a student name and score.

    •	Store names in a []string slice and scores in a []int slice.

    •	Names and scores should be aligned by index (e.g., names[0] corresponds to scores[0]).

    2.	View All Students and Scores:

    •	Display all students with their scores.

    •	Handle cases when there are no students.

    3.	Update a Student’s Score:

    •	Ask for the student’s name and new score.

    •	Update the corresponding score in the scores slice.

    4.	Delete a Student:

    •	Prompt for the student’s name and remove both the name and corresponding score.

    5.	Exit the Application:

    •	Provide an option to terminate the program.

Features to Implement:

    •	Use two slices:

    •	names []string for storing student names.

    •	scores []int for storing corresponding scores.

    •	Create functions for:

    •	Adding a student (addStudent)

    •	Viewing all students (viewStudents)

    •	Updating scores (updateScore)

    •	Deleting students (deleteStudent)

    •	Use for loops to:

    •	Iterate through slices for display, update, and deletion operations.

    •	Use basic datatypes (string, int) for names and scores.

Student Score Manager:

1. Add Student

2. View Students

3. Update Student Score

4. Delete Student

5. Exit

Enter choice: 1

Enter Student Name: Alice

Enter Score: 85

Student added successfully!

Enter choice: 2

Students and Scores:

1. Alice - 85

Enter choice: 3

Enter Student Name to Update: Alice

Enter New Score: 90

Score updated successfully!

Enter choice: 4

Enter Student Name to Delete: Alice

Student deleted successfully!

Enter choice: 2

No students found.

# Problem Statement:

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

1. Berlin 2) Madrid

2. Paris 4) Rome

Your answer: 3

Correct!

Question 2: 2 + 2 = ?

1. 3 2) 4

2. 5 4) 6

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

# Day 3

1. Convert String Inputs to Integers and Calculate Sum:

Write a program that takes five string inputs from the user, converts them to integers, and prints the sum.

2. Fahrenheit to Celsius Converter:

Create a function that converts Fahrenheit to Celsius. Take temperature input from the user and display the result with two decimal places.

3. Integer to Float Conversion and Average Calculation:

Take three integer inputs from the user, convert them to floats, and calculate the average using a function.

4. Boolean to Integer Conversion:

Write a function that converts a boolean value to an integer (true -> 1, false -> 0). Take user input as "true" or "false" and convert it.

5. Convert Float to Integer and Check Even/Odd:

Take a float input from the user, convert it to an integer, and determine whether the resulting integer is even or odd.

6. Hexadecimal to Decimal Converter:

Ask the user to enter a hexadecimal string (e.g., "1A"), convert it to an integer, and display the decimal equivalent.

7. Iterate Over a Slice of Strings and Convert to Integers:

Given a slice of string numbers ([]string{"10", "20", "30", "40"}), write a program to convert them into integers and calculate their total sum.

8. String to Float Conversion with Error Handling:

Take a string input representing a floating-point number, convert it to a float, and print the square of the number. Handle any conversion errors.

9. Read Multiple Values and Convert Types:

Read an integer, a float, and a string from the user. Convert the integer to a float and the float to an integer, then display the converted values.

10. Character to ASCII Value Converter:

Prompt the user to enter a single character, convert it to its ASCII integer value, and print the result.

# Day 3

# Problem Statement: Build a Sublist and Statistics Calculator Using the Slicing Operator in Golang

Overview:

Create a CLI application in Golang that performs various operations on a list of integers provided by the user. The application should use the slicing operator (:) to create sublists and calculate statistics like sum, average, and maximum of the sliced portion.

Requirements:

    1.	Input List of Integers:

    •	Prompt the user to enter a list of integers separated by spaces (e.g., 10 20 30 40 50).

    •	Store these integers in a slice.

    2.	Slice the List:

    •	Ask the user to input start and end indices.

    •	Use the slice operator to create a sublist from the original list.

    •	Validate that indices are within bounds and that start <= end.

    3.	Perform Operations on the Sublist:

    •	Display the sliced sublist.

    •	Calculate and display the sum of the sublist elements.

    •	Calculate and display the average of the sublist elements.

    •	Find and display the maximum element in the sublist.

    4.	Allow Multiple Operations:

    •	After displaying the results, ask if the user wants to perform another slicing operation or exit.

Example :

Enter integers separated by space: 10 20 30 40 50 60

Enter start index: 2

Enter end index: 5

Sliced Sublist: [30 40 50]

Sum: 120

Average: 40.00

Maximum: 50

Do you want to slice again? (yes/no): yes

Enter start index: 0

Enter end index: 3

Sliced Sublist: [10 20 30]

Sum: 60

Average: 20.00

Maximum: 30

Do you want to slice again? (yes/no): no

Exiting the application. Goodbye!
