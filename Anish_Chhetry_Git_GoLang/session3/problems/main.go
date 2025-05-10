package main

import (
	"fmt"
	"strconv"
)

var sum int

func main() {
	// 	1. Convert String Inputs to Integers and Calculate Sum:
	// Write a program that takes five string inputs from the user, converts them to integers, and prints the sum.
	fmt.Println("1. Convert String Inputs to Integers and Calculate Sum:")
	var str1, str2, str3, str4, str5 string
	fmt.Println("Enter 5 numbers: ")
	fmt.Scanf("%s %s %s%s %s\n", &str1, &str2, &str3, &str4, &str5)
	stringtoint(str1)
	stringtoint(str2)
	stringtoint(str3)
	stringtoint(str4)
	stringtoint(str5)
	fmt.Println("SUM = ", sum)

	// 2. Fahrenheit to Celsius Converter:
	// Create a function that converts Fahrenheit to Celsius. Take temperature input from the user and display the result with two decimal places.
	fmt.Println("2. Fahrenheit to Celsius Converter:")
	var fah int
	fmt.Println("Enter temperature in Fahrenheit: ")
	fmt.Scanf("%d\n", &fah)
	fmt.Println(FahrenheitToCelsius(fah))

	// 3. Integer to Float Conversion and Average Calculation:
	// Take three integer inputs from the user, convert them to floats, and calculate the average using a function.
	fmt.Println("3. Integer to Float Conversion and Average Calculation:")
	var num1, num2, num3 int
	fmt.Println("Enter 3 numbers: ")
	fmt.Scanf("%d %d %d\n", &num1, &num2, &num3)
	intToFloatAvg(num1, num2, num3)

	// 4. Boolean to Integer Conversion:
	// Write a function that converts a boolean value to an integer (true -> 1, false -> 0). Take user input as "true" or "false" and convert it.
	fmt.Println("4. Boolean to Integer Conversion:")
	var b1 bool
	fmt.Println("Enter true or false: ")
	fmt.Scanf("%t\n", &b1)
	fmt.Printf("%d\n", boolToInt(b1))

	// 5. Convert Float to Integer and Check Even/Odd:
	// Take a float input from the user, convert it to an integer, and determine whether the resulting integer is even or odd.
	fmt.Println("5. Convert Float to Integer and Check Even/Odd:")
	var a float32
	fmt.Println("Enter a float: ")
	fmt.Scanf("%f\n", &a)

	b := int(a)
	if b%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 6. Hexadecimal to Decimal Converter:
	// Ask the user to enter a hexadecimal string (e.g., "1A"), convert it to an integer, and display the decimal equivalent.
	fmt.Println("6. Hexadecimal to Decimal Converter:")
	var hex string
	fmt.Print("Hex: ")
	fmt.Scan(&hex)
	if dec, err := strconv.ParseInt(hex, 16, 0); err == nil {
		fmt.Printf("Decimal: %d\n", dec)
	} else {
		fmt.Println("Invalid hex")
	}

	// 7. Iterate Over a Slice of Strings and Convert to Integers:
	// Given a slice of string numbers ([]string{"10", "20", "30", "40"}), write a program to convert them into integers and calculate their total sum.
	fmt.Println("7. Iterate Over a Slice of Strings and Convert to Integers:")
	slice := []string{"10", "20", "30", "40"}
	stringtointSlice(slice)

	// 8. String to Float Conversion with Error Handling:
	// Take a string input representing a floating-point number, convert it to a float, and print the square of the number. Handle any conversion errors.
	fmt.Println("8. String to Float Conversion with Error Handling:")
	var s1 string
	fmt.Println("Enter a number")
	fmt.Scanf("%s\n", &s1)
	stringtofloat(s1)

	// 9. Read Multiple Values and Convert Types:
	// Read an integer, a float, and a string from the user. Convert the integer to a float and the float to an integer, then display the converted values.
	fmt.Println("9. Read Multiple Values and Convert Types:")
	var v1 int
	var v2 float32
	fmt.Println("Enter an integer")
	fmt.Scanf("%d\n", &v1)
	fmt.Println("Enter a float")
	fmt.Scanf("%f\n", &v2)

	val1 := float32(v1)
	val2 := int(v2)
	fmt.Printf("%f\n", val1)
	fmt.Printf("%d\n", val2)

	// 10. Character to ASCII Value Converter:
	// Prompt the user to enter a single character, convert it to its ASCII integer value, and print the result.
	fmt.Println("10. Character to ASCII Value Converter:")
	var c rune
	fmt.Print("Enter a char: ")
	fmt.Scanf("%c\n", &c)
	fmt.Printf("ASCII of '%c' is %d\n", c, c)
}

func stringtoint(str string) {

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Conversion not possible")
	}
	sum += num
}

func FahrenheitToCelsius(fah int) float32 {
	f := float32(fah)
	c := (f - 32) * (5.0 / 9.0)
	return c
}
func intToFloatAvg(num1 int, num2 int, num3 int) {
	average := (float32(num1) + float32(num2) + float32(num3)) / 3
	fmt.Println("Average is", average)
}
func boolToInt(b1 bool) int {
	if b1 {
		return 1
	} else {
		return 0
	}
}
func stringtointSlice(str []string) {
	sumslice := 0
	for i := 0; i < len(str); i++ {
		num, err := strconv.Atoi(str[i])
		if err != nil {
			fmt.Println("Conversion not possible")
		}
		sumslice += num
	}
	fmt.Println(sumslice)
}
func stringtofloat(s1 string) {
	num, err := strconv.ParseFloat(s1, 32)
	if err != nil {
		fmt.Println("Conversion not possible")
	}
	fmt.Println(num * num)
}
