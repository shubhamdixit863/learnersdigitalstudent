package main

import (
	"fmt"
	"math"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}
type Student struct {
	Name  string
	Marks int
	Grade string
}
type Rectangle struct {
	length  float64
	breadth float64
}
type Circle struct {
	radius        float64
	circumference float64
	area          float64
}

type Book struct {
	Title  string
	Author string
	ISBN   int
}

var contacts = make(map[string]string)

func main() {
	//1
	employees := []Employee{
		{ID: 1, Name: "Alice", Salary: 40000},
		{ID: 2, Name: "Bob", Salary: 80000},
	}

	threshold := 50000.0
	hikePercent := 10.0
	for e := range employees {
		GiveSalaryHike(&employees[e], threshold, hikePercent)
	}
	//2
	students := []Student{
		{Name: "Alice", Marks: 70},
		{Name: "Bob", Marks: 70},
	}
	for s := range students {
		AssignGrade(&students[s])
	}

	for s := range students {
		fmt.Printf("\nStudent name: %s, Student Marks: %d, Student Grade: %s", students[s].Name, students[s].Marks, students[s].Grade)
	}
	//3
	rec1 := Rectangle{length: 3.5, breadth: 2}
	fmt.Println("\n\nArea of rect is:", Area(rec1))

	//4
	cir1 := Circle{radius: 5}
	CircleCircumArea(&cir1)
	fmt.Printf("\nCircumference is: %.2f and Area is: %.2f", cir1.circumference, cir1.area)

	//5
	books := []Book{
		{Title: "Harry Potter", Author: "J.K Rowling", ISBN: 1234},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee", ISBN: 5678},
		{Title: "1984", Author: "George Orwell", ISBN: 9876},
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", ISBN: 3456},
	}
	title := "Harry Potter"
	fmt.Printf("\n\nSearching for book: %s", title)
	book, err := SearchBook(books, title)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nBooks found! Title:%s, Author:%s, ISBN:%d", book.Title, book.Author, book.ISBN)

	//6
	AddContact("Vicky", "12345678")
	AddContact("Alice", "987654")
	AddContact("Bob", "547654")
	AddContact("Charlie", "344354")
	name, contact, ok := GetContact("Vicky")
	if ok {
		fmt.Println("\nContact found", name, contact)
	} else {
		fmt.Println("\nContact not found")
	}
	DeleteContact("Vicky")
	name2, contact2, ok2 := GetContact("Vicky")
	if ok2 {
		fmt.Println("\nContact found", name2, contact2)
	} else {
		fmt.Println("\nContact not found")
	}

	//7
	fmt.Println("Frequency of digits in 1234512 is", DigitFreq(1234512))
	//8
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}
	arr2 := []int{2, 4, 6, 8, 10, 11, 12, 13}
	itrsec := Intersection(arr1, arr2)
	fmt.Println("intersection of arr: ", itrsec)

	//9
	fmt.Println("Duplicates in arr: ", Duplicate(arr1))

	//10
	var m = map[string]string{
		"name":  "vicky",
		"batch": "golang",
	}

	//11
	fmt.Println("Second Largest: ", SecondLargest(arr1))
	fmt.Println("Swapped map is:", SwapMap(m))

	//12
	fmt.Println("Missing numbers:", FindMissingNumbers(arr2, 13))

	//13
	arr3 := []int{1, 2, 3, 4, 5, 6, 7, 2, 2, 4, 2}
	fmt.Println("Chunked Array of size 2:", SplitSplice(arr3, 2))
	//15
	arr := []int{1, 2, 3, 4, 5, 6, 7, 2, 2, 4, 2}
	RemoveOccurance(&arr, 2)
	fmt.Println("Remove occurance of 2 in arr", arr)

	//17
	fmt.Println("Sum of Digit in 12345 is,", SumofDigit(12345))

	//18
	a := 4
	b := 5
	fmt.Printf("Before Swapping, a: %d, b: %d\n", a, b)
	a, b = Swap(a, b)
	fmt.Printf("Afters Swapping, a: %d, b: %d\n", a, b)

	//19
	fmt.Println("Trailing Zeroes of 100:", TrailingZeroes(100))
	//20
	fmt.Println("Is 32 power of two", PowerOfTwo(32))
	fmt.Println("Is 40 power of two", PowerOfTwo(40))
}

// 1.	Employee Records:
// Create a struct for an Employee with fields like ID, Name, and Salary. Write a function to give a salary hike to employees earning below a certain threshold.
func GiveSalaryHike(emp *Employee, threshold, hikePercent float64) {
	if emp.Salary < threshold {
		hikeAmount := emp.Salary * (hikePercent / 100)
		emp.Salary += hikeAmount
		fmt.Printf("Salary hiked for %s by %f%%. New Salary: %f\n", emp.Name, hikePercent, emp.Salary)
	} else {
		fmt.Printf("No hike for %s. Current Salary: %f\n", emp.Name, emp.Salary)
	}
}

// 	2.	Student Grade Management:=
// Define a Student struct with Name, Marks, and Grade. Write a function to assign grades based on marks.

func AssignGrade(stu *Student) {
	if stu.Marks > 90 {
		stu.Grade = "A"
	} else if stu.Marks > 80 {
		stu.Grade = "B"
	} else if stu.Marks > 70 {
		stu.Grade = "C"
	} else if stu.Marks > 60 {
		stu.Grade = "D"
	} else if stu.Marks > 40 {
		stu.Grade = "E"
	} else {
		stu.Grade = "F"
	}
}

