// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func comparison() {
// 	m1 := make(map[string]int)
// 	m1["hello"] = 1
// 	m2 := make(map[string]int)
// 	m2["hello"] = 1
// }

// func reverse(arr1 [5]int) {
// 	for i := len(arr1) - 1; i >= 0; i-- {
// 		fmt.Println(arr1[i])
// 	}
// }

// func sum(myslice []int) {
// 	var sum int
// 	for i := 0; i < len(myslice); i++ {
// 		sum = sum + myslice[i]
// 	}
// 	fmt.Println(sum)
// }

// func max(arr2 [5]int) (a int, b int) {
// 	var ma int = arr2[0]
// 	var mi int = arr2[0]
// 	for i := 0; i < len(arr2); i++ {
// 		if ma < arr2[i] {
// 			ma = arr2[i]
// 		} else if mi > arr2[i] {
// 			mi = arr2[i]
// 		}
// 	}
// 	return ma, mi
// }

// func remove(myslice2 []int) []int {
// 	for i := 0; i < len(myslice2); i++ {
// 		for j := i + 1; j < len(myslice2); j++ {
// 			if myslice2[i] == myslice2[j] {
// 				myslice2 = slices.Delete(myslice2, j, j+1)
// 			}
// 		}
// 	}
// 	return myslice2
// }

// func rotate(arr []int, k int) {
// 	n := len(arr)
// 	k = k % n
// 	copy(arr, append(arr[n-k:], arr[:n-k]...))
// 	fmt.Println(arr)
// }

// func even(myslice4 []int) []int {
// 	var newslice []int
// 	for i := 0; i < len(myslice4); i++ {
// 		if myslice4[i]%2 == 0 {
// 			newslice = append(newslice, myslice4[i])
// 		}
// 	}
// 	return newslice
// }

// func occurance(myslice5 []int) {
// 	var count int = 1
// 	for i := 0; i < len(myslice5); i++ {
// 		count = 1
// 		for j := i + 1; j < len(myslice5); j++ {
// 			if myslice5[i] == myslice5[j] {
// 				count++
// 				myslice5 = slices.Delete(myslice5, j, j+1)
// 			}
// 		}
// 		fmt.Println(myslice5[i], count)
// 	}
// }

// func wordFrequency(sentence string) {
// 	p := make(map[string]int)
// 	words := []rune(sentence)
// 	for _, word := range words {
// 		p[string(word)]++
// 	}
// 	for key, value := range p {
// 		fmt.Println(key, value)
// 	}
// }

// func twoSum(nums []int, target int) {
// 	m := make(map[int]int)
// 	for i, num := range nums {
// 		if j, found := m[target-num]; found {
// 			fmt.Println(i, j)
// 			return
// 		}
// 		m[num] = i
// 	}
// }

// func checkDuplicates(arr []int) {
// 	m := make(map[int]bool)
// 	for _, num := range arr {
// 		if m[num] {
// 			fmt.Println("Duplicate exists")
// 			return
// 		}
// 		m[num] = true
// 	}
// 	fmt.Println("No duplicates")
// }

// func fibonacci(n int) {
// 	a, b := 0, 1
// 	for i := 0; i < n; i++ {
// 		fmt.Println(a)
// 		a, b = b, a+b
// 	}
// }

// func palindromeNumber(n int) {
// 	rev, temp := 0, n
// 	for temp > 0 {
// 		digit := temp % 10
// 		rev = rev*10 + digit
// 		temp /= 10
// 	}
// 	if rev == n {
// 		fmt.Println("Palindrome")
// 	} else {
// 		fmt.Println("Not Palindrome")
// 	}
// }

// func missingNumber(arr []int, n int) {
// 	total := n * (n + 1) / 2
// 	sum := 0
// 	for _, num := range arr {
// 		sum += num
// 	}
// 	fmt.Println("Missing number:", total-sum)
// }

// func main() {
// 	p := make(map[string]int)
// 	p["name"] = 2
// 	p["john"] = 3
// 	fmt.Println(p)
// 	arr1 := [5]int{1, 2, 3, 4, 5}
// 	reverse(arr1)
// 	myslice := []int{1, 2, 3, 4, 5}
// 	sum(myslice)
// 	arr2 := [5]int{1, 2, 3, 4, 5}
// 	a, b := max(arr2)
// 	fmt.Println("Max", a, "Min", b)
// 	myslice2 := []int{5, 2, 3, 3, 5}
// 	myslice3 := remove(myslice2)
// 	fmt.Println(myslice3)
// 	rotate([]int{1, 2, 3, 4, 5}, 2)
// 	myslice4 := []int{1, 2, 3, 4, 5}
// 	newSlice := even(myslice4)
// 	fmt.Println(newSlice)
// 	myslice5 := []int{1, 2, 3, 2, 5}
// 	occurance(myslice5)
// 	wordFrequency("hello world")
// 	twoSum([]int{2, 7, 11, 15}, 9)
// 	checkDuplicates([]int{1, 2, 3, 4, 2})
// 	fibonacci(5)
// 	palindromeNumber(121)
// 	missingNumber([]int{1, 2, 4, 5}, 5)
// }
