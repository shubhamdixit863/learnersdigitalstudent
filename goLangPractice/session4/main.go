package main

import (
	"fmt"
	"math"
	"strings"
)

func arrReverse(arr *[5]int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

func slcSum(slc []int) int {
	sum := 0
	for _, v := range slc {
		sum += v
	}
	return sum
}

func minMax(arr [5]int) (int, int) {
	minNum := math.MaxInt64
	maxNum := math.MinInt64

	for i, _ := range arr {
		if arr[i] < minNum {
			minNum = arr[i]
		}
		if arr[i] > maxNum {
			maxNum = arr[i]
		}
	}
	return minNum, maxNum
}

func remDuplicate(slc *[]int) {
	slice := *slc
	n := len(slice)

	for i := 0; i < n; i++ {
		for j := i + 1; j < len(slice); { // Dynamically update length
			if slice[i] == slice[j] {
				// Remove element at index j
				slice = append(slice[:j], slice[j+1:]...)
			} else {
				j++ // Move to the next element only if no removal
			}
		}
	}
	*slc = slice // Update the original slice
}

func removeDuplicatesUsingMap(slc []int) []int {
	seen := make(map[int]bool) // Map to track seen elements
	result := []int{}          // New slice for unique elements

	for _, num := range slc {
		if !seen[num] { // If not already seen, add to result
			seen[num] = true
			result = append(result, num)
		}
	}

	return result // Return slice with unique elements
}

func rotateArr(arr [5]int, k int) []int {
	//i := 0
	//for i < k {
	//	arr[0], arr[1], arr[2], arr[3], arr[4] = arr[4], arr[0], arr[1], arr[2], arr[3]
	//	i++
	//}
	n := len(arr)

	newArr := make([]int, n)

	k = k % n

	for i := 0; i < n; i++ {
		newArr[(i+k)%n] = arr[i]
	}
	return newArr
}

func onlyEven(slc []int) []int {
	var result []int
	for _, num := range slc {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

func countInt(slc []int, num int) int {
	count := 0
	for _, val := range slc {
		if val == num {
			count++
		}
	}
	return count
}

func wordFreq(s string) map[string]int {
	mp := make(map[string]int)

	words := strings.Fields(strings.ToLower(s))
	for _, word := range words {
		mp[word]++
	}

	return mp
}

func countChar(s string) map[string]int {
	mp := map[string]int{}
	n := len(s)

	for i := 0; i < n; i++ {
		mp[string(s[i])]++
	}
	return mp
}

func twoSum(nums [5]int, target int) (int, int) {
	mp := make(map[int]int)

	for i, num := range nums {
		_, ok := mp[num]
		if ok {
			return mp[num], i
		} else {
			mp[target-num] = i
		}
	}
	return -1, -1
}

func checkDuplicate(arr [8]int) bool {
	mp := make(map[int]int)

	for _, num := range arr {
		mp[num]++
	}

	for _, v := range mp {
		if v > 1 {
			return true
		}
	}
	return false
}

func findUnique(s string) int {
	mp := make(map[string]int)

	n := len(s)
	for i := 0; i < n; i++ {
		mp[string(s[i])]++
	}

	for k, v := range mp {
		if v == 1 {

		}
	}
}

func main() {
	// Que: 1
	//arr := [5]int{1, 2, 3, 4, 5}
	//fmt.Println("Original Array: ", arr)
	//arrReverse(&arr)
	//fmt.Println("Reversed Array: ", arr)

	// Que: 2
	//slc := []int{1, 2, 3, 4, 5}
	//fmt.Println(slcSum(slc))

	// Que: 3
	//arr := [5]int{12, 87, 34, 14, 45}
	//minNum, maxNum := minMax(arr)
	//fmt.Println("minNum:", minNum, "maxNum:", maxNum)

	// Que: 4
	//s := []int{2, 3, 5, 2, 7, 11, 3, 3, 2, 2, 7, 13, 7, 11, 7}
	//fmt.Println(s)
	//remDuplicate(&s)
	//fmt.Println(s)

	//s := []int{2, 3, 5, 2, 7, 11, 3, 3, 2, 2, 7, 13, 7, 11, 7}
	//fmt.Println(s)
	//s1 := removeDuplicatesUsingMap(s)
	//fmt.Println(s1)

	// Que: 5
	//arr := [5]int{1, 2, 3, 4, 5}
	//fmt.Println("Original Array: ", arr)
	//arr1 := rotateArr(arr, 2)
	//fmt.Println("Rotated Array: ", arr1)

	// Que: 6
	//s := []int{2, 3, 5, 2, 7, 11, 3, 3, 2, 2, 7, 13, 7, 11, 7}
	//fmt.Println(s)
	//s1 := onlyEven(s)
	//fmt.Println(s1)

	// Que: 7
	//s := []int{2, 3, 5, 2, 7, 11, 3, 3, 2, 2, 7, 13, 7, 11, 7}
	//fmt.Println(s)
	//s1 := countInt(s, 7)
	//fmt.Println(s1)

	// Que: 8
	//s := "Hello world! Hello Go. Go is fun, and Go is powerful."
	//mp := wordFreq(s)
	//fmt.Println(mp)

	// Que: 9
	//s := "abcdacdbefg"
	//mp := countChar(s)
	//fmt.Println(mp)

	// Que: 10
	//arr := [5]int{1, 2, 3, 4, 5}
	//i1, i2 := twoSum(arr, 4)
	//fmt.Println(i1, i2)

	// Que: 11
	//arr := [8]int{1, 7, 2, 8, 3, 4, 5, 9}
	//b := checkDuplicate(arr)
	//fmt.Println(b)

	// Que: 12
	s := "abcbadcdefgfe"
	i := findUnique(s)
	fmt.Println(i)
}
