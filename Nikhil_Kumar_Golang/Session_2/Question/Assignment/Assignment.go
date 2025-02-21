/*
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
		11.	Sum of First N Natural Numbers
	Use a for loop to calculate the sum of the first N natural numbers.
		12.	Multiplication Table Generator
	Write a program that prints the multiplication table of a given number.
		13.	Fibonacci Series
	Print the first N terms of the Fibonacci series using a for loop.
		14.	Reverse a String
	Use a for loop to reverse a given string.
		15.	Prime Number Checker
	Write a program that checks if a given number is prime using a for loop
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


*/

package main

import (
	"fmt"
	"math"
)

func swap_two_integers() {

	var a, b int
	fmt.Print("Enter two numbers \n")
	fmt.Print("A : ")
	fmt.Scan(&a)
	fmt.Print("B : ")
	fmt.Scan(&b)
	a=a+b
	b=a-b
	a=a-b
	fmt.Printf("A : %d\nB : %d\n", a, b)
}
func area_of_circle()  {
	var radius float32
	fmt.Print("Enter radius of circle : ")
	fmt.Scan(&radius)
	area := math.Pi * radius * radius
	fmt.Printf("Area of circle : %.2f\n", area)
}
func celsius_to_fahrenheit()  {
	var celsius float32
	fmt.Print("Enter celsius : ")
	fmt.Scan(&celsius)
	fahrenheit := (celsius *(9/5)) + 32
	fmt.Printf("Temperature in fahrenheit : %.2f\n", fahrenheit)
}
func String_Concatenation()  {
	var str1, str2 string
	fmt.Print("Enter two strings\n")
	fmt.Print("Strings 1: ")
	fmt.Scan(&str1)
	fmt.Print("Strings 2: ")
	fmt.Scan(&str2)
	str3 := str1 + str2
	fmt.Printf("concatenated string : %s\n", str3)
}
func check_data_type(){
	var a int
	var b float32
	var c float64
	var d bool
	var e string
	var f error
	var g complex64

	
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)
	fmt.Printf("%T\n",d)
	fmt.Printf("%T\n",e)
	fmt.Printf("%T\n",f)
	fmt.Printf("%T\n",g)
}
func sum_of_two_num(a int ,b int) int{
	return a+b
}

func Variadic_Sum_Function(num ...int) int{
	sum := 0
	for i:=0; i<len(num); i++{
		sum += num[i]
	}
	return sum
}
func palindrome_check(str string) bool{
	var rev string
	for i:=len(str)-1; i>=0; i--{
		rev += string(str[i])
	}
	if str == rev{
		return true
	}
	return false
}
func factorial(i int)int{
	if i == 0{
		return 1
	}
	return i * factorial(i-1)
}
func sum_of_n_number(k int )int {
	var sum int
	for i:=1; i<=k; i++{
		sum += i
	}
	return sum
}
func multiplication_table_generator(num int)  {
	for i:=1; i<=10; i++{
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
}

func fibonacci(n int) {
	var f1 int 
	var f2 int
	var f3 int
	f1 = 1
	f2 = 1
	for i:=1; i<=n; i++{
		fmt.Println(f1)
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
}
func Reverse_a_String(str string) string{
	var rev string
	for i:=len(str)-1; i>=0; i--{
		rev += string(str[i])
	}
	return rev
}

func Prime_Number_Checker(num int) bool{
	for i:=2; i<num; i++{
		if num % i == 0{
			return false
		}
	}
	return true
}
func Odd_Even(num int) bool{
	if num % 2 == 0{
		return true
	}
	return false
}
func Largest_of_Three_Numbers(num1 int, num2 int, num3 int) int{
	if num1 > num2 && num1 > num3{
		return num1
	}else if num2 > num1 && num2 > num3{
		return num2
	}else{
		return num3
	}
}
func Leap_Year_Checker(num int) bool{
	if num % 4 == 0 && num % 100 != 0 || num % 400 == 0{
		return true
	}
	return false
}
func Grade_Calculator(num int) string{
	if num >= 90 && num <= 100{
		return "A"
	}else if num >= 80 && num <= 89{
		return "B"
	}else if num >= 70 && num <= 79{
		return "C"
	}else if num >= 60 && num <= 69{
		return "D"
	}else{
		return "F"
	}	
}
func Vowel_or_Consonant(char string) string{
	if char == "a" || char == "e" || char == "i" || char == "o" || char == "u" || char == "A" || char == "E" || char == "I" || char == "O" || char == "U"{
		return "Vowel"
	}else{
		return "Consonant"
	}
}

func main() {
	swap_two_integers()
	area_of_circle()
	celsius_to_fahrenheit()
	String_Concatenation()
	check_data_type()
	fmt.Println("sum of two numbers : ", sum_of_two_num(1,2))
	fmt.Println("variadic sum of numbers : ",Variadic_Sum_Function(1,2,3,4,5,6,7,8,9,10))
	fmt.Println("palindrome check : ", palindrome_check("aabbaa"))
	fmt.Println("factorial : ", factorial(5))
	fmt.Println("sum of n number : ", sum_of_n_number(5))
	fmt.Println("multiplication table generator : ")
	multiplication_table_generator(5)
	fmt.Println("fibonacci series : ")
	fibonacci(10)
	fmt.Println("reverse a string : ", Reverse_a_String("hello"))
	fmt.Println("prime number checker : ", Prime_Number_Checker(8))
	fmt.Println("odd even checker : ", Odd_Even(8))
	fmt.Println("largest of three numbers : ", Largest_of_Three_Numbers(1,2,3))
	fmt.Println("leap year checker : ", Leap_Year_Checker(2000))
	fmt.Println("grade calculator : ", Grade_Calculator(80))
	fmt.Println("vowel or consonant : ", Vowel_or_Consonant("a"))
}
