package main

import (
	"fmt"
	"strconv"
)

func sToI() {
	var a, b, c, d, e string

	fmt.Println("Enter the strings: ")
	fmt.Scan(&a, &b, &c, &d, &e)

	var a1, b1, c1, d1, e1 int

	a1, _ = strconv.Atoi(a)
	b1, _ = strconv.Atoi(b)
	c1, _ = strconv.Atoi(c)
	d1, _ = strconv.Atoi(d)
	e1, _ = strconv.Atoi(e)

	fmt.Printf("Type of a: %T, b: %T, c: %T, d: %T, e: %T", a1, b1, c1, d1, e1)
}

func tempConv(c float64) string {
	f := c*(9/5) + 32
	return fmt.Sprintf("%.2f", f)
}

func avg(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func boolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func evenOdd(a float64) string {
	temp := int(a)
	if temp%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func hexToDec(hex string) int64 {
	temp, _ := strconv.ParseInt(hex, 16, 64)
	return temp
}

func slcSum(s []string) int {
	sum := 0
	for i, _ := range s {
		a, _ := strconv.Atoi(s[i])
		sum += a
	}
	return sum
}

func floatToSquare(s string) (float64, error) {

	f, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return 0, err
	}

	square := f * f

	return square, nil
}

func readMultiple(num int, f float64, s string) {

	convF := float64(num)
	convI := int(f)

	fmt.Printf("The converted float is %f and int is %d and string input was %s", convF, convI, s)
}

func charToAscii(s string) int {
	if len(s) != 1 {
		fmt.Println("Please enter only one character.")
		return 0
	}

	asciiValue := int(s[0])
	return asciiValue

}

func main() {
	// Que: 1
	//sToI()

	// Que: 2
	//var c float64
	//fmt.Scan(&c)
	//fmt.Println(tempConv(c))

	// Que: 3
	//var a, b, c int
	//fmt.Scan(&a, &b, &c)
	//
	//f1 := float64(a)
	//f2 := float64(b)
	//f3 := float64(c)
	//
	//fmt.Println(avg(f1, f2, f3))

	// Que: 4
	//var b bool
	//fmt.Scan(&b)
	//fmt.Println(boolToInt(b))

	// Que: 5
	//var a float64
	//fmt.Scan(&a)
	//fmt.Println(evenOdd(a))

	// Que: 6
	//var hex string
	//fmt.Scan(&hex)
	//fmt.Println("Decimal is: ", hexToDec(hex))

	// Que: 7
	//slc := []string{"10", "20", "30", "40"}
	//fmt.Println(slcSum(slc))

	// Que: 8
	//sq, err := floatToSquare("2.5")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(sq)

	// Que: 9
	//var (
	//	a int
	//	b float64
	//	c string
	//)
	//fmt.Println("Enter the Integer: ")
	//fmt.Scan(&a)
	//fmt.Println("Enter the Float: ")
	//fmt.Scan(&b)
	//fmt.Println("Enter the String: ")
	//fmt.Scan(&c)
	//
	//readMultiple(a, b, c)

	// Que: 10
	var a string
	fmt.Println("Enter a Single Character: ")
	fmt.Scan(&a)

	fmt.Println(charToAscii(a))
}
