package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Reverse an array:
	var arr = [5]int{1, 2, 3, 4, 5}
	revArr1 := reverse_Array1(arr)
	fmt.Println("The Reversed array is : ", revArr1)

	revArr2 := reverse_Array2(arr)
	fmt.Println("(Method 2) The Reversed array is : ", revArr2)

	// 2. Sum of Elements in a Slice:
	var slc = []int{1, 2, 3, 4, 5}
	slc = append(slc, 6)

	sumElem := slice_Element_Sum(slc)
	fmt.Println("The Sum of Elements in a Slice is : ", sumElem)

	// 3. Find Maximum and Minimum in an Array:
	var arr2 = [6]int{3, 7, 9, 15, 78, 1}
	maxi, mini := max_min_Array(arr2)
	fmt.Println("The Maximum and Minimum in an array: ", maxi, mini)

	// 4. Remove Duplicates from a Slice:
	var slc2 = []int{1, 2, 2, 5, 4, 1, 3, 4, 5}
	slice_set := remove_duplicate_slice(slc2)
	fmt.Println("The Slice without Duplicates is: ", slice_set)

	str := "Hello"
	fmt.Println([]byte(str))
	fmt.Println([]rune(str))

	// 5. Rotate Array by k Steps:
	rotated_Arr := rotate_Array_K(arr2, 2)
	fmt.Println("The Rotated Array is", rotated_Arr)

	// 6. Find All Even Numbers in a Slice:
	var slc3 = []int{1, 2, 6, 5, 4, 9, 3, 8, 7}
	evenNums := find_Even_Numbers(slc3)
	fmt.Println("All The Even Numbers in a Slice are : ", evenNums)

	// 7. Count Occurrences of an Element:
	var slc4 = []int{1, 2, 2, 2, 2, 6, 5, 4, 3}
	elem_Count := count_Occur_Elem(slc4, 2)
	fmt.Println("The Occurrences of the element", 2, "is", elem_Count)

	// 8. Word Frequency Counter:
	str2 := "Hello World, Hello hello , I am Anubhav"
	word_map := word_Freq(str2)
	fmt.Println("Word Frequency Count: ", word_map)

	// 9. Character Frequency in a String:
	str3 := "HelloWorld"
	char_map := char_Freq(str3)
	fmt.Println("Character Frequency Count: ", char_map)

	// 10. Two Sum Problem:
	var arr3 = [6]int{3, 7, 9, 15, 78, 6}
	indices := two_Sum(arr3, 15)
	if indices == nil {
		fmt.Println("Error: No Two elements adds up to the Target..")
	} else {
		fmt.Println("Indices of elements that adds up to the target are ", indices)
	}

	// 11. Check for Duplicates in an Array:
	var arr4 = [6]int{3, 7, 9, 6, 8, 6}
	duplicates_found := duplicates_In_Array(arr4)
	fmt.Println("Are there any Duplicates in the Array?", duplicates_found)

	// 12. First Unique Character in a String:
	fst_uniq_ch := first_Unique_Char(str3)
	fmt.Println("First Unique Character in a String is at", fst_uniq_ch)

	// 13. Fibonacci Sequence Generator:
	n := 5
	fib_Seq := fibonacci(n)
	fmt.Println("First", n, "Fibonacci Numbers are", fib_Seq)

	// 14. Palindrome Number Checker:
	n1 := 121
	palindrome_Found := isPalindrome(n1)
	fmt.Println("The Number", n1, "is Palindrome?", palindrome_Found)

	// 15. Find Missing Number in a Range:
	var arr6 = [6]int{1, 2, 3, 4, 6, 7}
	missing_Num := find_Missing_Number(arr6)
	fmt.Println("The Missing Number from 1 to N is", missing_Num)

}

func reverse_Array1(arr [5]int) [5]int {

	n := len(arr)
	var revArr [5]int
	for i := 0; i < n; i++ {
		revArr[i] = arr[n-i-1]
	}

	return revArr
}

func reverse_Array2(arr [5]int) [5]int {
	n := len(arr)

	for i := 0; i < n/2; i++ {
		temp := arr[i]
		arr[i] = arr[n-i-1]
		arr[n-i-1] = temp
	}

	return arr
}

func slice_Element_Sum(slc []int) int {
	var sum int

	for _, value := range slc {
		sum += value
	}

	return sum
}

func max_min_Array(arr [6]int) (int, int) {
	maxi := arr[0]
	mini := arr[0]

	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] > maxi {
			maxi = arr[i]
		}
		if arr[i] < mini {
			mini = arr[i]
		}
	}

	return maxi, mini
}

func remove_duplicate_slice(slc []int) []int {
	m := make(map[int]int)
	var slice_set = []int{}

	for _, value := range slc {
		m[value]++
	}

	for key, _ := range m {
		slice_set = append(slice_set, key)
	}

	return slice_set

}

func rotate_Array_K(arr [6]int, k int) [6]int {
	n := len(arr)
	k = k % n
	var rotatedArr [6]int
	for i := 0; i < n; i++ {
		if i < k {
			rotatedArr[i] = arr[n+i-k]
		} else {
			rotatedArr[i] = arr[i-k]
		}
	}

	return rotatedArr
}

func find_Even_Numbers(slc []int) []int {
	var evenNums []int

	for _, value := range slc {
		if value%2 == 0 {
			evenNums = append(evenNums, value)
		}
	}
	return evenNums
}

func count_Occur_Elem(slc []int, target int) int {
	count := 0

	for _, value := range slc {
		if value == target {
			count++
		}
	}
	return count
}

func word_Freq(str string) map[string]int {
	words := strings.Fields(strings.ToLower(str))

	freq := make(map[string]int)

	for _, word := range words {
		freq[word]++
	}
	return freq
}

func char_Freq(str string) map[rune]int {

	freq := make(map[rune]int)

	for _, char := range str {
		freq[char]++
	}
	return freq
}

func two_Sum(arr [6]int, target int) []int {
	m := make(map[int]int)

	for i, value := range arr {
		rem := target - value
		if index, found := m[rem]; found {
			return []int{index, i}
		}
		m[value] = i
	}

	return nil
}

func duplicates_In_Array(arr [6]int) bool {
	m := make(map[int]bool)
	for _, value := range arr {
		if m[value] {
			return true
		}
		m[value] = true
	}
	return false
}

func first_Unique_Char(str string) int {
	freq := make(map[rune]int)
	for _, char := range str {
		freq[char]++
	}

	for i, char := range str {
		if freq[char] == 1 {
			return i
		}
	}

	return -1
}

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	fib := []int{0, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib
}

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}
	org, rev := num, 0

	for num > 0 {
		rem := num % 10
		rev = rev*10 + rem
		num /= 10
	}
	return org == rev
}

func find_Missing_Number(arr [6]int) int {
	n := len(arr) + 1
	expectedSum := n * (n + 1) / 2
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return expectedSum - sum
}
