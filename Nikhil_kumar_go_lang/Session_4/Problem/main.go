package main

import (
	"fmt"
)

func Reverse_and_Array(ar []int, n int) []int {
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = ar[n-i-1]

	}
	return temp
}
func Sum_of_element(sm []int) {
	sum := 0
	for i := 0; i < len(sm); i++ {
		sum += sm[i]
	}
	fmt.Println("sum of slice :", sum)
}

func max_min(arr []int) (int, int) {
	var max int = arr[0]
	var min int = arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	return max, min
}

func Remove_duplicate(arr []int) []int {
	var temp []int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				i++
				continue
			}
		}
		temp = append(temp, arr[i])
	}
	return temp
}

func rotate_array_by_k(sl []int, k int) {
	sl = append(sl[len(sl)-k:], sl[:len(sl)-k]...)
	fmt.Println("rotated array by k : ", sl)
}

func even_num_from_the_slice(sl []int) {
	for i := 0; i < len(sl); i++ {
		if sl[i]%2 == 0 {
			fmt.Println("even number : ", sl[i])
		}
	}
}

func count_occurrence(sl []int) {
	m := make(map[int]int)
	for i := 0; i < len(sl); i++ {
		m[(sl[i])]++
	}
	for k, v := range m {
		fmt.Println(k, "occurs", v, "times")
	}
}

func word_frequency_counter(str []string) {
	m := make(map[string]int)
	for i := 0; i < len(str); i++ {
		m[(str[i])]++
	}
	for k, v := range m {
		fmt.Println(k, "occurs", v, "times")
	}

}

func Character_Frequency() {
	name := "nikhil kumar"
	m := make(map[rune]int)
	for i := 0; i < len(name); i++ {
		m[rune(name[i])]++
	}
	for k, v := range m {
		fmt.Printf("%c occurs %d times\n", k, v)
	}
}

func Two_sum(arr []int, target int) {

	for i := 0; i < len(arr); i++ {
		flag := target - arr[i]
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			if arr[j] == flag {
				fmt.Println(" i : ", i, "arr :", arr[i], " j ", j, "arr : ", arr[j])
				break
			}
		}
	}
}

func Two_sum2(arr []int, target int) {
	mp := make(map[int]int)
	for idx, val := range arr {
		fl := target - val
		if idx2, ok := mp[fl]; ok {
			fmt.Println(val, idx2, idx)
			break
		}
		mp[val] = idx
	}
	fmt.Println("mp : ", mp)
}

func check_dup(arr []int) {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	for k, v := range m {
		if v > 1 {
			fmt.Println("duplicate element : ", k)
		}
	}
}

func unique_chat(str string) {
	m := make(map[rune]int)
	for i := 0; i < len(str); i++ {
		m[rune(str[i])]++
	}
	for k, v := range m {
		if v == 1 {
			fmt.Printf("unique character : %c \n", k)
		}
	}
}

func fibonacci(n int) {
	var f1 int
	var f2 int
	var f3 int
	f1 = 1
	f2 = 1
	fmt.Println("fibonaci : ")
	for i := 1; i <= n; i++ {
		fmt.Println(f1)
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
}

func check_palindrome(n int) {
	var temp int = n
	var rev int = 0
	for n > 0 {
		rev = rev*10 + n%10
		n = n / 10
	}
	if temp == rev {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}

func missing_number_in_range(arr []int) {
	var n int = arr[len(arr)-1]
	var sum int = (n * (n + 1)) / 2
	for i := 0; i < len(arr); i++ {
		sum = sum - arr[i]
	}
	fmt.Println("missing number in range : ", sum)
}

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(Reverse_and_Array(arr[:], len(arr)))

	var sum_of []int = []int{1, 2, 3, 4, 5}
	Sum_of_element(sum_of[:])
	max_min(sum_of[:])
	mx, mn := max_min(sum_of[:])
	fmt.Println("max and min : ", mx, mn)

	// var arr1 []int = []int{1, 1, 1, 1, 2, 3, 5, 4, 6, 8}
	var arr1 []int = []int{1, 1, 1, 1, 2, 3, 5, 4, 6, 8}
	fmt.Println("original slice : ", arr1)
	fmt.Println("remove duplicate : ", Remove_duplicate(arr1))
	rotate_array_by_k(arr1, 3)
	even_num_from_the_slice(arr1)

	count_occurrence(arr1)
	var str []string = []string{"nikhil", "kumar", "nikhil", "nikhil"}
	word_frequency_counter(str)

	Character_Frequency()
	// Two_sum(arr1, 2)
	Two_sum2(arr1, 2)

	check_dup(arr1)
	unique_chat("nikhil")
	fibonacci(9)

	check_palindrome(133331)

	var arrrr [5]int = [5]int{1, 2, 3, 4, 6}

	missing_number_in_range(arrrr[:])

}
