package main

import (
	"fmt"
	"math"
	"practise_Problems/circle"
	"practise_Problems/employee_management_system"
	"practise_Problems/library_management_system"
	"practise_Problems/rectangle"
	"practise_Problems/student_grade_management_system"
	"strconv"
)

func main() {
	//  1. Employee Records Management System:
	db := &employee_management_system.EmployeeDB

	employee_management_system.AddEmployee(db, 1, "Anubhav", 55000.00)
	employee_management_system.AddEmployee(db, 2, "Shirish", 45000.00)
	employee_management_system.AddEmployee(db, 3, "Rushil", 40000.00)
	employee_management_system.AddEmployee(db, 4, "Raushan", 42000.00)

	employee_management_system.GiveSalaryHike(db, 50000.00, 5)

	employee_management_system.DisplayEmployees(db)

	employee_management_system.RemoveEmployee(db, 4)

	// 2. Student Grade Management System:
	db2 := &student_grade_management_system.StudentDB

	student_grade_management_system.AddStudent(db2, "Anubhav", 95.0)
	student_grade_management_system.AddStudent(db2, "Shirish", 97.0)
	student_grade_management_system.AddStudent(db2, "Rushil", 85.0)
	student_grade_management_system.AddStudent(db2, "Raushan", 70.0)

	student_grade_management_system.DisplayStudents(db2)
	student_grade_management_system.RemoveStudent(db2, "Raushan")

	// 3. Rectangle Area Calculator:
	rectangle := rectangle.NewRectangle(30, 40)

	rectArea := rectangle.Area()
	fmt.Println("Area of the Rectangle is", rectArea)

	// 4. Circle Perimeter and Area:
	circle := circle.NewCircle(20)

	circleArea := circle.Area()
	circlePerimeter := circle.Perimeter()

	fmt.Println("Area of the Circle is", circleArea)
	fmt.Println("Perimeter of the Circle is", circlePerimeter)

	// 5. Library Management System:
	db3 := &library_management_system.BooksDB

	library_management_system.AddBook(db3, "Clean Code", "Robert C. Martin", "978-0132350884")
	library_management_system.AddBook(db3, "Go Programming", "Alan Donovan", "978-0134190440")

	library_management_system.DisplayBooks(db3)
	book, err := library_management_system.SearchBookByTitle(db3, "Go Programming")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Required Book:", book)
	}
	library_management_system.RemoveBook(db3, "978-0134190440")

	// 6. Phone Directory:
	phoneBook := make(map[string]string)

	AddContact(phoneBook, "Anubhav", "9876543210")
	AddContact(phoneBook, "Rushil", "9876234501")

	SearchContact(phoneBook, "Anubhav")
	DeleteContact(phoneBook, "Rushil")
	SearchContact(phoneBook, "Rushil")

	// 7. Frequency of Digits in an Integer:
	num := 12321
	digitMap := getDigitFrequency(num)

	fmt.Println("The Frequency of Digits are:", digitMap)

	// 8. Intersection of Two Arrays:
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 5, 6, 2, 8}

	intersection := arrayIntersection(arr1, arr2)
	fmt.Println("The Intersection of Two Arrays are", intersection)

	// 9. Find Duplicate Elements:
	arr3 := []int{1, 2, 2, 3, 4, 4, 4, 5}
	duplicateArr := findDuplicates(arr3)
	fmt.Println("The duplicate elements of the arr:", duplicateArr)

	// 10. Map Value Swapper:
	orgMap := map[string]string{
		"Anubhav":   "1",
		"Balmukund": "2",
		"Shirish":   "3",
	}

	swappedMap := swapMapVal(orgMap)
	fmt.Println("The Swapped Map is", swappedMap)

	// 11. Second Largest Element in a Slice:
	slc := []int{1, 2, 45, 69, 32, 23}
	secondLargest := findSecondLargestElement(slc)

	fmt.Println("The Second Largest Element in a Slice:", secondLargest)

	// 12. Find Missing Numbers in a Sequence:
	slc2 := []int{1, 3, 5, 7, 9}
	missingNums := findMissingNumbers(slc2, 9)
	fmt.Println("The Missing Numbers in the sequence are", missingNums)

	// 13. Split a Slice into Chunks:
	chunks := chunkSlice(slc2, 2)
	fmt.Println("The Chunks of the Slice are", chunks)

	// 14. Sliding Window Maximum:
	windowMaxArr := slidingWindowMaximum(slc, 2)
	fmt.Println("The maximum in every contiguous sub-slice of length k:", windowMaxArr)

	// 15. Remove All Occurrences of an Element:
	slc3 := []int{1, 2, 3, 2, 4, 2, 3}
	removeAllOccurrences(&slc3, 2)
	fmt.Println("Slice after removing all the Occurrences of an element", slc3)

	// 16. Type Conversion Utility:
	num2 := 16
	str := intToString(num2)
	fmt.Println(" Integer to String:", str)
	fmt.Printf("Type is %T\n\n", str)

	convInt, _ := stringToInt("101")
	fmt.Println(" String to Integer:", convInt)
	fmt.Printf("Type is %T\n\n", convInt)

	convFloat, _ := stringToFloat("99.99")
	fmt.Println(" String to Float:", convFloat)
	fmt.Printf("Type is %T\n\n", convFloat)

	// 17. Sum of Digits of an Integer:
	num3 := 12345
	digitSum := sumOfDigits(num3)
	fmt.Println("Sum of digits of the Integers is", digitSum)

	// 18. Swap Two Variables Without a Temp Variable:
	a, b := 4, 16
	a, b = swap(a, b)
	fmt.Println("Swapped variables a and b are", a, b)

	// 19. Trailing Zeroes in Factorial:
	fact := 25
	zeroCnt := trailingZeroes(fact)
	fmt.Println("Trailing Zeroes in Factorial", fact, "are", zeroCnt)

	// 20. Power of Two Checker:
	num4 := 6
	isPowTwo := isPowerOfTwo(num4)
	fmt.Println("The Given Number", num4, "is Power of Two?", isPowTwo)

}

