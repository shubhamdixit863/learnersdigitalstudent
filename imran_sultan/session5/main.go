package main

import (
	"fmt"
	"session5/practice"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	var nums3 []int
	for i < m && j < n {
		if nums1[i] < nums2[j] {

			nums3 = append(nums3, nums1[i])
			i++
		} else {
			nums3 = append(nums3, nums2[j])
			j++
		}
	}
	for i < m {

		nums3 = append(nums3, nums1[i])
		i++
	}
	for j < n {
		nums3 = append(nums3, nums2[j])
		j++
	}
	nums1 = nums3
	fmt.Println(nums1)
}

func main() {
	// nums1 := []int{1, 2, 3, 4}
	// nums2 := []int{5, 6, 7, 8}

	// merge(nums1, 4, nums2, 4)
	practice.SetValue()
}
