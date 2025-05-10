package main

import "fmt"

// 1
// func main() {

// 	arr := []int{1, 2, 3, 4}

// 	s1 := arr[:2]

// 	s2 := arr[1:3]

// 	s1 = append(s1, 10) // Modify slice s1

// 	// sliceAddress0 := &arr
// 	// sliceAddress1 := &s1
// 	// sliceAddress2 := &s2
// 	// fmt.Printf("The address of the slice is: %p\n", sliceAddress0)
// 	// fmt.Printf("The address of the slice is: %p\n", sliceAddress1)
// 	// fmt.Printf("The address of the slice is: %p\n", sliceAddress2)
// 	fmt.Println(arr, s1, &s2)

// }

//2
// func main() {

// 	arr := []int{1, 2, 3, 4, 5}

// 	s := arr[1:3] // 2, 3

// 	fmt.Println(len(s), cap(s)) // Length and capacity before appending
// 	// len is 2, cap is 4
// 	//if we insert upto two ele it will update same memory   but if insert more elements then new memory crested, will not affect original array
// 	s = append(s, 10, 20, 30)

// 	fmt.Println(arr, s)
// 	// arr=1 2 3 4 5
// 	//s 2,3,10,20,30

// }

//3
// func modifySlice(s []int) {

// 	s[0] = 100

// }

// func main() {

// 	arr := []int{1, 2, 3, 4, 5}

// 	s := arr[:3] // 1 2 3

// 	modifySlice(s)

// 	fmt.Println(arr)
// 	//100 2 3 4 5

// }

//4
// func main() {

// 	s1 := []int{1, 2, 3, 4}

// 	s2 := make([]int, len(s1))

// 	copy(s2, s1)

// 	s1[0] = 100

// 	fmt.Println(s1, s2)
// 	//s1 = 100 2 3 4
// 	//s2 = 1 2 3 4

// 	s1 = append(s1, 200)

// 	fmt.Println(s1, s2)
// 	//s1 = 100 2 3 4 200
// 	//s2 = 1 2 3 4

// }

//5
func main() {

	s := make([]int, 2, 3)

	s[0], s[1] = 10, 20

	s1 := append(s, 30)
	//s1 = 10 20 30

	s2 := append(s, 40)
	//s2 = 10 20 40

	fmt.Println(s, s1, s2)
	//s = 10 20
	//s1 = 10 20 40
	//s2 = 10 20 40

}