func AddContact(phoneBook map[string]string, name, phone string) {
	phoneBook[name] = phone
	fmt.Println(" Contact added:", name, "->", phone)
}

func DeleteContact(phoneBook map[string]string, name string) {
	if _, exists := phoneBook[name]; exists {
		delete(phoneBook, name)
		fmt.Println(" Contact deleted:", name)
	} else {
		fmt.Println(" Contact not found:", name)
	}
}

func SearchContact(phoneBook map[string]string, name string) {
	if phone, exists := phoneBook[name]; exists {
		fmt.Println(" Found:", name, "->", phone)
	} else {
		fmt.Println(" Contact not found:", name)
	}
}

func getDigitFrequency(num int) map[int]int {
	digitMap := make(map[int]int)

	for num > 0 {
		digit := num % 10
		num /= 10
		digitMap[digit]++
	}

	return digitMap
}

func arrayIntersection(arr1, arr2 []int) []int {
	elements := make(map[int]bool)
	var intersection []int

	for _, num := range arr1 {
		elements[num] = true
	}

	for _, num := range arr2 {
		if elements[num] {
			intersection = append(intersection, num)
			elements[num] = false
		}
	}

	return intersection
}

func findDuplicates(arr []int) []int {
	freq := make(map[int]int)
	duplicates := []int{}

	for _, num := range arr {
		freq[num]++
		if freq[num] == 2 {
			duplicates = append(duplicates, num)
		}
	}

	return duplicates
}

func swapMapVal(orgMap map[string]string) map[string]string {
	swappedMap := make(map[string]string)

	for key, value := range orgMap {
		swappedMap[value] = key
	}

	return swappedMap
}

func findSecondLargestElement(slc []int) int {
	maxi := slc[0]
	secondLargest := -1

	for _, value := range slc {
		if maxi < value {
			secondLargest = maxi
			maxi = value
		}
		if value < maxi && value > secondLargest {
			secondLargest = value
		}
	}

	return secondLargest
}

func findMissingNumbers(slc []int, n int) []int {
	found := make(map[int]bool)
	missingNums := []int{}

	for _, num := range slc {
		found[num] = true
	}

	for i := 1; i <= n; i++ {
		if !found[i] {
			missingNums = append(missingNums, i)
		}
	}

	return missingNums
}

func chunkSlice(nums []int, k int) [][]int {
	if k <= 0 {
		return nil
	}

	var chunks [][]int
	for i := 0; i < len(nums); i += k {
		end := i + k
		if end > len(nums) {
			end = len(nums)
		}
		chunks = append(chunks, nums[i:end])
	}

	return chunks
}

func slidingWindowMaximum(slc []int, k int) []int {
	windowMaxArr := []int{}

	n := len(slc)
	i, j := 0, k-1

	for j < n {
		maxi := math.MinInt
		for k := i; k <= j; k++ {
			if slc[k] > maxi {
				maxi = slc[k]
			}
		}
		windowMaxArr = append(windowMaxArr, maxi)
		i++
		j++
	}

	return windowMaxArr
}

func removeAllOccurrences(slce *[]int, target int) {
	slc := *slce
	count := 0
	for i, value := range slc {
		if value == target {
			slc = append(slc[:i], slc[i+1:]...)
			count++
		}
	}
	println(count, len(slc))

	*slce = slc[:len(slc)]
}

func intToString(num int) string {
	return strconv.Itoa(num)
}

func stringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func stringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func sumOfDigits(n int) int {
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func swap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b

	// Method 2:
	// a, b = b, a

	// Method 3:
	// return b, a
	return a, b
}

func trailingZeroes(n int) int {
	count := 0
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