//  3. Rectangle Area Calculator:
//
// Create a Rectangle struct with Length and Width. Write a method to calculate the area.
func Area(rec Rectangle) float64 {
	return rec.breadth * rec.length
}

//  4. Circle Perimeter and Area:
//
// Define a Circle struct with a Radius field. Write methods to calculate the circumference and area.
func CircleCircumArea(cir *Circle) {
	cir.area = (2 * math.Pi * cir.radius)
	cir.circumference = (math.Pi * cir.radius * cir.radius)
}

//  5. Library Book Management:
//
// Create a Book struct with fields Title, Author, and ISBN. Write a function to search for a book by title in a slice of books.
func SearchBook(books []Book, title string) (Book, error) {
	for _, b := range books {
		if b.Title == title {
			return b, nil
		}
	}
	return Book{}, fmt.Errorf("book not found")
}

//  6. Phone Directory:
//
// Create a phone directory using a map where names are keys and phone numbers are values. Implement functions to add, delete, and search for contacts.
func AddContact(name string, number string) {
	contacts[name] = number
}

func GetContact(name string) (string, string, bool) {
	number, found := contacts[name]
	return name, number, found
}

func DeleteContact(name string) {
	delete(contacts, name)
}

//  7. Frequency of Digits in an Integer:
//
// Given an integer, return a map containing the frequency of each digit.
func DigitFreq(num int) map[int]int {
	freq := make(map[int]int)
	for num > 0 {
		digit := num % 10
		freq[digit]++
		num /= 10
	}
	return freq
}

//  8. Intersection of Two Arrays:
//
// Write a function that finds the intersection of two integer slices using maps.
func Intersection(a, b []int) []int {
	m := make(map[int]bool)
	for _, num := range a {
		m[num] = true
	}

	intersection := make([]int, 0)
	for _, num := range b {
		if m[num] {
			intersection = append(intersection, num)
		}
	}
	return intersection
}

//  9. Find Duplicate Elements:
//
// Given a slice of integers, use a map to find and return the list of duplicates.
func Duplicate(a []int) []int {
	m := make(map[int]int)
	for _, num := range a {
		m[num]++
	}
	var duplicates []int
	for k, v := range m {
		if v > 1 {
			duplicates = append(duplicates, k)
		}
	}
	return duplicates
}

//  10. Map Value Swapper:
//
// Given a map of string keys and string values, create a new map where keys and values are swapped.
func SwapMap(m map[string]string) map[string]string {
	swapped := make(map[string]string)
	for k, v := range m {
		swapped[v] = k
	}
	return swapped
}

// 11.	Second Largest Element in a Slice:
// Write a function to find the second largest element in a slice of integers.
func SecondLargest(arr []int) int {
	maxm := math.MinInt
	secMaxm := -1
	for _, n := range arr {
		if n > maxm {
			secMaxm = maxm
			maxm = n
		} else if n > secMaxm && n != maxm {
			secMaxm = n
		}
	}
	return secMaxm
}

//  12. Find Missing Numbers in a Sequence:
//
// Given a slice of integers from 1 to n with some numbers missing, find all missing numbers.
func FindMissingNumbers(arr []int, n int) []int {
	missing := make(map[int]bool)
	for i := 1; i <= n; i++ {
		missing[i] = true
	}
	for _, n := range arr {
		missing[n] = false
	}

	var result []int
	for k, v := range missing {
		if v {
			result = append(result, k)
		}
	}
	return result
}

//  13. Split a Slice into Chunks:
//
// Write a function to split a slice into smaller chunks of size k.
func SplitSplice(arr []int, k int) [][]int {
	var ans [][]int
	for i := 0; i < len(arr); {
		var chunk []int
		for j := 0; j < k && i < len(arr); j++ {
			chunk = append(chunk, arr[i])
			i++
		}
		ans = append(ans, chunk)
	}
	return ans

}

//  14. Sliding Window Maximum:
//
// Implement a function to find the maximum in every contiguous sub-slice of length k.
// func SlidingWindowMaxm(arr []int, k int) []int {
// 	maxm := math.MinInt
// 	var ans []int
// 	for i := 0; i < k; i++ {
// 		maxm = max(maxm, arr[i])
// 	}
// 	ans = append(ans, maxm)
// 	start:=0
// 	end:=k
// 	for end<len(arr){
// 		if arr[end]>maxm
// 	}
// }

//  15. Remove All Occurrences of an Element:
//
// Given a slice and a value, remove all occurrences of that value without creating a new slice.
func RemoveOccurance(arr *[]int, num int) {
	for i := 0; i < len(*arr); {
		if (*arr)[i] == num {
			*arr = append((*arr)[:i], (*arr)[i+1:]...)
		} else {
			i++
		}
	}

}

// 	16.	Type Conversion Utility:
// Write functions to convert between basic data types (int to string, string to float, etc.).

//  17. Sum of Digits of an Integer:
//
// Create a function that calculates the sum of all digits in an integer.
func SumofDigit(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

//  18. Swap Two Variables Without a Temp Variable:
//
// Implement a function to swap two integers without using a third variable.
func Swap(a int, b int) (int, int) {
	// a = a + b
	// b = a - b
	// a = a - b
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

//  19. Trailing Zeroes in Factorial:
//
// Write a function to find the number of trailing zeroes in the factorial of a number.
func TrailingZeroes(num int) int {
	count := 0
	for num >= 5 {
		num /= 5
		count += num
	}
	return count
}

//  20. Power of Two Checker:
//
// Write a function to check if a given integer is a power of two.
func PowerOfTwo(num int) bool {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count == 1
}
