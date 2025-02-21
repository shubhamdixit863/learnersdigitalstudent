

take user details as input name age company parents name
conver the age into days
print out the whole details

//day 2 20 2 2025

	1.	Swap Two Integers

Write a function that takes two integers and swaps their values without using a third variable.

	2.	Area of a Circle

Create a program that takes the radius as a float64 and returns the area.

Hint: Use math.Pi.

	3.	Convert Celsius to Fahrenheit

Write a function that converts a temperature from Celsius to Fahrenheit.

	4.	String Concatenation

Accept two strings from the user and concatenate them using a function.

	5.	Check Data Types

Write a program that takes any variable and prints its type using %T in fmt.Printf.





	6.	Sum of Two Numbers

Write a function that takes two integers and returns their sum.

	7.	Variadic Sum Function

Create a variadic function that accepts any number of integers and returns their sum.

	8.	Max in a Variadic Function

Write a variadic function that returns the maximum value among the provided integers.

	9.	Palindrome Checker (String)

Write a function that checks whether a given string is a palindrome.

	10.	Factorial Using Recursion

Implement a recursive function to find the factorial of a given number.





	

	12.	Multiplication Table Generator

Write a program that prints the multiplication table of a given number.

	13.	Fibonacci Series

Print the first N terms of the Fibonacci series using a for loop.

	14.	Reverse a String

Use a for loop to reverse a given string.

	15.	Prime Number Checker

Write a program that checks if a given number is prime using a for loop.



	16.	Even or Odd

Write a program that checks whether an integer is even or odd.

	17.	Largest of Three Numbers

Accept three integers and determine the largest using if-else conditions.

	18.	Leap Year Checker

Create a function that checks if a given year is a leap year.

	19.	Grade Calculator

Given a score between 0 and 100, assign grades:

	•	90-100: A

	•	80-89: B

	•	70-79: C

	•	60-69: D

	•	Below 60: F

	20.	Vowel or Consonant

Accept a single character and determine if it is a vowel or a consonant.



This paste expires in <1 hour. Public IP access. Share whatever you see with others in seconds with Context.Terms of ServiceReport this





//pronlem 2
Ÿ“ Problem: Student Management System



Objective:

Build a Student Management System in Go that performs the following tasks:

1. Add Students:

â€¢ Take input for student details: Name (string), Age (int), and Grades (variadic float64).

â€¢ Store the student details in a slice of structs.

2. View All Students:

â€¢ Display all student information in a tabular format.

3. Find Student by Name:

â€¢ Allow the user to search for a student by name and display their details.

4. Calculate Average Grade:

â€¢ Implement a function to calculate the average of a studentâ€™s grades.

5. Find Top Performer:

â€¢ Find and display the student with the highest average grade.

6. Exit:

â€¢ Exit the program.

ðŸ” Requirements

â€¢ Use:

â€¢ Structs for student data.

â€¢ Variadic functions to handle multiple grades input.

â€¢ For loops for iterating over student records.

â€¢ If-else statements for menu options and validations.

â€¢ Functions for modularity (e.g., addStudent, viewStudents, findStudent, calculateAverage, findTopPerformer).

â€¢ Take user input to perform the operations.

ðŸŽ¯ Example Menu (User Interaction)

==== Student Management System ====

1. Add Student

2. View All Students

3. Find Student by Name

4. Find Top Performer

5. Exit

Enter your choice: 1



Enter student name: Alice

Enter student age: 20

Enter grades (separated by space, end with -1): 85 90 95 -1



Student added successfully!



==== Student Management System ====

1. Add Student

2. View All Students

3. Find Student by Name

4. Find Top Performer

5. Exit

Enter your choice: 2



Name: Alice, Age: 20, Grades: [85 90 95], Average: 90.00



==== Student Management System ====

1. Find Top Performer

Enter your choice: 4



Top Performer: Alice with an average grade of 90.00

ðŸ§© Functional Requirements



âœ… addStudent()

â€¢ Takes name, age, and a variadic list of grades as input.

â€¢ Appends the new student to a slice.



âœ… viewStudents()

â€¢ Displays all students with their details and average grade.



âœ… findStudent(name string)

â€¢ Searches for a student by name (case-insensitive).

â€¢ Displays details if found or shows a â€œnot foundâ€ message.



âœ… calculateAverage(grades ...float64) float64

â€¢ Returns the average of provided grades.



âœ… findTopPerformer()

â€¢ Iterates over all students and finds the one with the highest average.

ðŸ› ï¸ Bonus Challenges:

â€¢ Validate that grades are between 0 and 100.

â€¢ Prevent adding students with the same name.

â€¢ Use a map for faster student name lookup.

â€¢ Allow deletion of a student.

