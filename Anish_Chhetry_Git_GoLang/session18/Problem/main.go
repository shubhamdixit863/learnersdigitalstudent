package main

import (
	"log"
)

func main() {
	slc := []int{4, 3, 1, 2, 5, 3, 1, 1}
	log.Println("unsorted:", slc)
	log.Println("sorted:", mergeSort(slc))
}

func mergeSort(slc []int) []int {
	if len(slc) <= 1 {
		return slc
	}
	mid := len(slc) / 2
	var left, right []int
	leftChan := make(chan []int)
	rightChan := make(chan []int)

	//fan out
	go func() {
		left = mergeSort(slc[:mid])
		leftChan <- left
		close(leftChan)
	}()

	go func() {
		right = mergeSort(slc[mid:])
		rightChan <- right
		close(rightChan)
	}()

	//fan in
	leftPart, rightPart := <-leftChan, <-rightChan
	return merge(leftPart, rightPart)

}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
