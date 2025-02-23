package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//1
	StringToInt()

	//2
	FahrenheitToCelsius()

	//3
	IntToFloat()

	//4
	BoolToInt()

	//5
	FloatToInt()

	//6
	var hex string
	fmt.Println("Enter hexadecimal number:")
	fmt.Scanln(&hex)
	fmt.Println(HexadecimalToDecimal(hex))

	//7
	s := []string{"1", "2", "3", "4"}
	StringSliceToInt(s)

	//8
	StringToFloat()

	//9
	MultipleInput()

	//10
	CharToASCII()
}

// 1. Convert String Inputs to Integers and Calculate Sum:
// Write a program that takes five string inputs from the user, converts them to integers, and prints the sum.
func StringToInt() {
	sum := 0
	for i := 1; i <= 5; i++ {
		var s string
		fmt.Println("Please enter string", i)
		fmt.Scanln(&s)
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += num
	}
	fmt.Println("Sum is", sum)
}

// 2. Fahrenheit to Celsius Converter:
// Create a function that converts Fahrenheit to Celsius. Take temperature input from the user and display the result with two decimal places.
func FahrenheitToCelsius() {
	var fahrenheit float64
	fmt.Println("Enter temperature in Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f Fahrenheit is equal to %.2f Celsius\n", fahrenheit, celsius)
}

// 3  Integer to Float Conversion and Average Calculation:
// Take three integer inputs from the user, convert them to floats, and calculate the average using a function.
func IntToFloat() {
	sum := 0.0
	for i := 1; i <= 3; i++ {
		var num int
		fmt.Println("Please enter integer", i)
		fmt.Scanln(&num)
		f := float64(num)
		sum += f
	}
	fmt.Println("Avg is", sum/3)
}

//4. Boolean to Integer Conversion:
//Write a function that converts a boolean value to an integer (true -> 1, false -> 0). Take user input as "true" or "false" and convert it.

func BoolToInt() {
	var b bool
	fmt.Println("Enter bool value true or false: ")
	fmt.Scanln(&b)
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	fmt.Println("Bool to int is:", i)
}

// 5 Convert Float to Integer and Check Even/Odd:
// Take a float input from the user, convert it to an integer, and determine whether the resulting integer is even or odd.
func FloatToInt() {
	var f float64
	fmt.Println("Enter float number")
	fmt.Scanln(&f)
	i := int(f)
	if i&1 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

// 6 Hexadecimal to Decimal Converter:
// Ask the user to enter a hexadecimal string (e.g., "1A"), convert it to an integer, and display the decimal equivalent.
func HexadecimalToDecimal(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		c := s[i]
		num := 0
		if c == 'A' || c == 'B' || c == 'C' || c == 'D' || c == 'E' || c == 'F' {
			num = 10 + int(c-'A')
		} else {
			num, _ = strconv.Atoi(string(c))
		}
		ans += num * int(math.Pow(16, float64(len(s)-i-1)))
	}
	return ans
}

// 7. Iterate Over a Slice of Strings and Convert to Integers:
//Given a slice of string numbers ([]string{"10", "20", "30", "40"}), write a program to convert them into integers and calculate their total sum.

func StringSliceToInt(str []string) {
	sum := 0
	for i := 0; i < len(str); i++ {
		num, err := strconv.Atoi(str[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += num
	}
	fmt.Println("Sum is", sum)
}

// 8 String to Float Conversion with Error Handling:
// Take a string input representing a floating-point number, convert it to a float, and print the square of the number. Handle any conversion errors.
func StringToFloat() {
	var s string
	fmt.Println("Please enter string")
	fmt.Scanln(&s)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("square is", f*f)
}

// 9. Read Multiple Values and Convert Types:
// Read an integer, a float, and a string from the user. Convert the integer to a float and the float to an integer, then display the converted values.
func MultipleInput() {
	var num int
	fmt.Println("Please enter integer")
	fmt.Scanln(&num)
	f := float64(num)
	fmt.Printf("Int to Float is %f\n", f)

	var num2 float64
	fmt.Println("Please enter float")
	fmt.Scanln(&num2)
	i := int(num2)
	fmt.Println("Float to Int is", i)

}

// 10. Character to ASCII Value Converter:
// Prompt the user to enter a single character, convert it to its ASCII integer value, and print the result.
func CharToASCII() {
	c := 'a'
	fmt.Println("Enter a single character")
	fmt.Scanf("%c", &c)
	i := int(c)
	fmt.Println("ascii to int:", i)
}
func CharToASCII2() {
	var s string
	fmt.Println("Enter a single character")
	fmt.Scanln(&s)
	if len(s) > 1 {
		fmt.Println("More than 1 character entered")
	} else {
		c := s[0]
		i := int(c)
		fmt.Println("ascii to int:", i)
	}
}
