package main

import (
	"fmt"
	"math"
)

func main() {
	//1. Swap to integers
	num1 := 2
	num2 := 4
	num1, num2 = SwapTwoNumbers(num1, num2)
	fmt.Printf("Swapped numbers are %d, %d", num1, num2)

	//2. Area of Circle
	var r float64 = 3
	fmt.Printf("area of circle of radius %f is %f.", r, CircleArea(r))

	//3. Celsius to Fahrenheit
	cel := 0.0
	fmt.Println(CelsiusToFahrenheit(cel))

	//4. string concatenation
	var s1, s2 string
	fmt.Println("Enter first string")
	fmt.Scanln((&s1))
	fmt.Println("Enter second string")
	fmt.Scanln((&s2))
	fmt.Println(Concatenation(s1, s2))

	//5 check data type
	c := "adsfghj"
	fmt.Printf("Data type is %T", c)

	//6 sum of two numbers
	fmt.Println("Sum is: ", Sum(15, 25))

	//7 variadic sum
	fmt.Println("Variadic Sum is: ", VariadicSum(1, 3, 5, 7, 2, 9, 10))

	//8 max in variadic func
	fmt.Println("Max num is: ", MaxNum(2, 5, 7, 2, 54, 6, 99))

	//9 palindrome checker
	fmt.Println("Palindrome: ", IsPalindrome("madam"))

	// 10 factorial of a number
	fmt.Println("Factorial of 5 is:", Factorial(5))

	//11 sum of first n natural numbers
	fmt.Println("Sum from 1 upto 100: ", SumOfNaturalNumbers(100))

	// 12 Mjltiplication table generator
	Multiplication(12)

	// 13 Fibonacci series
	Fibonacci(10)

	//14 Reverse String
	fmt.Println("Reverse of vicky is:", Reverse("Vicky"))

	//15 Prime no checker
	fmt.Println("Is Prime 23: ", IsPrime(20))

	//16 Even or Odd
	fmt.Println("10 is:", EvenOrOdd(10))
	fmt.Println("11 is:", EvenOrOdd(11))

	// 17 Largest of 3 numbers
	fmt.Println("Largest of three: ", Largest(14, 10, 8))

	// 18 Lear year
	year := 1900
	fmt.Println("Is", year, "leap year:", IsLeapYear(year))

	// 19 Grade calc
	marks := 80
	fmt.Printf("Grade for marks %d : %s", marks, GradeCalc(marks))

	// 20 vowel or const
	cc := "v"
	fmt.Println("Is", cc, "vowel:", IsVowel(cc))
}

func SwapTwoNumbers(num1 int, num2 int) (int, int) {
	return num2, num1
}

func CircleArea(radius float64) float64 {
	const pi = math.Pi
	return pi * radius * radius
}
func CelsiusToFahrenheit(celsius float64) float64 {
	fahrenheit := (celsius*9)/5 + 32
	return fahrenheit
}
func Concatenation(s1 string, s2 string) string {
	return s1 + s2
}
func Sum(num1 int, num2 int) int {
	return num1 + num2
}
func VariadicSum(n ...int) int {
	result := 0
	for i := 0; i < len(n); i++ {
		result += n[i]
	}
	return result
}
func MaxNum(n ...int) int {
	max := math.MinInt
	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
	}
	return max
}
func IsPalindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
func Factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * Factorial(num-1)
}
func SumOfNaturalNumbers(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}
func Multiplication(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "*", i, "=", num*i)
	}
}
func Fibonacci(num int) {
	a := 0
	b := 1
	for i := 0; i < num; i++ {
		fmt.Println(a)
		c := a
		a = b
		b = b + c
	}
}
func Reverse(s string) string {
	rev := ""
	for i := len(s) - 1; i >= 0; i-- {
		rev = rev + string(s[i])
	}
	return rev
}
func IsPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func EvenOrOdd(num int) string {
	if num&1 == 1 {
		return "Odd"
	}
	return "Even"
}
func Largest(a int, b int, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}
}
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
func GradeCalc(marks int) string {
	if marks >= 90 {
		return "A"
	} else if marks >= 80 {
		return "B"
	} else if marks >= 70 {
		return "C"
	} else if marks >= 60 {
		return "D"
	} else {
		return "F"
	}
}
func IsVowel(c string) bool {
	if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" || c == "A" || c == "E" || c == "I" || c == "O" || c == "U" {
		return true
	}
	return false
}
