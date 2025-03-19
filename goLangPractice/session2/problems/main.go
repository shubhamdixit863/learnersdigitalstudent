package main

import (
	"fmt"
	"math"
	"strings"
)

func interchange(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func area(r float64) float64 {
	return math.Pi * r * r
}

func tempConv(c float64) float64 {
	f := c*(9/5) + 32
	return f
}

func strConcat(a, b string) string {
	return a + b
}

func varType(a interface{}) {
	fmt.Printf("Type: %T\n", a)
}

func sum(a, b int) int {
	return a + b
}

func sum2(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func maxVal(a ...int) int {
	m := math.MinInt
	fmt.Println(m)
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

func Palindrome(s string) bool {
	s = strings.ToLower(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumOfN(N int) int {
	sum := 0
	for i := 1; i <= N; i++ {
		sum += i
	}
	return sum
}

func table(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n, " * ", i, " = ", n*i)

	}
}

func Fibonacci(n int) {
	temp1 := 1
	temp2 := 2

	for i := 1; i <= n; i++ {
		if i == 1 || i == 2 {
			fmt.Print(i, " ")
		} else {
			temp := temp2
			temp2 = temp2 + temp1
			temp1 = temp
			fmt.Print(temp2, " ")
		}
	}
}

func strReverse(s string) string {
	var temp string
	for i := len(s) - 1; i >= 0; i-- {
		temp += string(s[i])
	}
	return temp
}

func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func oddEven(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func largestOfThree(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c && b > a {
		return b
	} else {
		return c
	}
}

func leapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

func gradeCalc(n int) string {
	if n >= 90 && n <= 100 {
		return "A"
	} else if n >= 80 && n <= 89 {
		return "B"
	} else if n >= 70 && n <= 79 {
		return "c"
	} else if n >= 60 && n <= 69 {
		return "D"
	} else {
		return "F"
	}
}

func vowelConsonent(c string) string {
	c = strings.ToLower(c)
	if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
		return "Vowel"
	}
	return "Consonent"
}

func main() {
	// Que: 1
	//a := 5
	//b := 7
	//fmt.Println(interchange(a, b))

	// Que: 2
	//fmt.Println(area(3))

	// Que: 3
	//fmt.Println(tempConv(0))

	// Que: 4
	//fmt.Println(strConcat("Sagar ", "Sinha"))

	// Que: 5
	//varType(100)

	// Que: 6
	//fmt.Println(sum(1, 2))

	// Que: 7
	//fmt.Println(sum2(5, 7, 9))

	// Que:8
	//fmt.Println(maxVal(9, 5, 501, 209, 1, 2, 3))

	// Que: 9
	//fmt.Println(Palindrome("123321"))

	// Que: 10
	//fmt.Println(factorial(5))

	// Que: 11
	//fmt.Println(sumOfN(6))

	// Que: 12
	//table(9)

	// Que: 13
	//Fibonacci(10)

	// Que: 14
	//fmt.Println(strReverse("hello world"))

	// Que: 15
	//fmt.Println(prime(21))

	// Que: 16
	//fmt.Println(oddEven(9))

	// Que: 17
	//fmt.Println(largestOfThree(9, 7, 11))

	// Que: 18
	//fmt.Println(leapYear(2000))

	// Que: 19
	//fmt.Println(gradeCalc(70))

	// Que: 20
	fmt.Println(vowelConsonent("z"))
}
