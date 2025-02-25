/*1. Convert String Inputs to Integers and Calculate Sum:
Write a program that takes five string inputs from the user, converts them to integers, and prints the sum.

2. Fahrenheit to Celsius Converter:
Create a function that converts Fahrenheit to Celsius. Take temperature input from the user and display the result with two decimal places.

3. Integer to Float Conversion and Average Calculation:
Take three integer inputs from the user, convert them to floats, and calculate the average using a function.

4. Boolean to Integer Conversion:
Write a function that converts a boolean value to an integer (true -> 1, false -> 0). Take user input as "true" or "false" and convert it.

5. Convert Float to Integer and Check Even/Odd:
Take a float input from the user, convert it to an integer, and determine whether the resulting integer is even or odd.

6. Hexadecimal to Decimal Converter:
Ask the user to enter a hexadecimal string (e.g., "1A"), convert it to an integer, and display the decimal equivalent.

7. Iterate Over a Slice of Strings and Convert to Integers:
Given a slice of string numbers ([]string{"10", "20", "30", "40"}), write a program to convert them into integers and calculate their total sum.

8. String to Float Conversion with Error Handling:
Take a string input representing a floating-point number, convert it to a float, and print the square of the number. Handle any conversion errors.

9. Read Multiple Values and Convert Types:
Read an integer, a float, and a string from the user. Convert the integer to a float and the float to an integer, then display the converted values.

10. Character to ASCII Value Converter:
Prompt the user to enter a single character, convert it to its ASCII integer value, and print the result.

*/

package main

import (
	"fmt"
	"strconv"
)

func calc_sum()  {
	fmt.Println("1. Convert String Inputs to Integers and Calculate Sum:")
	fmt.Println("enter 5 numbers")
	var inp string
	var sum int
	for i := 0; i < 5; i++ {
		fmt.Println("enter ", i+1, " numbers :")
		fmt.Scan(&inp)
		flag,err := strconv.Atoi(inp)
		if err != nil {
			fmt.Println(err)
		}else {
			sum+=flag
		}
	}
	fmt.Println("Sum:", sum)
}


func converter()  {
	fmt.Println("2. Fahrenheit to Celsius Converter:")
	var fahrenheit int
	fmt.Println("Enter temperature in Fahrenheit:")
	fmt.Scan(&fahrenheit)
	celsius := float64((fahrenheit - 32) * 5 / 9)
	fmt.Printf("Temperature in Celsius: %.2f\n", celsius)

}
func average()  {
	fmt.Println("3. Integer to Float Conversion and Average Calculation:")
	var num [3]int
	var sum int
	for i := 0; i < 3; i++ {
		fmt.Println("Enter", i+1, "number:")
		fmt.Scan(&num[i])
		sum += num[i]
	}
	avg := float32(sum / 3)
	fmt.Printf("Average: %.2f\n", avg)

}
func boot_to_int()  {
	fmt.Println("4. Boolean to Integer Conversion:")
	var inp string
	fmt.Println("Enter a boolean value (true or false):")
	fmt.Scan(&inp)
	var result int
	if inp == "true" {
		result = 1
	} else {
		result = 0
	}
	fmt.Println("Integer value:", result)
}
func even_odd()  {
	fmt.Println("5. Convert Float to Integer and Check Even/Odd:")
	var inp float64
	fmt.Println("Enter a floating-point number:")
	fmt.Scan(&inp)
	var result int
	result = int(inp)
	if result%2 == 0 {
		fmt.Println("Number is even.")
	} else {
		fmt.Println("Number is odd.")
	}
}

func hex_to_dec()  {
	fmt.Println("6. Hexadecimal to Decimal Converter:")
	var inp string
	fmt.Println("Enter a hexadecimal number:")
	fmt.Scan(&inp)
	dec, err := strconv.ParseInt(inp, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Decimal value:", dec)
}
func iterate_over_slice()  {
	fmt.Println("7. Iterate Over a Slice of Strings and Convert to Integers:")
	str := []string{"10", "20", "30", "40"}
	fmt.Println("enter string :")
	fmt.Scan(&str)
	var sum int
	for i := 0; i < len(str); i++ {
		num, err := strconv.Atoi(str[i])
		if err != nil {
			fmt.Println(err)
		}
		sum += num
	}

	fmt.Println("Sum:", sum)
}
func string_to_float()  {
	fmt.Println("8. String to Float Conversion with Error Handling:")
	var inp string
	fmt.Println("Enter a string:")
	fmt.Scan(&inp)
	flt, err := strconv.ParseFloat(inp, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Float squarred value: %.2f", flt*flt)
}
func read_multiple_values()  {
	fmt.Println("9. Read Multiple Values and Convert Types:")
	var inp string
	var num int
	var flt float64
	fmt.Println("Enter a string:")
	fmt.Scan(&inp)
	intVal, err := strconv.Atoi(inp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("strinf converted into integer:", intVal)
	}
	fmt.Println("Enter an integer:")
	fmt.Scan(&num)
	fmt.Printf("integer converted into float: %.2f", float64(num))
	fmt.Println("Enter a floating-point number:")
	fmt.Scan(&flt)	
	fmt.Println("float converted into int :", int(flt))
}
func convert_to_ascii()  {
	fmt.Println("10. Character to ASCII Value Converter:")
	var inp string
	fmt.Println("Enter a character:")
	fmt.Scan(&inp)
	ascii := int(inp[0])
	fmt.Println("ASCII value:", ascii)
}
func main() {
	calc_sum()
	converter()
	average()
	boot_to_int()
	even_odd()
	hex_to_dec()
	iterate_over_slice()
	string_to_float()
	read_multiple_values()
	convert_to_ascii()
}