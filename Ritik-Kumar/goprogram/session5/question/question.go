package question

import (
	"math"
	"strconv"
	"strings"
)

// 1. Employee Records
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func GiveSalaryHike(emp *Employee, threshold float64, hike float64) {
	if emp.Salary < threshold {
		emp.Salary += hike
	}
}

// 2. Student Grade Management
type Student struct {
	Name  string
	Marks int
	Grade string
}

func AssignGrade(s *Student) {
	switch {
	case s.Marks >= 90:
		s.Grade = "A"
	case s.Marks >= 80:
		s.Grade = "B"
	case s.Marks >= 70:
		s.Grade = "C"
	default:
		s.Grade = "D"
	}
}

// 3. Rectangle Area Calculator
type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// 4. Circle Perimeter and Area

type Circle struct {
	Radius float64
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 5. Library Book Management
type Book struct {
	Title, Author, ISBN string
}

func SearchBook(books []Book, title string) *Book {
	for _, book := range books {
		if strings.EqualFold(book.Title, title) {
			return &book
		}
	}
	return nil
}

// 6. Phone Directory
var phoneDirectory = make(map[string]string)

func AddContact(name, number string) {
	phoneDirectory[name] = number
}

func DeleteContact(name string) {
	delete(phoneDirectory, name)
}

func SearchContact(name string) string {
	if num, exists := phoneDirectory[name]; exists {
		return num
	}
	return "Not Found"
}

// 7. Frequency of Digits in an Integer
func DigitFrequency(n int) map[int]int {
	freq := make(map[int]int)
	for n > 0 {
		digit := n % 10
		freq[digit]++
		n /= 10
	}
	return freq
}

// 8. Intersection of Two Arrays
func ArrayIntersection(arr1, arr2 []int) []int {
	m := make(map[int]bool)
	var intersection []int

	for _, num := range arr1 {
		m[num] = true
	}
	for _, num := range arr2 {
		if m[num] {
			intersection = append(intersection, num)
		}
	}
	return intersection
}

// 9. Find Duplicate Elements
func FindDuplicates(arr []int) []int {
	m := make(map[int]int)
	var duplicates []int

	for _, num := range arr {
		m[num]++
		if m[num] == 2 {
			duplicates = append(duplicates, num)
		}
	}
	return duplicates
}

// 10. Map Value Swapper
func SwapMap(m map[string]string) map[string]string {
	swapped := make(map[string]string)
	for key, value := range m {
		swapped[value] = key
	}
	return swapped
}

// 11. Second Largest Element
func SecondLargest(arr []int) int {
	first, second := math.MinInt, math.MinInt
	for _, num := range arr {
		if num > first {
			second = first
			first = num
		} else if num > second && num != first {
			second = num
		}
	}
	return second
}

// 12. Find Missing Numbers
func FindMissingNumbers(arr []int, n int) []int {
	missing := []int{}
	exists := make(map[int]bool)

	for _, num := range arr {
		exists[num] = true
	}
	for i := 1; i <= n; i++ {
		if !exists[i] {
			missing = append(missing, i)
		}
	}
	return missing
}

// 13. Split a Slice into Chunks
func ChunkSlice(arr []int, size int) [][]int {
	var chunks [][]int
	for size < len(arr) {
		arr, chunks = arr[size:], append(chunks, arr[0:size])
	}
	chunks = append(chunks, arr)
	return chunks
}

// 14. Sliding Window Maximum
func SlidingWindowMax(arr []int, k int) []int {
	if len(arr) == 0 {
		return nil
	}
	result := []int{}
	for i := 0; i <= len(arr)-k; i++ {
		max := arr[i]
		for j := i; j < i+k; j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		result = append(result, max)
	}
	return result
}

// 15. Remove All Occurrences
func RemoveElement(arr []int, val int) []int {
	var result []int
	for _, num := range arr {
		if num != val {
			result = append(result, num)
		}
	}
	return result
}

// 16. Type Conversion Utility
func IntToString(n int) string {
	return strconv.Itoa(n)
}

func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// 17. Sum of Digits
func SumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// 18. Swap Two Variables Without a Temp Variable
func Swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

// 19. Trailing Zeroes in Factorial
func TrailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}

// 20. Power of Two Checker
func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
