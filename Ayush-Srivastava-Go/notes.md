# Q.1

You have to create a banking application that takes the following things from the user

1.  Name
2.  Age
3.  Bank Balance

# Q.2

Write a function called fizzbuzz
if num is div by 3 print fizz
if num is div by 5 print buzz
if num is div by both print fizzbuzz

# Q.3

We will take input from the command line and store it into a variable

# Q.4

Write a program to take the user details as input convert the age into days Print out the details
Name age address company parents name

# Q.5

Create a calculator

# Day 2

1. Swap Two Integers

Write a function that takes two integers and swaps their values without using a third variable.

2. Area of a Circle

Create a program that takes the radius as a float64 and returns the area.

Hint: Use math.Pi.

3. Convert Celsius to Fahrenheit

Write a function that converts a temperature from Celsius to Fahrenheit.

4. String Concatenation

Accept two strings from the user and concatenate them using a function.

5. Check Data Types

Write a program that takes any variable and prints its type using %T in fmt.Printf.

6.  Sum of Two Numbers

Write a function that takes two integers and returns their sum.

7. Variadic Sum Function

Create a variadic function that accepts any number of integers and returns their sum.

8. Max in a Variadic Function

Write a variadic function that returns the maximum value among the provided integers.

9.  Palindrome Checker (String)

Write a function that checks whether a given string is a palindrome.

10. Factorial Using Recursion

Implement a recursive function to find the factorial of a given number.

11. Sum of First N Natural Numbers

Use a for loop to calculate the sum of the first N natural numbers.

12. Multiplication Table Generator

Write a program that prints the multiplication table of a given number.

13. Fibonacci Series

Print the first N terms of the Fibonacci series using a for loop.

14. Reverse a String

Use a for loop to reverse a given string.

15. Prime Number Checker

Write a program that checks if a given number is prime using a for loop.

16. Even or Odd

Write a program that checks whether an integer is even or odd.

17. Largest of Three Numbers

Accept three integers and determine the largest using if-else conditions.

18. Leap Year Checker

Create a function that checks if a given year is a leap year.

19. Grade Calculator

Given a score between 0 and 100, assign grades:

    •	90-100: A

    •	80-89: B

    •	70-79: C

    •	60-69: D

    •	Below 60: F

20. Vowel or Consonant

Accept a single character and determine if it is a vowel or a consonant.

21. Problem Statement:

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
