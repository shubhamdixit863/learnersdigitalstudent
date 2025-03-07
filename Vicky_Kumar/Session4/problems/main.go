package main

import (
	"fmt"
	"math"
)

func main() {
	//1
	arr := [5]int{1, 2, 3, 4, 5}
	revArr := Reverse(arr)
	fmt.Println("Reversed array is:", revArr)
	//2
	var slc = []int{1, 2, 3, 4, 4, 5}
	fmt.Println("sum of slice is:", Sum(slc))
	//3
	minm, maxm := MinMax(arr)
	fmt.Println("Minm is", minm, "and Maxm is", maxm)
	//4
	fmt.Println("Removed Duplicates from Slice:", RemoveDuplicate(slc))

	//5
	fmt.Println("Rotate array by k steps:", Rotate(arr, 2))

	//6
	fmt.Println("Even slice is:", EvenSlice(slc))

	//7
	fmt.Println("no. of occurace of 2 is:", OccCount(slc, 2))

	//8
	wordCount("This This is is is Vicky")

	//9
	fmt.Println("character count is:", CharCount("something"))

	// fmt.Println([]byte("こんにちは"))
	// fmt.Println([]rune("こんにちは"))

	//10
	idx1, idx2 := TwoSum(arr, 7)
	fmt.Println("two sum indx: ", idx1, idx2)

	//11
	arr2 := [5]int{1, 2, 2, 1, 3}
	fmt.Println("Removed Duplicates from array:", DuplicateArray(arr2))

	//12
	fmt.Println("Index of first non repeating character:", IdxOfNonRepeating("helloh"))

	//13
	Fibonacci(10)

	//14

	fmt.Println("IS palindrome:1001 ", Palindrome(1001))
	//15
	arr3 := []int{1, 2, 3, 4, 6, 7}
	fmt.Println("missing number is:", MissingNumber(arr3))

}

// 1.	Reverse an Array:
// Write a function to reverse an array of integers.
func Reverse(arr [5]int) [5]int {
	for i := 0; i < len(arr)/2; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = temp
	}
	return arr
}

// 2.	Sum of Elements in a Slice:
// Create a function that takes a slice of integers and returns the sum of all elements.
func Sum(slc []int) int {
	sum := 0
	for i := 0; i < len(slc); i++ {
		sum += slc[i]
	}
	return sum
}

// 3.	Find Maximum and Minimum in an Array:
// Given an integer array, find and return the maximum and minimum elements.
func MinMax(arr [5]int) (int, int) {
	minm := math.MaxInt
	maxm := math.MinInt
	for i := 0; i < len(arr); i++ {
		minm = min(minm, arr[i])
		maxm = max(maxm, arr[i])
	}
	return minm, maxm
}

// 4.	Remove Duplicates from a Slice:
// Write a function to remove all duplicate integers from a slice.
func RemoveDuplicate(slc []int) []int {
	for i := 0; i < len(slc); i++ {
		for j := i + 1; j < len(slc); {
			if slc[i] == slc[j] {
				slc = append(slc[:j], slc[j+1:]...)
			} else {
				j++
			}
		}
	}
	return slc
}

//  5. Rotate Array by k Steps:
//
// Rotate the elements of an array to the right by k steps.
func Rotate(arr [5]int, k int) [5]int {
	var rarray [5]int
	for i := 0; i < len(arr); i++ {
		rarray[(i+k)%len(arr)] = arr[i]
	}
	return rarray
}

// 	6.	Find All Even Numbers in a Slice:

// Given a slice of integers, return a new slice containing only even numbers.
func EvenSlice(slc []int) []int {
	var eveSlc []int
	for i := 0; i < len(slc); i++ {
		if slc[i]&1 == 0 {
			eveSlc = append(eveSlc, slc[i])
		}
	}
	return eveSlc
}

// 	7.	Count Occurrences of an Element:

// Write a function that counts how many times a specific integer appears in a slice.
func OccCount(slc []int, k int) int {
	count := 0
	for i := 0; i < len(slc); i++ {
		if slc[i] == k {
			count++
		}
	}
	return count
}

// 	8.	Word Frequency Counter:
// Given a string, count the frequency of each word using a map.

func wordCount(sentence string) {
	word := ""
	m := make(map[string]int)
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' {
			m[word]++
			word = ""
		} else {
			word += string(sentence[i])
		}
	}
	m[word]++
	fmt.Println("Word Count is", m)
}

//  9. Character Frequency in a String:
//
// Write a function that returns a map with the count of each character in a given string.
func CharCount(str string) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(str); i++ {
		m[string(str[i])]++
	}
	return m
}

//  10. Two Sum Problem:
//
// Given an array and a target integer, find the indices of the two numbers that add up to the target using a map.
func TwoSum(arr [5]int, sum int) (int, int) {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = i
		_, ok2 := m[sum-arr[i]]
		if ok2 {
			return m[arr[i]], m[sum-arr[i]]
		}
	}
	return -1, -1
}

//  11. Check for Duplicates in an Array:
//
// Determine if an array contains any duplicate elements using a map.
func DuplicateArray(arr [5]int) []int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = i
	}
	var newArr []int
	for k := range m {
		newArr = append(newArr, k)
	}
	return newArr

}

// 	12.	First Unique Character in a String:
// Find the index of the first non-repeating character in a string using a map.

func IdxOfNonRepeating(str string) int {
	count := make(map[rune]int)
	for _, ch := range str {
		count[ch]++
	}
	for i, c := range str {
		if count[c] == 1 {
			return i
		}
	}
	return -1
}

// 13.	Fibonacci Sequence Generator:
// Write a function that generates the first n Fibonacci numbers using a for loop.
func Fibonacci(num int) {
	a := 0
	b := 1
	for i := 0; i < num; i++ {
		fmt.Println(a)
		c := a
		a = b
		b = b + c
	}
}

//  14. Palindrome Number Checker:
//
// Write a function to check if a given integer is a palindrome (without converting it to a string).
func Palindrome(num int) bool {
	n := num
	rev := 0
	for {
		d := n % 10
		rev = rev*10 + d
		n /= 10
		if n == 0 {
			break
		}
	}
	return num == rev

}

//  15. Find Missing Number in a Range:
//
// Given an array containing n-1 distinct integers from 1 to n, find the missing number.
func MissingNumber(arr []int) int {
	sum := 0
	n := len(arr) + 1
	for i := 0; i < n-1; i++ {
		sum += arr[i]
	}
	return ((n * (n + 1)) / 2) - sum
}
