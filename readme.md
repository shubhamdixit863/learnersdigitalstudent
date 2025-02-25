# Go Course Material

## Session 1: Basics

### Problem Set 1.1

1. **Banking Application**
   - Create a banking application that takes the following from the user:
     - Name
     - Age
     - Bank balance

2. **FizzBuzz**
   - Create a function called Fizzbuzz that takes a number as an argument
   - If number is divisible by 3, print "fizz"
   - If number is divisible by 5, print "buzz"
   - If number is divisible by both 3 and 5, print "fizzbuzz"

3. **User Details**
   - Write a program to take user details as input:
     - Name
     - Age
     - Address
   - Convert the age into days
   - Print the complete details

4. **Calculator**
   - Make a calculator application

## Session 2: Functions and Control Flow

### Problem Set 2.1

1. **Swap Two Integers**
   - Write a function that takes two integers and swaps their values without using a third variable

2. **Area of a Circle**
   - Create a program that takes the radius as a float64 and returns the area
   - Hint: Use math.Pi

3. **Convert Celsius to Fahrenheit**
   - Write a function that converts a temperature from Celsius to Fahrenheit

4. **String Concatenation**
   - Accept two strings from the user and concatenate them using a function

5. **Check Data Types**
   - Write a program that takes any variable and prints its type using %T in fmt.Printf

6. **Sum of Two Numbers**
   - Write a function that takes two integers and returns their sum

7. **Variadic Sum Function**
   - Create a variadic function that accepts any number of integers and returns their sum

8. **Max in a Variadic Function**
   - Write a variadic function that returns the maximum value among the provided integers

9. **Palindrome Checker (String)**
   - Write a function that checks whether a given string is a palindrome

10. **Factorial Using Recursion**
    - Implement a recursive function to find the factorial of a given number

11. **Sum of First N Natural Numbers**
    - Use a for loop to calculate the sum of the first N natural numbers

12. **Multiplication Table Generator**
    - Write a program that prints the multiplication table of a given number

13. **Fibonacci Series**
    - Print the first N terms of the Fibonacci series using a for loop

14. **Reverse a String**
    - Use a for loop to reverse a given string

15. **Prime Number Checker**
    - Write a program that checks if a given number is prime using a for loop

16. **Even or Odd**
    - Write a program that checks whether an integer is even or odd

17. **Largest of Three Numbers**
    - Accept three integers and determine the largest using if-else conditions

18. **Leap Year Checker**
    - Create a function that checks if a given year is a leap year

19. **Grade Calculator**
    - Given a score between 0 and 100, assign grades:
      - 90-100: A
      - 80-89: B
      - 70-79: C
      - 60-69: D
      - Below 60: F

20. **Vowel or Consonant**
    - Accept a single character and determine if it is a vowel or a consonant

### Practical 2.1: Quiz Game

Create a Quiz Game in Go that:

1. **Displays a menu with options:**
   - Start Quiz
   - View Score
   - Exit

2. **Start Quiz:**
   - Ask the user 5 multiple-choice questions
   - For each question:
     - Display the question and 4 options
     - Take user input (1-4) for the answer
     - Validate the answer and update the score

3. **View Score:**
   - Display the user's total score

4. **Exit:**
   - Quit the application with a thank-you message

**Concepts Covered:**
- Data types: Integers, strings, booleans
- Functions: Modularize quiz logic
- Variadic function: To validate multiple correct answers
- For loops: To iterate through questions
- If-else: To check answers and handle user choices
- User input: For answers and menu selection

**Menu Example:**
```
=== Quiz Game ===

1. Start Quiz
2. View Score
3. Exit
Enter your choice: 1

Question 1: What is the capital of France?
1. Berlin  2. Madrid
3. Paris   4. Rome
Your answer: 3
Correct!

Question 2: 2 + 2 = ?
1. 3  2. 4
3. 5  4. 6
Your answer: 2
Correct!

=== Quiz Completed! ===
Your Score: 2 out of 5
```

