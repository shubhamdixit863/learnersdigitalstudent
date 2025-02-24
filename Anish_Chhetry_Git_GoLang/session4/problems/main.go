package main

import (
	"fmt"
	"strings"
)

func main() {
	// 	1.	Reverse an Array:
	// Write a function to reverse an array of integers.
	// The function should take an array of integers as input and return the reversed array.
	Array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original Array:", Array1)
	fmt.Println("Reversed Arrat:", reverseArray(Array1))

	// 	2.	Sum of Elements in a Slice:
	// Create a function that takes a slice of integers and returns the sum of all elements.
	Array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original Array:", Array2)
	fmt.Println("Sum is:", sumArray(Array2))

	// 	3.	Find Maximum and Minimum in an Array:
	// Given an integer array, find and return the maximum and minimum elements.
	Array3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original Array:", Array3)
	max, min := maxMin(Array3)
	fmt.Printf("Max is %d and Min is %d:\n", max, min)

	// 	4.	Remove Duplicates from a Slice:
	// Write a function to remove all duplicate integers from a slice.
	Array4 := []int{1, 2, 2, 3, 3}
	fmt.Println("Original Array:", Array4)
	fmt.Println("Array after removing duplicates:", removeDup(Array4))

	// 	5.	Rotate Array by k Steps:
	// Rotate the elements of an array to the right by k steps.
	Array5 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original Array:", Array5)
	fmt.Println("Array after rotation:", rotateArray(Array5, 2))

	// 	6.	Find All Even Numbers in a Slice:
	// Given a slice of integers, return a new slice containing only even numbers.
	Slice6 := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice: ", Slice6)
	fmt.Println("Even numbers in the slice:", evenNumbers(Slice6))

	// 	7.	Count Occurrences of an Element:
	// Write a function that counts how many times a specific integer appears in a slice.
	Slice7 := []int{1, 1, 2, 3, 1}
	fmt.Println("Original Slice: ", Slice7)
	fmt.Println("Total count of the specific integer is:", countOcc(Slice7, 1))

	// 	8.	Word Frequency Counter:
	// Given a string, count the frequency of each word using a map.
	String8 := "Hello world Hello world Hello world"
	fmt.Println("Original String: ", String8)
	fmt.Println("Word frequency counter: ", wordFrequency(String8))

	// 	9.	Character Frequency in a String:
	// Write a function that returns a map with the count of each character in a given string.
	String9 := "Hello world"
	fmt.Println("Original String: ", String9)
	fmt.Println("Character frequency counter: ", charFrequency(String9))

	// 	10.	Two Sum Problem:
	// Given an array and a target integer, find the indices of the two numbers that add up to the target using a map.
	Array10 := [5]int{1, 2, 3, 4, 5}
	target := 8
	fmt.Println("Original Array:", Array10)
	fmt.Println("Indices of the two numbers that add up to the target", target, ": ", twoSum(Array10, target))

	// 	11.	Check for Duplicates in an Array:
	// Determine if an array contains any duplicate elements using a map.
	Array11 := [5]int{1, 2, 3, 4, 1}
	fmt.Println("Original Array:", Array11)
	fmt.Println(checkDuplicates(Array11))

	// 	12.	First Unique Character in a String:
	// Find the index of the first non-repeating character in a string using a map.
	String12 := "aabcdeee"
	fmt.Println("Original String: ", String12)
	fmt.Println("Index of the first non-repeating character: ", firstUniqueChar(String12))

	// 13.	Fibonacci Sequence Generator:
	// Write a function that generates the first n Fibonacci numbers using a for loop.
	n := 10
	fmt.Println("Fibonacci sequence up to", n, "numbers: ", fibonacci(n))

	//  14. Palindrome Number Checker:
	// Write a function to check if a given integer is a palindrome (without converting it to a string).
	num := 12321
	fmt.Println("Is the number", num, "a palindrome?", isPalindrome(num))

	//  15. Find Missing Number in a Range:
	// Given an array containing n-1 distinct integers from 1 to n, find the missing number.
	Array15 := [5]int{1, 2, 4, 5}
	fmt.Println("Original Array:", Array15)
	fmt.Println("Missing number in the range: ", findMissingNumber(Array15))

}
func reverseArray(arr [5]int) []int {
	var reversed []int
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}
func sumArray(arr [5]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
func maxMin(arr [5]int) (int, int) {
	max := arr[0]
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}
func removeDup(arr []int) []int {
	var newarr []int
	for i := 0; i < len(arr); i++ {
		if !contains(newarr, arr[i]) {
			newarr = append(newarr, arr[i])
		}
	}
	return newarr
}
func contains(arr []int, val int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}
func rotateArray(arr [5]int, k int) []int {
	var newarr []int
	for i := len(arr) - k; i < len(arr); i++ {
		newarr = append(newarr, arr[i])
	}
	for i := 0; i < len(arr)-k; i++ {
		newarr = append(newarr, arr[i])
	}
	return newarr
}
func evenNumbers(slice []int) []int {
	var newarr []int
	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			newarr = append(newarr, slice[i])
		}
	}
	return newarr
}
func countOcc(slice []int, num int) int {
	var count int
	for i := 0; i < len(slice); i++ {
		if slice[i] == num {
			count++
		}
	}
	return count
}
func wordFrequency(phrase string) map[string]int {
	wordFreq := make(map[string]int)
	for _, word := range strings.Fields(phrase) {
		wordFreq[word]++
	}
	return wordFreq
}
func charFrequency(phrase string) map[rune]int {
	charFreq := make(map[rune]int)
	for _, char := range phrase {
		charFreq[char]++
	}
	return charFreq
}
func twoSum(nums [5]int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if numMap[target-nums[i]] > 0 {
			return []int{numMap[target-nums[i]], i}
		}
		numMap[nums[i]] = i
		fmt.Println(numMap)
	}
	return []int{}
}
func checkDuplicates(arr [5]int) string {
	m := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if m[arr[i]] {
			return "Duplicates found"
		}
		m[arr[i]] = true
	}
	return "No Duplicates found"
}
func firstUniqueChar(phrase string) int {
	charMap := make(map[rune]int)
	for _, char := range phrase {
		charMap[char]++
	}
	for i, char := range phrase {
		if charMap[char] == 1 {
			return i
		}
	}
	return -1
}
func fibonacci(n int) []int {
	var fib []int
	fib = append(fib, 0)
	fib = append(fib, 1)
	for i := 2; i <= n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib
}

func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	reversed := 0
	original := n
	for n > 0 {
		reversed = reversed*10 + n%10
		n = n / 10
	}
	return original == reversed
}
func findMissingNumber(arr [5]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += i + 1
	}
	for i := 0; i < len(arr); i++ {
		sum -= arr[i]
	}
	return sum
}
