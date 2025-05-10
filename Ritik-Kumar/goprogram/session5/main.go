package main

import (
	"fmt"
	"session5/question"
)

func main() {
	// 1. Employee Salary Hike
	emp := question.Employee{ID: 1, Name: "Alice", Salary: 40000}
	question.GiveSalaryHike(&emp, 50000, 5000)
	fmt.Println("Updated Salary:", emp.Salary)

	// 2. Student Grade Management
	student := question.Student{Name: "Bob", Marks: 85}
	question.AssignGrade(&student)
	fmt.Println("Student Grade:", student.Name, "->", student.Grade)

	// 3. Rectangle Area
	rect := question.Rectangle{Length: 10, Width: 5}
	fmt.Println("Rectangle Area:", rect.Area())

	// 4. Circle Perimeter and Area
	circle := question.Circle{Radius: 7}
	fmt.Println("Circle Circumference:", circle.Circumference())
	fmt.Println("Circle Area:", circle.Area())

	// 5. Library Book Management
	books := []question.Book{
		{Title: "Go Programming", Author: "John", ISBN: "12345"},
		{Title: "Data Structures", Author: "Jane", ISBN: "67890"},
	}
	book := question.SearchBook(books, "Go Programming")
	if book != nil {
		fmt.Println("Book found:", book.Title, "by", book.Author)
	} else {
		fmt.Println("Book not found")
	}

	// 6. Phone Directory
	question.AddContact("Alice", "9876543210")
	question.AddContact("Bob", "8765432109")
	fmt.Println("Bob's Phone Number:", question.SearchContact("Bob"))
	question.DeleteContact("Bob")
	fmt.Println("Bob's Phone Number after deletion:", question.SearchContact("Bob"))

	// 7. Frequency of Digits
	fmt.Println("Digit Frequency of 112358:", question.DigitFrequency(112358))

	// 8. Intersection of Two Arrays
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 4, 5, 6, 7}
	fmt.Println("Intersection:", question.ArrayIntersection(arr1, arr2))

	// 9. Find Duplicates
	arr := []int{1, 2, 3, 2, 4, 5, 1, 6}
	fmt.Println("Duplicate Elements:", question.FindDuplicates(arr))

	// 10. Map Value Swapper
	m := map[string]string{"a": "apple", "b": "banana"}
	fmt.Println("Swapped Map:", question.SwapMap(m))

	// 11. Second Largest Element
	fmt.Println("Second Largest:", question.SecondLargest(arr1))

	// 12. Find Missing Numbers
	fmt.Println("Missing Numbers:", question.FindMissingNumbers([]int{1, 2, 4, 5, 7}, 7))

	// 13. Split a Slice into Chunks
	fmt.Println("Chunked Slice:", question.ChunkSlice(arr1, 2))

	// 14. Sliding Window Maximum
	fmt.Println("Sliding Window Max:", question.SlidingWindowMax(arr1, 3))

	// 15. Remove All Occurrences
	fmt.Println("Array after removal:", question.RemoveElement(arr, 2))

	// 16. Type Conversion Utility
	fmt.Println("Int to String:", question.IntToString(123))
	strToFloat, _ := question.StringToFloat("45.67")
	fmt.Println("String to Float:", strToFloat)

	// 17. Sum of Digits
	fmt.Println("Sum of Digits:", question.SumOfDigits(1234))

	// 18. Swap Two Variables Without a Temp Variable
	a, b := 5, 10
	a, b = question.Swap(a, b)
	fmt.Println("Swapped Values:", a, b)

	// 19. Trailing Zeroes in Factorial
	fmt.Println("Trailing Zeroes in 100!:", question.TrailingZeroes(100))

	// 20. Power of Two Checker
	fmt.Println("Is 16 a power of two?", question.IsPowerOfTwo(16))
	fmt.Println("Is 18 a power of two?", question.IsPowerOfTwo(18))
}
