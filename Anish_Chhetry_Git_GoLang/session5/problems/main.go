package main

import (
	"fmt"
	"strconv"
)

type employee struct {
	id     int
	name   string
	salary float32
}

func salaryHike(emp []employee, threshold float32, hike float32) {
	for i := range emp {
		if emp[i].salary <= threshold {
			emp[i].salary *= 1 + hike
		}
	}
}

type student struct {
	name  string
	marks int
	grade string
}

func assignGrades(stu []student) {
	for i := range stu {
		if stu[i].marks >= 90 {
			stu[i].grade = "A"
		} else if stu[i].marks >= 80 {
			stu[i].grade = "B"
		} else if stu[i].marks >= 70 {
			stu[i].grade = "C"
		} else if stu[i].marks >= 60 {
			stu[i].grade = "D"
		} else {
			stu[i].grade = "F"
		}
	}
}

type rectangle struct {
	length float32
	width  float32
	area   float32
}

func rectArea(rect []rectangle) {
	for i := range rect {
		rect[i].area = rect[i].length * rect[i].width
	}
}

type circle struct {
	radius    float32
	perimeter float32
	area      float32
}

func circAreaPeri(circ []circle) {
	for i := range circ {
		circ[i].area = 3.14 * circ[i].radius * circ[i].radius
		circ[i].perimeter = 2 * 3.14 * circ[i].radius
	}
}

type books struct {
	title  string
	author string
	isbn   int
}

func findbook(book []books, title string) {
	for i := range book {
		if book[i].title == title {
			fmt.Println(book[i])
		}
	}
}

func phoneadd(pmap map[string]int, name string, phone int) {
	pmap[name] = phone
}
func phonedelete(pmap map[string]int, name string) {
	delete(pmap, name)
}
func phonesearch(pmap map[string]int, name string) {
	fmt.Println(pmap[name])
}

func digiFreq(num string) {
	m := make(map[int]int)
	numlist := []int{}
	for _, digit := range num {
		value, _ := strconv.Atoi(string(digit))
		numlist = append(numlist, value)
	}
	for i := 0; i < len(numlist); i++ {
		switch numlist[i] {
		case 1:
			m[1]++
		case 2:
			m[2]++
		case 3:
			m[3]++
		case 4:
			m[4]++
		case 5:
			m[5]++
		case 6:
			m[6]++
		case 7:
			m[7]++
		case 8:
			m[8]++
		case 9:
			m[9]++
		case 0:
			m[0]++
		}

	}
	fmt.Println(m)
}

func intersection(slice1 []int, slice2 []int) []int {
	m := make(map[int]int)
	for _, num := range slice1 {
		m[num]++
	}
	inter := []int{}
	for _, num := range slice2 {
		if m[num] > 0 {
			inter = append(inter, num)
		}
	}
	return inter
}

func dup(slice []int) []int {
	m := make(map[int]int)
	dups := []int{}

	for _, num := range slice {
		m[num]++
		if m[num] > 1 {
			dups = append(dups, num)
		}
	}
	return dups
}
func swapmap(m1 map[string]string) map[string]string {
	m2 := make(map[string]string)
	for i, j := range m1 {
		m2[j] = i
	}
	return m2
}
func secondlargest(slice []int) int {
	largest := 0
	slargest := 0
	for i := range slice {
		if slice[i] > largest {
			slargest = largest
			largest = slice[i]
		} else if i > slargest && slice[i] != largest {
			slargest = slice[i]
		}
	}
	return slargest
}
func findMissingNumber(slice []int, n int) []int {
	m := make(map[int]bool)
	newslice := []int{}
	for _, i := range slice {
		m[i] = true
	}
	for i := 1; i <= n; i++ {
		if !m[i] {
			newslice = append(newslice, i)
		}
	}

	return newslice
}
func chunks(slice []int, k int) {
	for i := 0; i < len(slice); i += k {
		end := i + k
		if end > len(slice) {
			end = len(slice)
		}
		fmt.Println(slice[i:end])
	}
}
func slidingWindow(slice []int, k int) {
	for i := 0; i <= len(slice)-k; i++ {
		end := i + k
		if end > len(slice) {
			end = len(slice)
		}
		fmt.Println(slice[i:end], max(slice[i:end]))

	}
}
func max(slice []int) int {
	max := 0
	for i := 0; i < len(slice); i++ {
		for j := 1; j < len(slice); j++ {
			if slice[i] > max {
				max = slice[i]
			}
		}
	}
	return max
}
func removeOcc(slice []int, value int) []int {
	newslice := []int{}
	for _, i := range slice {
		if i != value {
			newslice = append(newslice, i)
		}
	}
	return newslice
}
func intToString(i int) string {
	return strconv.Itoa(i)
}

func stringToFloat(s string) float64 {
	value, _ := strconv.ParseFloat(s, 64)
	return value
}

func stringToInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
func digiSum(num string) {
	numlist := []int{}
	sum := 0
	for _, digit := range num {
		value, _ := strconv.Atoi(string(digit))
		numlist = append(numlist, value)
	}
	for _, i := range numlist {
		sum += i
	}
	fmt.Println(sum)
}
func swapvar(var1, var2 int) (int, int) {
	var1, var2 = var2, var1
	return var1, var2
}
func fact(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}
func trail0(num int) int {
	count := 0
	fact := fact(num)
	fmt.Println("Factorial of", num, "is :", fact)
	for i := 1; i != 0; {
		if fact%10 == 0 {
			count++
			fact /= 10
		} else {
			return count
		}
	}
	return count
}
func isPowerOf2(num int) bool {
	for i := 0; i < num/2; i++ {
		if 2<<i == num {
			return true
		}
	}
	return false
}
func main() {
	// 	1.	Employee Records:
	// Create a struct for an Employee with fields like ID, Name, and Salary. Write a function to give a salary hike to employees earning below a certain threshold.
	emp := []employee{{1, "Anish", 300}, {2, "Kishor", 400}, {3, "Mega", 200}, {4, "Aditya", 500}}
	fmt.Println("Before Salary Hike: \n", emp)
	salaryHike(emp, 350, 0.2)
	fmt.Println("After Salary Hike: \n", emp)

	// 	2.	Student Grade Management:
	// Define a Student struct with Name, Marks, and Grade. Write a function to assign grades based on marks.
	stu := []student{{"Anish", 70, ""}, {"Kishor", 80, ""}, {"Megha", 90, ""}}
	fmt.Println("Before student grade assignment: \n", stu)
	assignGrades(stu)
	fmt.Println("After student grade assignment: \n", stu)
	// 	3.	Rectangle Area Calculator:
	// Create a Rectangle struct with Length and Width. Write a method to calculate the area.
	rect := []rectangle{{10, 23, 0}, {99, 22, 0}, {12, 82, 0}}
	fmt.Println("Before area calculation: \n", rect)
	rectArea(rect)
	fmt.Println("After area calculation: \n", rect)
	// 	4.	Circle Perimeter and Area:
	// Define a Circle struct with a Radius field. Write methods to calculate the circumference and area.
	circ := []circle{{10, 0, 0}, {99, 0, 0}, {12, 0, 0}}
	fmt.Println("Before area,perimeter calculation: \n", circ)
	circAreaPeri(circ)
	fmt.Println("After area, perimeter calculation: \n", circ)
	// 	5.	Library Book Management:
	// Create a Book struct with fields Title, Author, and ISBN. Write a function to search for a book by title in a slice of books.
	book := []books{{"The Forest", "James", 123}, {"The wind in the air", "Jack", 456}}
	fmt.Println("Available Books: \n", book)
	titlename := "The Forest"
	findbook(book, titlename)
	// 	6.	Phone Directory:
	// Create a phone directory using a map where names are keys and phone numbers are values. Implement functions to add, delete, and search for contacts.
	phone := make(map[string]int)
	phoneadd(phone, "Anish", 123)
	phoneadd(phone, "Mega", 5762)
	phoneadd(phone, "Aditya", 87356373)
	fmt.Println(phone)
	phonedelete(phone, "Mega")
	fmt.Println(phone)
	phonesearch(phone, "Aditya")
	// 	7.	Frequency of Digits in an Integer:
	// Given an integer, return a map containing the frequency of each digit.
	num := 111223
	slice := strconv.Itoa(num)
	digiFreq(slice)

	// 	8.	Intersection of Two Arrays:
	// Write a function that finds the intersection of two integer slices using maps.
	intslice1 := []int{1, 2, 3, 4, 5}
	intslice2 := []int{4, 5, 6, 7, 8}
	fmt.Println(intersection(intslice1, intslice2))
	// 	9.	Find Duplicate Elements:
	// Given a slice of integers, use a map to find and return the list of duplicates.
	intslice3 := []int{1, 1, 3, 3, 5}
	fmt.Println(dup(intslice3))
	// 	10.	Map Value Swapper:
	// Given a map of string keys and string values, create a new map where keys and values are swapped.
	m1 := map[string]string{"a": "abc", "x": "xyz"}
	fmt.Println(m1)
	fmt.Println(swapmap(m1))

	// 	11.	Second Largest Element in a Slice:
	// Write a function to find the second largest element in a slice of integers.
	slice3 := []int{13, 2, 3, 41}
	fmt.Println(secondlargest(slice3))
	// 	12.	Find Missing Numbers in a Sequence:
	// Given a slice of integers from 1 to n with some numbers missing, find all missing numbers.
	slice4 := []int{1, 2, 4, 5}
	n := 10
	fmt.Println("Original Array:", slice4)
	fmt.Println("Missing number in the range: ", findMissingNumber(slice4, n))

	// 	13.	Split a Slice into Chunks:
	// Write a function to split a slice into smaller chunks of size k.
	slice5 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 2
	chunks(slice5, k1)

	// 	14.	Sliding Window Maximum:
	// Implement a function to find the maximum in every contiguous sub-slice of length k.
	sliced := []int{4, 3, 2, 7, 21, 77, 22, 23}
	k2 := 4
	slidingWindow(sliced, k2)

	// 	15.	Remove All Occurrences of an Element:
	// Given a slice and a value, remove all occurrences of that value without creating a new slice.
	slice6 := []int{1, 2, 3, 4, 5, 1, 1}
	value := 1
	fmt.Println(removeOcc(slice6, value))
	// 	16.	Type Conversion Utility:
	// Write functions to convert between basic data types (int to string, string to float, etc.).
	fmt.Println(intToString(12))
	fmt.Println(stringToFloat("123.2"))
	fmt.Println(stringToInt("12"))
	// 	17.	Sum of Digits of an Integer:
	// Create a function that calculates the sum of all digits in an integer.
	num2 := 111223
	slice7 := strconv.Itoa(num2)
	digiSum(slice7)

	// 	18.	Swap Two Variables Without a Temp Variable:
	// Implement a function to swap two integers without using a third variable.
	var1 := 1
	var2 := 2
	fmt.Println("Before:", var1, var2)
	var1, var2 = swapvar(var1, var2)
	fmt.Println("After:", var1, var2)
	// 	19.	Trailing Zeroes in Factorial:
	// Write a function to find the number of trailing zeroes in the factorial of a number.
	num3 := 18
	fmt.Println(trail0(num3))
	// 	20.	Power of Two Checker:
	// Write a function to check if a given integer is a power of two.
	num4 := 16
	fmt.Println(isPowerOf2(num4))
}