**Functional Requirements:**
- Menu Handling
  - Use a for loop to keep showing the menu until the user exits
- startQuiz() function
  - Ask 5 questions using a for loop
  - Take user input for answers
  - Update the global score variable
- checkAnswer(correctOption int, userAnswer int) bool function
  - Takes the correct option and user answer
  - Returns true if correct, false otherwise
- viewScore() function
  - Displays the total score
- validateInput(allowedOptions ...int) int variadic function
  - Keeps prompting until the user enters a valid option

**Bonus Features:**
- Add a timer for each question (challenge feature)
- Randomize question order
- Add hints for tricky questions

## Session 3: Type Conversion

### Problem Set 3.1

1. **Convert String Inputs to Integers and Calculate Sum:**
   - Write a program that takes five string inputs from the user, converts them to integers, and prints the sum

2. **Fahrenheit to Celsius Converter:**
   - Create a function that converts Fahrenheit to Celsius
   - Take temperature input from the user and display the result with two decimal places

3. **Integer to Float Conversion and Average Calculation:**
   - Take three integer inputs from the user, convert them to floats, and calculate the average using a function

4. **Boolean to Integer Conversion:**
   - Write a function that converts a boolean value to an integer (true -> 1, false -> 0)
   - Take user input as "true" or "false" and convert it

5. **Convert Float to Integer and Check Even/Odd:**
   - Take a float input from the user, convert it to an integer, and determine whether the resulting integer is even or odd

6. **Hexadecimal to Decimal Converter:**
   - Ask the user to enter a hexadecimal string (e.g., "1A"), convert it to an integer, and display the decimal equivalent

7. **Iterate Over a Slice of Strings and Convert to Integers:**
   - Given a slice of string numbers ([]string{"10", "20", "30", "40"}), write a program to convert them into integers and calculate their total sum

8. **String to Float Conversion with Error Handling:**
   - Take a string input representing a floating-point number, convert it to a float, and print the square of the number
   - Handle any conversion errors

9. **Read Multiple Values and Convert Types:**
   - Read an integer, a float, and a string from the user
   - Convert the integer to a float and the float to an integer, then display the converted values

10. **Character to ASCII Value Converter:**
    - Prompt the user to enter a single character, convert it to its ASCII integer value, and print the result

### Practical 3.1: Student Score Management CLI

**Problem Statement:** Build a Student Score Management CLI Application in Golang (Without Structs)

**Overview:**
Create a Command-Line Interface (CLI) application in Golang that allows users to add, view, update, and delete student scores.

1. **Add a Student:**
   - Prompt the user to enter a student name and score
   - Store names in a []string slice and scores in a []int slice
   - Names and scores should be aligned by index (e.g., names[0] corresponds to scores[0])

2. **View All Students and Scores:**
   - Display all students with their scores
   - Handle cases when there are no students

3. **Update a Student's Score:**
   - Ask for the student's name and new score
   - Update the corresponding score in the scores slice

4. **Delete a Student:**
   - Prompt for the student's name and remove both the name and corresponding score

5. **Exit the Application:**
   - Provide an option to terminate the program

**Features to Implement:**
- Use two slices:
  - names []string for storing student names
  - scores []int for storing corresponding scores
- Create functions for:
  - Adding a student (addStudent)
  - Viewing all students (viewStudents)
  - Updating scores (updateScore)
  - Deleting students (deleteStudent)
- Use for loops to:
  - Iterate through slices for display, update, and deletion operations
- Use basic datatypes (string, int) for names and scores

**Example:**
```
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
```

### Practical 3.2: Sublist and Statistics Calculator

**Problem Statement:** Build a Sublist and Statistics Calculator Using the Slicing Operator in Golang

**Overview:**
Create a CLI application in Golang that performs various operations on a list of integers provided by the user. The application should use the slicing operator (:) to create sublists and calculate statistics like sum, average, and maximum of the sliced portion.

