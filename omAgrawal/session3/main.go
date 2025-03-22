package main

import (
	"fmt"
	"strconv"

	"math"
)

// 1. Convert String Inputs to Integers and Calculate Sum: // Write a program that takes five string inputs from the user, converts them to integers, and prints the sum.

// 2. Fahrenheit to Celsius Converter:
// Create a function that converts Fahrenheit to Celsius. Take temperature input from the user and display the result with two decimal places.

// 3. Integer to Float Conversion and Average Calculation:

// Take three integer inputs from the user, convert them to floats, and calculate the average using a function.

func intToFloat(i1 int, i2 int, i3 int) {
	f1 := float64(i1)
	f2 := float64(i2)
	f3 := float64(i3)

	var floatSum float64 = (f1 + f2 + f3) / 3

	fmt.Println("avarage of 3 int oin float is", floatSum)

}

// 4. Boolean to Integer Conversion:
// Write a function that converts a boolean value to an integer (true -> 1, false -> 0). Take user input as "true" or "false" and convert it.

func boolToInt(b bool) {
	fmt.Println(b, "Conversion of bool to int ")
	if b {
		fmt.Println(1)
	} else {
		fmt.Println(0)

	}

}

// 5. Convert Float to Integer and Check Even/Odd:
// Take a float input from the user, convert it to an integer, and determine whether the resulting integer is even or odd.

func evenOdd(f float64) {
	intf := int(f)
	if intf%2 == 0 {
		fmt.Println("is a even number")

	} else {
		fmt.Println("is a odd number")

	}
}

// 6. Hexadecimal to Decimal Converter:
// Ask the user to enter a hexadecimal string (e.g., "1A"), convert it to an integer, and display the decimal equivalent.

func HextoDeciaml(str string) {

	var ans float64 = 0
	var plac float64 = 0

	for i := len(str) - 1; i >= 0; i-- {
		x := str[i]

		inx, err := strconv.Atoi(string(x))

		if err != nil {
			if x == 'A' {
				ans += 10 * math.Pow(16, plac)
				plac++

			}
			if x == 'B' {
				ans += 11 * math.Pow(16, plac)
				plac++
			}
			if x == 'C' {
				ans += 12 * math.Pow(16, plac)
				plac++
			}
			if x == 'D' {
				ans += 13 * math.Pow(16, plac)
				plac++
			}
			if x == 'E' {
				ans += 14 * math.Pow(16, plac)
				plac++
			}
			if x == 'F' {
				ans += 16 * math.Pow(16, plac)
				plac++
			}
		} else {
			ans += float64(inx) * math.Pow(16, plac)
			plac++
		}

	}

	fmt.Println("hexadecimal ", ans)
}

// 7. Iterate Over a Slice of Strings and Convert to Integers:
// Given a slice of string numbers ([]string{"10", "20", "30", "40"}), write a program to convert them into integers and calculate their total sum.
func sliceadd(s []string) {
	vans := 0
	for i := 0; i < len(s); i++ {
		v, _ := strconv.Atoi(s[i])
		vans += v

	}
	println("sum of slice of string ", vans)
}

// 8. String to Float Conversion with Error Handling:
// Take a string input representing a floating-point number, convert it to a float, and print the square of the number. Handle any conversion errors.
func stringFloat(s string) {

	va, err := strconv.ParseFloat(s, 64)

	if err != nil {
		fmt.Println("string conversion failed")
	}

	fmt.Println("convert the string to var ", va)
}

// 9. Read Multiple Values and Convert Types:

// Read an integer, a float, and a string from the user. Convert the integer to a float and the float to an integer, then display the converted values.

func conversion() {
	var i int
	var f float64
	var s string

	fmt.Println("eneter one integer float and string ")
	fmt.Scanln(&i, &f, &s)

	ifloat := float64(i)
	ftoi := int(f)

	sum := ifloat + float64(ftoi)

	fmt.Println("sum of int and float", sum)

}

// 10. Character to ASCII Value Converter:

// Prompt the user to enter a single character, convert it to its ASCII integer value, and print the result.

func main() {

	s1 := "123"
	s2 := "2334"
	s3 := "34"
	s4 := "345"
	s5 := "66"

	i1, _ := strconv.Atoi(s1)
	i2, _ := strconv.Atoi(s2)
	i3, _ := strconv.Atoi(s3)
	i4, _ := strconv.Atoi(s4)
	i5, _ := strconv.Atoi(s5)
	i7 := i1 + i2 + i3 + i4 + i5

	intToFloat(2, 3, 4)
	boolToInt(true)
	evenOdd(766.09878)

	fmt.Println("sum of s1 and s2 s3 s4 s5 is ", i7)
	s := []string{"2", "3", "5"}
	sliceadd(s)

	stringFloat("44.44")
	conversion()

	HextoDeciaml("4A")

	var cha rune
	fmt.Println("Enter a character:")
	d, _ := fmt.Scanf("%c", &cha)

	m := int(d)
	fmt.Printf("The ASCII value of '%c' is: %d\n", cha, m)

}

// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	i := 2 // Index of element to delete
// 	fmt.Println(&s[2])
// 	s = slices.Delete(s, i, i+1)
// 	fmt.Println(s) // Output: [1 2 4 5]
// 	fmt.Println(&s[2])

// var c string
// 	fmt.Scanln(&c)
// 	fmt.Println(int(c[0]))

// }
