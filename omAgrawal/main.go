// // package main

// // import (
// // 	"fmt"
// // )

// // func main() {

// // 	var name string
// // 	var age int
// // 	var address string
// // 	var company string
// // 	var parents string

// // 	fmt.Scan()

// // 	fmt.Println("Enter your name: ")
// // 	fmt.Scanln("%s", &name)
// // 	fmt.Println("Enter your age: ")
// // 	fmt.Scanln(&age)
// // 	fmt.Println("Enter your address: ")
// // 	fmt.Scanln(&address)
// // 	fmt.Println("Enter your company: ")
// // 	fmt.Scanln(&company)
// // 	fmt.Println("Enter your parents: ")
// // 	fmt.Scanln(&parents)

// // 	fmt.Println("Name:", name)
// // 	fmt.Println("Age:", age)
// // 	fmt.Println("Address: ", address)
// // 	fmt.Println("Company: ", company)
// // 	fmt.Println("Company:", company)
// // 	fmt.Println("Parents:", parents)

// // 	fmt.Println("Number of days he lived", age*365)
// // 	//
// // 	//fmt.Println("Enter your name: ")

// // }
// package main

// import "fmt"

// // func main() {
// // 	var name string
// // 	var age int
// // 	fmt.Print("Enter name and age: ")
// // 	fmt.Scanf("%s %d \n", &name, &age)
// // 	fmt.Scanf("%s %d", &name, &age)
// // 	fmt.Println("Name:", name, "Age:", age)
// // 	//

// // }

// func main() {
// 	var a int
// 	var b int
// 	var opr string
// 	fmt.Println("enter two number")
// 	fmt.Scanln(&a)
// 	fmt.Scanln(&b)
// 	fmt.Scanln(&opr)
// 	fmt.Println("Enter a operation + - * /")
// 	if opr == "+" {
// 		fmt.Println(a + b)
// 	} else if opr == "*" {
// 		fmt.Println(a * b)
// 	} else if opr == "-" {
// 		fmt.Println(a - b)
// 	} else if opr == "/" {
// 		fmt.Println(a / b)
// 	}

// }

package main

import (
	"fmt"
	"math"
	"nexturn_Go/services"
)

// swap two number
func swap(num1 int, num2 int) (int, int) {
	return num2, num1

}

//areaa of ciecle

func area(radius float64) float64 {
	return math.Pi * radius * radius
}

//celcius to farrenhite

func convert(num int) int {
	return (num * 9 / 5) + 32
}

// string concatenation
func concatenation(s1 string, s2 string) string {
	return s1 + s2
}

// sum 6
func sum(n int, m int) int { return n + m }

// 7 variadic function
func variadic(n ...int) int {
	summ := 0
	for i := 0; i < len(n); i++ {
		summ += i
	}
	return summ
}

// 9.	Palindrome Checker (String)
// Write a function that checks whether a given string is a palindrome.
// func palinfrome(str ...string) {

// 	var s1 string

// 	for i := 0; i < len(str); i++ {
// 		s1+=str[i] // takes a rune

// 	}
// 	if(s1==str) {
// 		fmt.Println("palindrome");}
// }

// 10.	Factorial Using Recursion
// Implement a recursive function to find the factorial of a given number.

func factorial(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)

}

// 11.	Sum of First N Natural Numbers // Use a for loop to calculate the sum of the first N natural numbers.
func sumFirstN(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

// 12.	Multiplication Table Generator
// Write a program that prints the multiplication table of a given number.
func multiplicationTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}

// 13.	Fibonacci Series // Print the first N terms of the Fibonacci series using a for loop.

func fibonnanci(num int) {
	n1 := 0
	n2 := 1

	fmt.Printf("%d ,%d ", n1, n2)
	for i := 0; i <= num-2; i++ {
		n3 := n2 + n1
		n1 = n2
		n2 = n3
		fmt.Printf("%d ", n3)

	}
}

// 14.	Reverse a String// Use a for loop to reverse a given string.
func reverse(s string) {
	var s11 string
	for i := len(s) - 1; i >= 0; i-- {
		s11 += string(s[i])

	}
	fmt.Println(s11)
}

// 15.	Prime Number Checker
func primeChecker(n int) {
	p := true

	for i := 2; i < n; i++ {

		if n%i == 0 {
			p = false
		}

	}

	if p {
		fmt.Println("n is a prime number")
	} else {
		fmt.Println("n is not a prime number")
	}
}

// Write a program that checks if a given number is prime using a for loop.

// 	16.	Even or Odd

func evenOdd(n int) {
	if n%2 == 0 {
		fmt.Println(" even")
	} else {
		fmt.Println(" odd")

	}
}

// Write a program that checks whether an integer is even or odd.

// 17.	Largest of Three Numbers
func larger3(a int, b int, c int) {
	if a > b && a > c {
		fmt.Println("a")

	} else if a < b && b > c {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}
}

// Accept three integers and determine the largest using if-else conditions.

// 	18.	Leap Year Checker

// Create a function that checks if a given year is a leap year.
func leapYear(n int) {
	if n%400 == 0 {
		fmt.Println("leap")

	} else if n%100 == 0 {
		fmt.Println("not a leap")
	} else if n%4 == 0 {
		fmt.Println("leap ")
	}
}

// 	19.	Grade Calculator

// Given a score between 0 and 100, assign grades:

// 	•	90-100: A

// 	•	80-89: B

// 	•	70-79: C

// 	•	60-69: D

// 	•	Below 60: F

func graded(n int) {
	if n >= 90 {
		fmt.Println("A")
	} else if n >= 80 {
		fmt.Println("B")
	} else if n > 70 {
		fmt.Println("C")
	} else if n > 60 {
		fmt.Println("D")

	}
}

// 	20.	Vowel or Consonant

// Accept a single character and determine if it is a vowel or a consonant.

func main() {

	a := 3
	b := 2

	fmt.Println(services.GetService())
	fmt.Println(services.GetService2())

	a, b = swap(a, b)
	fmt.Println("Swapped a,b is ")
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("area of circle with radius 30")
	area := area(30.0)
	fmt.Printf("area of circle %f \n", area)

	fmt.Println("tenpreture from celsius to farren")
	f := convert(70)
	fmt.Printf("cenlsius to farrenhite %d \n", f)

	str := concatenation("string 1 +  ", "string 2")

	fmt.Printf("%s \n", str)

	//data type of any variable

	fmt.Printf("data type of variable %T \n", f)

	sum := sum(8, 90)
	fmt.Printf("sum of two number 5 , 90  %d \n", sum)

	sumvariadic := variadic(1, 2, 3, 4, 5, 6)
	fmt.Println("sum of 1,2,3,4,,5,6 is  : ", sumvariadic)
	//10
	fmt.Printf("palindrome of a number ")
	fmt.Println(factorial(6))
	//11
	fmt.Printf("sum of first n  number ")
	fmt.Println(sumFirstN(6))
	//12
	fmt.Printf("multiplication table of n  ")
	multiplicationTable(6)

	//13
	fibonnanci(12)

	//14
	reverse("hiiHello")

}
