package main

import (
	"fmt"
	"session4/booking"
	"sort"
)

//	func fill(m map[string]int) {
//		m["imran"] = 102
//	}
func reverse(arr []int) {
	k := len(arr) / 2
	j := len(arr) - 1

	for i := 0; i < k; i++ {

		a := arr[i]
		arr[i] = arr[j]
		arr[j] = a
		j--

	}
	fmt.Println(arr)
}
func sumSlice(s []int) {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	fmt.Println(sum)

}
func maxMin(arr []int) {
	sort.Ints(arr)
	fmt.Println(arr[0], arr[len(arr)-1])

}
func removeDuplicate(nums []int) int {
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	// var newarr []int
	i := 0
	for key, _ := range mp {
		nums[i] = key
		i++
	}
	// sort.Ints(newarr)
	// nums = newarr
	// fmt.Println(newarr)
	fmt.Println(nums)

	return i

}
func reverseByksteps(arr []int, k int) {
	arr1 := arr[k+1:]
	arr2 := arr[:k+1]
	arr = append(arr1, arr2...)
	fmt.Println(arr)

}
func towSum(arr []int, target int) (int, int) {
	mp := make(map[int]int)
	var i int
	var j int
	for i := 0; i < len(arr); i++ {
		mp[arr[i]] = i
	}
	for k, v := range mp {
		next := target - k
		idx, ok := mp[next]
		if ok {
			i = v
			j = idx
			break
		}

	}
	return i, j
}

func main() {

	// var my_slice_1 = make([]int, 4, 7)
	// fmt.Println(my_slice_1)
	// p := make(map[string]int)
	// fmt.Println(p == nil)
	// // fmt.Println(p["hello"])
	// p["name"] = 1
	// p["imran"] = 7
	// fmt.Println(p)
	// var m map[string]int

	// fmt.Println(m == nil)
	// fmt.Println(m["hee"])
	// fill(p)
	// fmt.Println(p["imran"])
	// var slice_1 = make([]int, 4, 7)
	// var slice_2 = make([]int, 4, 7)
	// // fmt.Println(slice_1 == slice_2)
	// m1 := make(map[string]int)
	// m2 := make(map[string]int)
	// // fmt.Println(m1 == m2)
	// aar1 := [2]int{1, 2}
	// aar2 := [2]int{1, 2}
	// fmt.Println(aar1 == aar2)
	// arr := []int{1, 1, 2}
	// fmt.Println(removeDuplicate(arr))
	// fmt.Println(arr)

	// reverse(arr)
	// sumSlice(arr)
	// maxMin(arr)
	// removeDuplicate(arr)
	// fmt.Println(towSum(arr, 6))
	booking.Sepprating()

}
