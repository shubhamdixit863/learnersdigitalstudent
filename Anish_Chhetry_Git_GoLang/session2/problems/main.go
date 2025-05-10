package main

import "fmt"

func main() {
	// 1.	Swap Two Integers
	// Write a function that takes two integers and swaps their values without using a third variable.
	swapTwoIntegers(1, 2)

	// 2.	Area of a Circle
	// Create a program that takes the radius as a float64 and returns the area.
	// Hint: Use math.Pi.
	areaOfCircle(5)

	//  3. Convert Celsius to Fahrenheit
	// Write a function that converts a temperature from Celsius to Fahrenheit.
	celsiusToFahrenheit(0)

	//  4. String Concatenation
	// Accept two strings from the user and concatenate them using a function.
	stringConcatenation("hello", " world")

	//  5. Check Data Types
	// Write a program that takes any variable and prints its type using %T in fmt.Printf.
	checkDataType(10)

	//  6. Sum of Two Numbers
	// Write a function that takes two integers and returns their sum.
	sumOfTwoNumbers(5, 10)

	//  7. Variadic Sum Function
	// Create a variadic function that accepts any number of integers and returns their sum.
	sumOfVariadicNumbers(1, 2, 3, 4, 5)

	//  8. Max in a Variadic Function
	// Write a variadic function that returns the maximum value among the provided integers.
	maxInVariadicFunction(1, 2, 3, 4, 5)

	//  9. Palindrome Checker (String)
	// Write a function that checks whether a given string is a palindrome.
	palindrome("malayalam")

	// 	10.	Factorial Using Recursion
	// Implement a recursive function to find the factorial of a given number.
	factorial(5)

	// 	11.	Sum of First N Natural Numbers
	// Use a for loop to calculate the sum of the first N natural numbers.
	sumOfFirstNNaturalNumbers(10)

	// 	12.	Multiplication Table Generator
	// Write a program that prints the multiplication table of a given number.
	multiplicationTable(5)

	// 	13.	Fibonacci Series
	// Print the first N terms of the Fibonacci series using a for loop.
	fibonacciSeries(10)

	// 	14.	Reverse a String
	// Use a for loop to reverse a given string.
	reverseString("hello")

	// 	15.	Prime Number Checker
	// Write a program that checks if a given number is prime using a for loop.
	isPrime(6)

	// 	16.	Even or Odd
	// Write a program that checks whether an integer is even or odd.
	evenodd(3)

	// 	17.	Largest of Three Numbers
	// Accept three integers and determine the largest using if-else conditions.
	largestOfThree(5, 15, 15)

	// 	18.	Leap Year Checker
	// Create a function that checks if a given year is a leap year.
	leapYear(2020)

	// 	19.	Grade Calculator
	// Given a score between 0 and 100, assign grades:
	// 	•	90-100: A
	// 	•	80-89: B
	// 	•	70-79: C
	// 	•	60-69: D
	// 	•	Below 60: F
	gradeCalculator(85)

	// 	20.	Vowel or Consonant
	// Accept a single character and determine if it is a vowel or a consonant.
	vowelOrConsonant("a")

}

func swapTwoIntegers(a int, b int) {
	fmt.Println("before swap", a, b)
	a, b = b, a
	fmt.Println("after swap", a, b)
}
func areaOfCircle(radius float64) {
	fmt.Println(3.14 * radius * radius)
}

func celsiusToFahrenheit(celsius float64) {
	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Println(fahrenheit)
}
func stringConcatenation(str1 string, str2 string) {
	fmt.Println(str1 + str2)
}
func checkDataType(data interface{}) {
	fmt.Printf("%T\n", data)
}
func sumOfTwoNumbers(a int, b int) {
	fmt.Println(a + b)
}
func sumOfVariadicNumbers(numbers ...int) {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println(sum)
}
func maxInVariadicFunction(numbers ...int) {
	max := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	fmt.Println(max)
}
func palindrome(str string) {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	if str == reversed {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
func factorial(num int) {
	fact := 1
	for i := num; i > 0; i-- {
		fact *= i
	}
	fmt.Println(fact)
}
func sumOfFirstNNaturalNumbers(num int) {
	sum := 0
	for i := num; i > 0; i-- {
		sum += i
	}
	fmt.Println(sum)
}
func multiplicationTable(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
}
func fibonacciSeries(num int) {
	a, b := 0, 1
	for i := 0; i < num; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
}
func reverseString(str string) {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	fmt.Println("\n", reversed)
}
func isPrime(num int) {
	if num <= 1 {
		fmt.Println(false)
		return
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)
}
func evenodd(num int) {
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
func largestOfThree(n1 int, n2 int, n3 int) {
	if n1 >= n2 && n1 >= n3 {
		fmt.Println(n1)
	} else if n2 >= n1 && n2 >= n3 {
		fmt.Println(n2)
	} else if n3 >= n1 && n3 >= n2 {
		fmt.Println(n3)
	}
}
func leapYear(year int) {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
func gradeCalculator(marks int) {
	if marks >= 90 {
		fmt.Println("A")
	} else if marks >= 80 {
		fmt.Println("B")
	} else if marks >= 70 {
		fmt.Println("C")
	} else if marks >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
func vowelOrConsonant(char string) {
	if char == "A" || char == "E" || char == "I" || char == "O" || char == "U" || char == "a" || char == "e" || char == "i" || char == "o" || char == "u" {
		fmt.Println("Vowel")
	} else {
		fmt.Println("Consonant")
	}
}
