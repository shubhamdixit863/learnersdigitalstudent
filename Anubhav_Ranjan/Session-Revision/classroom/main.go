package main

import (
	"classroom/add"
	mySubtract "classroom/subtract"
	"classroom/subtract/delete"
	"fmt"
)

const h = 9

var (
	i = 1
	j = 2
)

func main() {
	// const h = 9 // generlly we dont do it..we define outside the function.
	fmt.Println("Hello Beautiful World!")

	var c int
	c = 9
	fmt.Println("C =", c)

	y := 14
	fmt.Println("Y=", y)

	add.Sum()
	mySubtract.Sub()
	delete.Delt()

	var arr [6]int
	fmt.Print("Array:", arr)

}

// func main() {

// 	arr := []int{1, 2, 3, 4}

// 	s1 := arr[:2]

// 	s2 := arr[2:3]

// 	s1 = append(s1, 10) // Modify slice s1
// 	fmt.Println(len(s1), cap(s1))
// 	fmt.Println(len(s2), cap(s2))
// 	fmt.Println(arr, s1, s2)

// }
