package main

import (
	"fmt"
	"strconv"
)

func main() {
	stringToInt()
		temperature()
	average()
	boolToint()
	flaotToInt()
	hexaToDeci()
	sliceConvert()
	flaotError()
	multpleconvert()
	asciic_to_int()
}

// Convert String Inputs to Integers and Calculate Sum:

// Write a program that takes five string inputs from the user, converts them to integers, and prints the sum.

func stringToInt() {
	fmt.Println("conver string to int ")
	var sum int
	var a string
	for i := 0; i < 5; i++ {
		fmt.Printf("enter %d number ", i+1)
		fmt.Scanln(&a)
		b, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("error")
		}
		sum += b

	}
	fmt.Println("sum : ", sum)
}

// Fahrenheit to Celsius Converter:

// Create a function that converts Fahrenheit to Celsius. Take temperature input from the user and display the result with two decimal places.

func temperature() {
	fmt.Println("convert fahrenheit to celsius")
	fmt.Println("enter input ")
	var a int
	fmt.Scanln(&a)
	b := float32(a)
	c := (5 * (b - 31)) / 7
	fmt.Println("celsius  : ", c)
}

// . Integer to Float Conversion and Average Calculation:

// Take three integer inputs from the user, convert them to floats, and calculate the average using a function.

func average() {
	fmt.Println("average of three")
	sum := 0
	var a int
	for i := 0; i < 3; i++ {
		fmt.Printf("enter %d number ", i+1)
		fmt.Scanln(&a)
		sum += a
	}
	b := float32(sum / 3)
	fmt.Printf("average : %.2f", b)
}

// Boolean to Integer Conversion:

// Write a function that converts a boolean value to an integer (true -> 1, false -> 0). Take user input as "true" or "false" and convert it.

func boolToint() {
	fmt.Println("bool to int ")
	fmt.Println("enter bool value ")
	var a string
	fmt.Scanln(&a)
	if a == "false" || a == "False" {
		fmt.Println("0")
	} else if a == "True" || a == "true" {
		fmt.Println("1")
	}
}

// Convert Float to Integer and Check Even/Odd:

// Take a float input from the user, convert it to an integer, and determine whether the resulting integer is even or odd.

func flaotToInt() {
	fmt.Println("float to int ")
	fmt.Println("enter value ")
	var a float32
	fmt.Scanln(&a)
	b := int(a)
	if b%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

}

// Hexadecimal to Decimal Converter:

// Ask the user to enter a hexadecimal string (e.g., "1A"), convert it to an integer, and display the decimal equivalent.

func hexaToDeci() {
	fmt.Println("Hexadecimal to decimal convert")
	fmt.Println("enter hexadecimal value ")
	var a string
	fmt.Scanln(&a)
	b, err := strconv.ParseInt(a, 16, 32)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("decimal : ", b)
}

// Iterate Over a Slice of Strings and Convert to Integers:

// Given a slice of string numbers ([]string{"10", "20", "30", "40"}), write a program to convert them into integers and calculate their total sum.

func sliceConvert(){
	fmt.Println("convert slice to int ")
  a:= []string{"10","20", "30", "40 "}
	sum:=0
	for i:=0; i<len(a)-1;i++{
		b,_:=strconv.Atoi(a[i])
		sum+=b
	}
	fmt.Println("sum : ",sum)
}

// String to Float Conversion with Error Handling:

// Take a string input representing a floating-point number, convert it to a float, and print the square of the number. Handle any conversion errors.

func flaotError(){
	fmt.Println("string to float and print square")
	fmt.Println("enter value ")
	var inp string 
	fmt.Scanln(&inp)
	F,err := strconv.ParseFloat(inp,32)
	if err!=nil {
		fmt.Println("error")
	}else {fmt.Println("sqaure : ", F*F)}
}

// Read Multiple Values and Convert Types:

// Read an integer, a float, and a string from the user. Convert the integer to a float and the float to an integer, then display the converted values.


func multpleconvert(){
	fmt.Println("multiple convert")
	fmt.Println("enter int value ")
	var a int
	fmt.Scanln(&a)
	b := float32(a)
	fmt.Println(b)
	fmt.Println("enter float value ")
	var c float32
	fmt.Scanln(&c)
	d:= int(c)
	fmt.Println(d)
	
}

// Character to ASCII Value Converter:

// Prompt the user to enter a single character, convert it to its ASCII integer value, and print the result.

func asciic_to_int(){
	fmt.Println("enter single character for ascci value ")
	var a string
	fmt.Scanln(&a)
	b:=int(a[0])
	fmt.Println(b)
}