package main

import (
	"fmt"
	"math"
)

func main() {
	swap()
	areaOfCircle()
	CelsiustoFahrenheit()
	stringConcatenation()
	datatype()
	sum()
	fmt.Println(" variadic function sum ")
	variadicSum(1, 2, 2, 4, 5, 6)
	fmt.Println("variadic function max")
	variadicMAx(1, 3, 5, 7, 8, 10)
	palindromeString()

	fmt.Println("enter number for factorial")
	var a int
	fmt.Scanln(&a)
	b := factorial(a)
	fmt.Println("factorial : ", b)

	SumN()
	multiplication()
	Fibonacci()
	reverseString()
	PrimeNumber()
	evenOdd()
	largestThree()
	leapYear()
	grade()
	vowel()
}

// 1.	Swap Two Integers

// Write a function that takes two integers and swaps their values without using a third variable.

func swap() {
	fmt.Println("Enter number to swap ")
	var num1 int
	var num2 int
	fmt.Scanln(&num1, &num2)
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println("swaped numbers : ", num1, num2)
}

// Area of a Circle

// Create a program that takes the radius as a float64 and returns the area.

func areaOfCircle() {
	fmt.Println("Enter radius")
	var r float32
	fmt.Scanln(&r)
	var area float32
	area = math.Pi * r * r
	fmt.Println("Area : ", area)
}

// Convert Celsius to Fahrenheit

// Write a function that converts a temperature from Celsius to Fahrenheit

func CelsiustoFahrenheit() {
	fmt.Println("Enter value to convert")
	var a float32
	fmt.Scanln(&a)
	var b float32
	b = ((9 * a) / 5) + 32
	fmt.Println("fahrenheit : ", b)
}

// String Concatenation

// Accept two strings from the user and concatenate them using a function.

func stringConcatenation() {
	fmt.Println("enter first string ")
	var s1 string
	fmt.Scanln(&s1)
	var s2 string
	fmt.Println("enter second string")
	fmt.Scanln(&s2)
	fmt.Printf("%s %s", s1, s2)
}

// 5.	Check Data Types

// Write a program that takes any variable and prints its type using %T in fmt.Printf.

func datatype() {
	var a int
	var b string
	var c bool
	var d float32
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}

// Sum of Two Numbers

// Write a function that takes two integers and returns their sum.

func sum() {
	fmt.Println("enter first number to add ")
	var a int
	fmt.Scanln(&a)
	fmt.Println("enter second number to add ")
	var b int
	fmt.Scanln(&b)
	fmt.Println("ans : ", a+b)
}

// Variadic Sum Function

// Create a variadic function that accepts any number of integers and returns their sum.

func variadicSum(n ...int) {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	fmt.Println("ans : ", sum)
}

// Max in a Variadic Function

// Write a variadic function that returns the maximum value among the provided integers.

func variadicMAx(n ...int) {
	var a int
	a = n[0]
	for i := 0; i < len(n); i++ {
		if a < n[i] {
			a = n[i]
		}
	}
	fmt.Println("maximum number : ", a)
}

// Palindrome Checker (String)

// Write a function that checks whether a given string is a palindrome.

func palindromeString() {
	fmt.Println("enter string to check palindrome")
	var s string
	fmt.Scanln(&s)

	for a, b := 0, len(s)-1; a < b; a, b = a+1, b-1 {
		if s[a] != s[b] {
			fmt.Println("Not Palindrome")
			return
		}
	}
	fmt.Println("Palindrome")

}

// Factorial Using Recursion

// Implement a recursive function to find the factorial of a given number.

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// 11.	Sum of First N Natural Numbers

// Use a for loop to calculate the sum of the first N natural numbers.
func SumN() {
	fmt.Println("enter number ")
	var n int
	fmt.Scanln(&n)
	var sum int
	sum = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("sum of first %d natural number is %d", n, sum)
}

// Multiplication Table Generator

// Write a program that prints the multiplication table of a given number.

func multiplication() {
	fmt.Println("enter number for multiplication table")
	var n int
	fmt.Scanln(&n)
	for i := 1; i < 11; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}

// Fibonacci Series

// Print the first N terms of the Fibonacci series using a for loop.

func Fibonacci() {
	fmt.Println("enter number to get terms ")
	var n int
	fmt.Scanln(&n)
	var a int
	a = 0
	var b int
	b = 1
	for i := 0; i < n; i++ {
		fmt.Println(a)
		next := a + b
		a = b
		b = next

	}
}

// Reverse a String

// Use a for loop to reverse a given string.

func reverseString() {
	fmt.Println("enter string for reverse")
	var s string
	fmt.Scanln(&s)
	var s1 string
	s1 = ""
	for i := len(s) - 1; i >= 0; i-- {
		s1 += string(s[i])
	}
	fmt.Println(s1)
}

// Prime Number Checker

// Write a program that checks if a given number is prime using a for loop.

func PrimeNumber() {
	var num int
	fmt.Println("enter number to check prime ")
	fmt.Scanln(&num)
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			fmt.Println("not Prime")
			return
		}
	}
	fmt.Println("Prime")
}

// Even or Odd

// Write a program that checks whether an integer is even or odd.

func evenOdd() {
	fmt.Println("enter number to check even and odd")
	var n int
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("odd")
	}
}

// Largest of Three Numbers

// Accept three integers and determine the largest using if-else conditions.

func largestThree() {
	fmt.Println("largest among three numbers")
	fmt.Println(" enter first number ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println(" enter second number ")
	var num2 int
	fmt.Scanln(&num2)
	fmt.Println(" enter third number ")
	var num3 int
	fmt.Scanln(&num3)
	if num1 > num2 && num1 > num3 {
		fmt.Println("largest : ", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Println("largest : ", num2)
	} else {
		fmt.Println("largest : ", num3)
	}
}

// Leap Year Checker

// Create a function that checks if a given year is a leap year.

func leapYear() {

	fmt.Println("enter year")
	var n int
	fmt.Scanln(&n)
	if (n%4 == 0 && n%100 != 0) || n%400 == 0 {
		fmt.Println("leap year")
	} else {
		fmt.Println("not leap year")
	}
}

// Grade Calculator

// Given a score between 0 and 100, assign grades:

// 	•	90-100: A

// 	•	80-89: B

// 	•	70-79: C

// 	•	60-69: D

// 	•	Below 60: F

func grade() {
	fmt.Println("grading calculator")
	fmt.Println("enter number between 1 to 100")
	var n int
	fmt.Scanln(&n)
	if n >= 90 && n <= 100 {
		fmt.Println(" grade A")
	} else if n >= 80 && n <= 89 {
		fmt.Println(" grade B")
	} else if n >= 70 && n <= 79 {
		fmt.Println(" grade C")
	} else if n >= 60 && n <= 69 {
		fmt.Println(" grade D")
	} else if n < 60 {
		fmt.Println(" grade F")
	}
}

// Vowel or Consonant

// Accept a single character and determine if it is a vowel or a consonant.

func vowel() {
	fmt.Println("enter single character ")
	var s string
	fmt.Scanln(&s)
	char := s[0]
	if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
		fmt.Println("vowel")
	} else {
		fmt.Println("consonant")
	}
}
