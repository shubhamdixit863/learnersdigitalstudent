package main

import "fmt"

func swap(a1 int, a2 int) (int, int) {
	a1, a2 = a2, a1
	return a1, a2
}
func area(radius float64)float64{
	return 3.14*radius*radius
}
func fah(c float64)float64{
	return (c*9/5)+32
}
func concat(s1 string,s2 string)string{
	return s1 + s2
}
func sum(a3 int,a4 int)int{
	return a3+a4
}
func variadic(n2 ...int)int{
	var sum=0
	for i:=0;i<len(n2);i++{
		sum+=n2[i]
	}
	return sum
}
func max(n2 ...int)int{
	var maxi int=n2[0]
	for i:=1;i<len(n2);i++{
		if(maxi<n2[i]){
			maxi=n2[i]
		}
	}
	return maxi
}
// func stringcheck(s1 string)string{
// 	var new string;
// 	var j int =0
// 	for i:= len(s1)-1;i>(len(s1)-1)/2;i--{
// 		s1[j]= s1[i]
// 	}
// }
func fact(n1 int)int{
	if(n1==0){
		return 0
	}else if(n1==1){
		return 1
	}
		return n1*fact(n1-1)
}
func table(n1 int){
	for i:=1;i<=10;i++{
		var tb int=n1*i
		fmt.Println(tb)
	}
}
func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		temp := a + b
		a = b
		b = temp
	}
	fmt.Println()
}

func reverseString(s1 string) string {
	var rev string
	for i := len(s1) - 1; i >= 0; i-- {
		rev += string(s1[i])
	}
	return rev
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func evenOdd(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func largestOfThree(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}

func isLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

func gradeCalculator(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func vowelOrConsonant(ch byte) string {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return "Vowel"
	default:
		return "Consonant"
	}
}

func main() {
	c,d := swap(10, 5)
	r:=area(5)
	h:=fah(32)
	s1:=concat("str","ptr")
	fmt.Println(c,d)
	fmt.Println(r,h)
	fmt.Println(s1)
	var i int=43
	fmt.Printf("its type using %T",i)
	s:=sum(5,4)
	fmt.Println(s)
	t:=variadic(1,2,3)
	println(t)
	p:=max(3,6,2,8)
	println(p)
	b:=fact(5)
	println(b)
	table(3)
	fibonacci(10)
	rev := reverseString("hello")
	fmt.Println(rev)
	fmt.Println(isPrime(17))
	fmt.Println(evenOdd(7))
	fmt.Println(largestOfThree(3, 9, 5))
	fmt.Println(isLeapYear(2024))
	fmt.Println(gradeCalculator(85))
	fmt.Println(vowelOrConsonant('a'))

}
// 1.	Swap Two Integers

// Write a function that takes two integers and swaps their values without using a third variable.

// 	2.	Area of a Circle

// Create a program that takes the radius as a float64 and returns the area.

// Hint: Use math.Pi.

// 	3.	Convert Celsius to Fahrenheit

// Write a function that converts a temperature from Celsius to Fahrenheit.

// 	4.	String Concatenation

// Accept two strings from the user and concatenate them using a function.

// 	5.	Check Data Types

// Write a program that takes any variable and prints its type using %T in fmt.Printf.





// 	6.	Sum of Two Numbers

// Write a function that takes two integers and returns their sum.

// 	7.	Variadic Sum Function

// Create a variadic function that accepts any number of integers and returns their sum.

// 	8.	Max in a Variadic Function

// Write a variadic function that returns the maximum value among the provided integers.

// 	9.	Palindrome Checker (String)

// Write a function that checks whether a given string is a palindrome.

// 	10.	Factorial Using Recursion

// Implement a recursive function to find the factorial of a given number.





// 	11.	Sum of First N Natural Numbers

// Use a for loop to calculate the sum of the first N natural numbers.

// 	12.	Multiplication Table Generator

// Write a program that prints the multiplication table of a given number.

// 	13.	Fibonacci Series

// Print the first N terms of the Fibonacci series using a for loop.

// 	14.	Reverse a String

// Use a for loop to reverse a given string.

// 	15.	Prime Number Checker

// Write a program that checks if a given number is prime using a for loop.



// 	16.	Even or Odd

// Write a program that checks whether an integer is even or odd.

// 	17.	Largest of Three Numbers

// Accept three integers and determine the largest using if-else conditions.

// 	18.	Leap Year Checker

// Create a function that checks if a given year is a leap year.

// 	19.	Grade Calculator

// Given a score between 0 and 100, assign grades:

// 	•	90-100: A

// 	•	80-89: B

// 	•	70-79: C

// 	•	60-69: D

// 	•	Below 60: F

// 	20.	Vowel or Consonant

// Accept a single character and determine if it is a vowel or a consonant





// ðŸ“ Problem: Student Management System



// Objective:

// Build a Student Management System in Go that performs the following tasks:

// 1. Add Students:

// â€¢ Take input for student details: Name (string), Age (int), and Grades (variadic float64).

// â€¢ Store the student details in a slice of structs.

// 2. View All Students:

// â€¢ Display all student information in a tabular format.

// 3. Find Student by Name:

// â€¢ Allow the user to search for a student by name and display their details.

// 4. Calculate Average Grade:

// â€¢ Implement a function to calculate the average of a studentâ€™s grades.

// 5. Find Top Performer:

// â€¢ Find and display the student with the highest average grade.

// 6. Exit:

// â€¢ Exit the program.

// ðŸ” Requirements

// â€¢ Use:

// â€¢ Structs for student data.

// â€¢ Variadic functions to handle multiple grades input.

// â€¢ For loops for iterating over student records.

// â€¢ If-else statements for menu options and validations.

// â€¢ Functions for modularity (e.g., addStudent, viewStudents, findStudent, calculateAverage, findTopPerformer).

// â€¢ Take user input to perform the operations.

// ðŸŽ¯ Example Menu (User Interaction)

// ==== Student Management System ====

// 1. Add Student

// 2. View All Students

// 3. Find Student by Name

// 4. Find Top Performer

// 5. Exit

// Enter your choice: 1



// Enter student name: Alice

// Enter student age: 20

// Enter grades (separated by space, end with -1): 85 90 95 -1



// Student added successfully!



// ==== Student Management System ====

// 1. Add Student

// 2. View All Students

// 3. Find Student by Name

// 4. Find Top Performer

// 5. Exit

// Enter your choice: 2



// Name: Alice, Age: 20, Grades: [85 90 95], Average: 90.00



// ==== Student Management System ====

// 1. Find Top Performer

// Enter your choice: 4



// Top Performer: Alice with an average grade of 90.00

// ðŸ§© Functional Requirements



// âœ… addStudent()

// â€¢ Takes name, age, and a variadic list of grades as input.

// â€¢ Appends the new student to a slice.



// âœ… viewStudents()

// â€¢ Displays all students with their details and average grade.



// âœ… findStudent(name string)

// â€¢ Searches for a student by name (case-insensitive).

// â€¢ Displays details if found or shows a â€œnot foundâ€ message.



// âœ… calculateAverage(grades ...float64) float64

// â€¢ Returns the average of provided grades.



// âœ… findTopPerformer()

// â€¢ Iterates over all students and finds the one with the highest average.

// ðŸ› ï¸ Bonus Challenges:

// â€¢ Validate that grades are between 0 and 100.

// â€¢ Prevent adding students with the same name.

// â€¢ Use a map for faster student name lookup.

// â€¢ Allow deletion of a student