**Requirements:**

1. **Input List of Integers:**
   - Prompt the user to enter a list of integers separated by spaces (e.g., 10 20 30 40 50)
   - Store these integers in a slice

2. **Slice the List:**
   - Ask the user to input start and end indices
   - Use the slice operator to create a sublist from the original list
   - Validate that indices are within bounds and that start <= end

3. **Perform Operations on the Sublist:**
   - Display the sliced sublist
   - Calculate and display the sum of the sublist elements
   - Calculate and display the average of the sublist elements
   - Find and display the maximum element in the sublist

4. **Allow Multiple Operations:**
   - After displaying the results, ask if the user wants to perform another slicing operation or exit

**Example:**
```
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
```

## Session 4: Arrays, Slices, and Maps

### Problem Set 4.1

1. **Reverse an Array:**
   - Write a function to reverse an array of integers

2. **Sum of Elements in a Slice:**
   - Create a function that takes a slice of integers and returns the sum of all elements

3. **Find Maximum and Minimum in an Array:**
   - Given an integer array, find and return the maximum and minimum elements

4. **Remove Duplicates from a Slice:**
   - Write a function to remove all duplicate integers from a slice

5. **Rotate Array by k Steps:**
   - Rotate the elements of an array to the right by k steps

6. **Find All Even Numbers in a Slice:**
   - Given a slice of integers, return a new slice containing only even numbers

7. **Count Occurrences of an Element:**
   - Write a function that counts how many times a specific integer appears in a slice

8. **Word Frequency Counter:**
   - Given a string, count the frequency of each word using a map

9. **Character Frequency in a String:**
   - Write a function that returns a map with the count of each character in a given string

10. **Two Sum Problem:**
    - Given an array and a target integer, find the indices of the two numbers that add up to the target using a map

11. **Check for Duplicates in an Array:**
    - Determine if an array contains any duplicate elements using a map

12. **First Unique Character in a String:**
    - Find the index of the first non-repeating character in a string using a map

13. **Fibonacci Sequence Generator:**
    - Write a function that generates the first n Fibonacci numbers using a for loop

14. **Palindrome Number Checker:**
    - Write a function to check if a given integer is a palindrome (without converting it to a string)

15. **Find Missing Number in a Range:**
    - Given an array containing n-1 distinct integers from 1 to n, find the missing number

### Practical 4.1: Flight Booking Analysis

**Problem Statement:** You've been asked to build a tool for analyzing flight bookings for an airline company. The system should process booking data and provide useful information like total revenue per flight, the number of bookings in each seat class, and the most frequent flyers.

**Example Data:**
A list of bookings, where each booking is a string formatted like this:
"PassengerName:FlightNumber:SeatClass:Price"

```go
bookings := []string{
    "Alice:FL123:Economy:120.50",
    "Bob:FL123:Business:450.00",
    "Charlie:FL456:Economy:150.75",
    "Alice:FL456:Economy:150.75",
    "Bob:FL123:Economy:120.50",
}
```

**Tasks:**
1. **Calculate total revenue for each flight**
2. **Count how many bookings there are for each seat class** (Economy, Business, First)
3. **Figure out which passenger has the most bookings**
4. **Find the flight that earned the most revenue**
5. **Ignore invalid bookings**
   - Skip entries that are incorrectly formatted, have negative prices, or contain missing information
6. **Avoid counting duplicate bookings**
   - If the same passenger books the same flight and seat class more than once, it should only count once

**Edge Cases:**
- Bookings with empty strings or wrong formats
- Passengers booking multiple flights
- Flights with no valid bookings

**Expected Output:**
```
Revenue per flight:
FL123: $690.50
FL456: $301.50

Bookings by seat class:
Economy: 4
Business: 1
First: 0

Passenger with the most bookings: Alice (2 bookings)
Flight with the highest revenue: FL123 ($690.50)
```