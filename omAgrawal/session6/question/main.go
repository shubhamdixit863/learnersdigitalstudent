// package main

// import "fmt"

// func main() {

// 	arr := []int{1, 2, 3, 4, 5}

// 	s := arr[1:]

// 	fmt.Println(len(s), cap(s)) // Length and capacity before appending

// 	s = append(s, 10, 20, 30)

// 	fmt.Println(arr, s)

// }

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4}

	s1 := arr[1:2]
	fmt.Println(len(s1), cap(s1))
	s2 := arr[1:]
	fmt.Println(len(s2), cap(s2))

	s1 = append(s1, 10, 20,30) // Modify slice s1

	fmt.Println(arr, s1, s2)

}
